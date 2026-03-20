# gols — Tu propio comando `ls` en Go

Reimplementación del comando `ls` de Unix con colores, iconos y opciones de ordenamiento.

## Arquitectura general

El proyecto tiene **dos archivos de código**:

| Archivo | Responsabilidad |
|---------|----------------|
| `gols.go` | Define los **tipos de datos** (structs, constantes) y los **estilos visuales** |
| `main.go` | Contiene la **lógica del programa**: parseo de flags, lectura de directorio, filtrado, ordenamiento e impresión |

En Go, todos los archivos del mismo `package` comparten scope. Por eso `main.go` puede usar `file`, `styleFileType` y `mapStyleByFileType` definidos en `gols.go` sin importarlos — son del mismo paquete `main`.

---

## Tipos y estilos (`gols.go`)

### Constantes con `iota`

```go
type FileType int

const (
    fileRegular    FileType = iota  // 0
    fileDirectory                   // 1
    fileExecutable                  // 2
    fileCompress                    // 3
    fileImage                       // 4
    fileLink                        // 5
)
```

`iota` es el generador de enumeraciones de Go. Empieza en 0 y se incrementa automáticamente en cada línea del bloque `const`. Es la forma idiomática de crear "enums" en Go. `type FileType int` crea un tipo distinto de `int`, lo que aporta **type safety**: el compilador rechaza asignar un `int` arbitrario a un `FileType` sin conversión explícita.

### El struct `file`

```go
type file struct {
    name             string
    fileType         FileType
    isDir            bool
    isHidden         bool
    userName         string
    groupName        string
    size             int64
    modificationTime time.Time
    mode             string
}
```

Modelo central del programa. Cada archivo/directorio que se lee del sistema se convierte en un `file`. `fileType` usa el tipo `FileType` que corresponde a las constantes `iota`.

### Mapa de estilos

```go
var mapStyleByFileType = map[FileType]styleFileType{
    fileRegular:    {icon: "📄"},
    fileDirectory:  {icon: "🗂️", color: color.FgBlue, symbol: "/"},
    fileExecutable: {icon: "🚀", color: color.FgGreen, symbol: "*"},
    // ...
}
```

Patrón de **mapa de configuración**: en vez de usar un `switch` enorme para decidir cómo mostrar cada tipo de archivo, se centraliza la configuración en un mapa. Agregar un nuevo tipo de archivo es tan simple como agregar una línea.

### Funciones de color

```go
var (
    blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
    green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
    // ...
)
```

`SprintFunc()` retorna una **función** (`func(a ...interface{}) string`). `blue`, `green`, etc. no son strings — son **funciones** que puedes llamar: `blue("hola")` devuelve el string "hola" coloreado en azul. Este patrón se llama **function value** en Go.

---

## Lógica del programa (`main.go`)

### Paso 1: Parseo de flags

```go
flagPattern := flag.String("p", "", "filter by pattern")
flagAll := flag.Bool("a", false, "all files including hidden files")
flagNumberRecords := flag.Int("n", 0, "number of records")
hasOrderByTime := flag.Bool("t", false, "sort by time")
hasOrderBySize := flag.Bool("s", false, "sort by size")
hasOrderReverse := flag.Bool("r", false, "reverse order")
flag.Parse()
```

Flags disponibles:

| Flag | Descripción |
|------|-------------|
| `-p "patron"` | Filtra archivos por expresión regular |
| `-a` | Muestra archivos ocultos |
| `-n 5` | Limita a N resultados |
| `-t` | Ordena por tiempo de modificación |
| `-s` | Ordena por tamaño |
| `-r` | Invierte el orden |

`flag.String()` retorna un `*string` (puntero), no un `string`. Por eso se usa `*flagPattern`, `*flagAll`, etc. — se **dereferencia** el puntero para obtener el valor. `flag.Arg(0)` obtiene el primer argumento **posicional** (no-flag) y se guarda en `dirPath`, permitiendo `gols -a /ruta/directorio`.

### Paso 2: Lectura del directorio y filtrado

```go
dirs, err := os.ReadDir(dirPath)
if err != nil {
    fmt.Fprintf(os.Stderr, "error reading directory: %v\n", err)
    os.Exit(1)
}

var pattern *regexp.Regexp
if *flagPattern != "" {
    pattern, err = regexp.Compile("(?i)" + *flagPattern)
    if err != nil {
        fmt.Fprintf(os.Stderr, "invalid pattern: %v\n", err)
        os.Exit(1)
    }
}
```

`os.ReadDir` lee todas las entradas del directorio. La regex se compila **una sola vez** antes del loop con `regexp.Compile`, en vez de recompilarla en cada iteración (mejora de rendimiento). Los errores se reportan a `os.Stderr` con un mensaje limpio y `os.Exit(1)`, en vez de usar `panic` que genera un stack trace difícil de leer para el usuario.

Se itera sobre las entradas aplicando **tres filtros**:

1. **Archivos ocultos**: si el archivo es oculto y no se pasó `-a`, se salta
2. **Patrón regex**: si se compiló un patrón (`pattern != nil`), solo incluye archivos cuyo nombre coincida. El prefijo `(?i)` activa case-insensitive matching
3. **Construcción del struct**: cada archivo que pasa los filtros se convierte en un `file` via `getFile()`

### Paso 3: Clasificación del tipo de archivo — `setFile()`

```go
func setFile(f *file) {
    switch {
    case isLink(*f):
        f.fileType = fileLink
    case f.isDir:
        f.fileType = fileDirectory
    case isExecutable(*f):
        f.fileType = fileExecutable
    case isCompress(*f):
        f.fileType = fileCompress
    case isImage(*f):
        f.fileType = fileImage
    default:
        f.fileType = fileRegular
    }
}
```

- **Puntero como parámetro** (`f *file`): recibe un puntero para poder **modificar** el struct original. Sin el `*`, Go haría una copia y los cambios se perderían.
- **`switch` sin expresión** (`switch {}`): "switch booleano" — cada `case` es una condición booleana independiente. Se evalúan en orden y el primero que sea `true` gana. El orden importa: un enlace simbólico podría también ser ejecutable, pero se clasifica como enlace primero.

### Paso 4: Ordenamiento con genéricos

```go
func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
    if isReverse {
        return i > j
    }
    return i < j
}
```

`mySort[T constraints.Ordered]` es una **función genérica**. El parámetro de tipo `T` está restringido por `constraints.Ordered`, que incluye todos los tipos que soportan `<` y `>` (strings, ints, floats, etc.).

Una sola función sirve para ordenar por nombre (`string`), por tamaño (`int64`) y por tiempo (`int64` del timestamp Unix). Sin genéricos, se necesitarían tres funciones separadas.

#### Prioridad de ordenamiento

```go
if !*hasOrderBySize || !*hasOrderByTime {
    orderByName(fs, *hasOrderReverse)
}
if *hasOrderBySize && !*hasOrderByTime {
    orderBySize(fs, *hasOrderReverse)
}
if *hasOrderByTime {
    orderByTime(fs, *hasOrderReverse)
}
```

Prioridad: **tiempo > tamaño > nombre**. `orderByName` se ejecuta casi siempre como paso base, y luego se re-ordena si se pidió. Como `SliceStable` mantiene el orden previo para elementos iguales, archivos con el mismo tamaño quedan ordenados por nombre.

### Paso 5: Impresión

```go
func printList(fs []file, nRecords int) {
    for _, file := range fs[:nRecords] {
        style := mapStyleByFileType[file.fileType]
        fmt.Printf("%s %s %s %10d %s %s %s %s %s\n", ...)
    }
}
```

`fs[:nRecords]` toma los primeros `nRecords` elementos. `%10d` alinea el tamaño a 10 caracteres a la derecha para columnas limpias.

### Detección de archivos especiales

- **`isLink`**: verifica si el modo empieza con `L` (convención Unix para symlinks)
- **`isExecutable`**: en Windows busca `.exe`, en Unix/Mac busca el bit `x` en permisos
- **`isCompress`/`isImage`**: verifican extensiones conocidas
- **`isHidden`**: en Unix un archivo oculto empieza con `.`, en Windows usa atributos del sistema

`runtime.GOOS` es una constante que el compilador establece según la plataforma, permitiendo código multiplataforma.

---

## Dependencias externas

| Paquete | Para qué sirve |
|---------|----------------|
| `github.com/fatih/color` | Colores ANSI en terminal |
| `github.com/AJRDRGZ/fileinfo` | Obtener usuario/grupo del archivo y detectar archivos ocultos (multiplataforma) |
| `golang.org/x/exp/constraints` | El tipo `constraints.Ordered` para genéricos |

---

## Flujo completo

```
CLI args → flag.Parse() → compilar regex (si -p)
    → os.ReadDir(dirPath)
    → filtrar ocultos → filtrar por regex compilada
    → clasificar tipo (setFile)
    → ordenar (nombre → tamaño → tiempo)
    → limitar N registros
    → imprimir con colores e iconos
    → errores a stderr con os.Exit(1)
```

## Uso

```bash
# Listar directorio actual
go run .

# Listar con archivos ocultos
go run . -a

# Filtrar por patrón
go run . -p "\.go$"

# Ordenar por tamaño, invertido, solo 5 resultados
go run . -s -r -n 5

# Ordenar por tiempo de modificación
go run . -t /ruta/directorio
```
