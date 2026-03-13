# Módulo 02 — Punteros, Colecciones y Estructuras

Este módulo cubre los tipos compuestos de Go: desde punteros y manejo de memoria hasta arrays, slices, maps y structs.

---

## Contenido

| # | Tema | Carpeta |
|---|------|---------|
| 1 | [Punteros](#1-punteros) | `1_punteros/` |
| 2 | [Arrays](#2-arrays) | `2_arrays/` |
| 3 | [Slices (Parte 1)](#3-slices---referencia-a-arrays) | `3_slices/` |
| 4 | [Slices (Parte 2)](#4-slices---length-capacity-y-append) | `4_slices_2/` |
| 5 | [Maps](#5-maps) | `5_maps/` |
| 6 | [Estructuras](#6-estructuras-structs) | `6_estructuras/` |

---

## 1. Punteros

**Archivo:** [`1_punteros/main.go`](1_punteros/main.go)

### Concepto

Un puntero es una variable que almacena la **dirección de memoria** de otra variable, no su valor directamente.

```
Variable:  color = "🟢"       → almacena el valor
Puntero:   ptr   = 0xc0000... → almacena la dirección donde está "🟢"
```

### Operadores clave

| Operador | Nombre | Uso | Ejemplo |
|----------|--------|-----|---------|
| `&` | Address-of | Obtiene la dirección de una variable | `ptr := &color` |
| `*` | Dereference | Accede al valor en la dirección | `fmt.Println(*ptr)` → `"🟢"` |
| `*T` | Tipo puntero | Declara un puntero al tipo T | `var ptr *string` |

### Cómo funciona

```go
color := "🟢"

var pointerColor *string   // declara puntero a string (valor cero: nil)
pointerColor = &color      // apunta a la dirección de color

fmt.Println(*pointerColor) // "🟢" — lee el valor a través del puntero
*pointerColor = "🟦"       // modifica el valor original a través del puntero
fmt.Println(color)          // "🟦" — el original cambió
```

### Zero value de un puntero

Un puntero sin inicializar tiene valor `nil`. Intentar dereferenciar un puntero `nil` causa un **panic** en tiempo de ejecución.

```go
var ptr *string       // ptr == nil
fmt.Println(*ptr)     // PANIC: nil pointer dereference
```

### Casos de uso en el mundo real

- **Modificar valores en funciones:** Go pasa argumentos por valor. Si quieres que una función modifique una variable externa, pasas un puntero:
  ```go
  func updateStatus(status *string) {
      *status = "completed"
  }
  ```
- **Evitar copias costosas:** pasar structs grandes por puntero en lugar de copiar toda la estructura:
  ```go
  func processOrder(order *Order) { ... }  // eficiente
  func processOrder(order Order) { ... }   // copia todo el struct
  ```
- **Valores opcionales (nil como ausencia):** un puntero puede ser `nil`, indicando "sin valor":
  ```go
  type Config struct {
      Timeout *int  // nil = usar valor por defecto
  }
  ```
- **Receptores de métodos:** en Go, los métodos que modifican el struct usan receptor de puntero:
  ```go
  func (u *User) SetName(name string) {
      u.Name = name
  }
  ```

### Errores comunes

| Error | Causa | Solución |
|-------|-------|----------|
| `nil pointer dereference` | Dereferenciar un puntero `nil` | Verificar `if ptr != nil` antes de usar `*ptr` |
| Mutation inesperada | Múltiples punteros al mismo dato | Ser consciente de quién tiene acceso al puntero |

---

## 2. Arrays

**Archivo:** [`2_arrays/main.go`](2_arrays/main.go)

### Concepto

Un array es una colección de elementos del **mismo tipo** con un **tamaño fijo** definido en tiempo de compilación.

```go
var flags [3]string                          // array de 3 strings (zero value: ["", "", ""])
flags = [3]string{"🇻🇦", "🇨🇽", "🇸🇭"}          // inicialización explícita
autoFlags := [...]string{"🇻🇦", "🇨🇽", "🇸🇭"}   // el compilador infiere el tamaño (3)
```

### Características clave

- **Tamaño fijo:** `[3]string` y `[4]string` son **tipos diferentes** e incompatibles.
- **Valor, no referencia:** asignar un array a otra variable **copia** todos los elementos.
- **Índice base 0:** el primer elemento está en la posición `0`.

### Arrays vs Slices

| Característica | Array | Slice |
|----------------|-------|-------|
| Tamaño | Fijo (parte del tipo) | Dinámico |
| Asignación | Copia completa | Copia la referencia |
| Uso en la práctica | Poco común | Muy común |

### Casos de uso en el mundo real

Los arrays puros se usan poco en Go (se prefieren slices), pero tienen su lugar:

- **Buffers de tamaño fijo:** `[256]byte` para leer bloques de datos.
- **Claves criptográficas:** `[32]byte` para SHA-256 hashes.
- **Coordenadas:** `[3]float64` para puntos en 3D (x, y, z).
- **Días de la semana:** `[7]string` — cuando sabes que siempre serán exactamente 7.
- **Direcciones IP:** `[4]byte` para IPv4.

---

## 3. Slices — Referencia a arrays

**Archivo:** [`3_slices/main.go`](3_slices/main.go)

### Concepto

Un slice es una **vista flexible** sobre un array subyacente. No almacena datos directamente — **referencia** un segmento de un array.

```
Array:   [🍕, 🍔, 🌮, 🍣, 🍜]
Slice 1:  ^---------^            things[:3] → [🍕, 🍔, 🌮]
Slice 2:              ^------^   things[3:] → [🍣, 🍜]
```

### Sintaxis de slicing

| Expresión | Resultado | Descripción |
|-----------|-----------|-------------|
| `a[:]` | Todo el array | Desde el inicio hasta el final |
| `a[:3]` | Primeros 3 | Desde el inicio hasta índice 3 (excluido) |
| `a[3:]` | Desde el 4to | Desde índice 3 hasta el final |
| `a[1:4]` | Del 2do al 4to | Desde índice 1 hasta índice 4 (excluido) |

### Dato clave: los slices comparten memoria

Modificar un elemento del slice **modifica el array original** (y cualquier otro slice que apunte al mismo segmento):

```go
things := [5]string{"🍕", "🍔", "🌮", "🍣", "🍜"}
slice := things[:3]    // [🍕, 🍔, 🌮]
slice[0] = "🥗"        // modifica things[0] también
fmt.Println(things[0]) // "🥗" — el array original cambió
```

### Casos de uso en el mundo real

- **Procesamiento de lotes:** dividir una lista de items en chunks para procesamiento paralelo:
  ```go
  batch1 := items[:batchSize]
  batch2 := items[batchSize:]
  ```
- **Parsing de protocolos:** extraer headers y body de un paquete de red:
  ```go
  header := packet[:headerLen]
  body := packet[headerLen:]
  ```
- **Ventana deslizante:** algoritmos que analizan subconjuntos de datos (últimos N elementos).

---

## 4. Slices — Length, Capacity y Append

**Archivo:** [`4_slices_2/main.go`](4_slices_2/main.go)

### Concepto

Un slice tiene tres propiedades internas:

| Propiedad | Función | Obtener |
|-----------|---------|---------|
| **Pointer** | Dirección del primer elemento | — |
| **Length** | Cantidad de elementos en el slice | `len(s)` |
| **Capacity** | Cantidad de elementos en el array subyacente (desde el inicio del slice) | `cap(s)` |

### Crear slices con `make`

```go
s := make([]string, 0, 3)  // length=0, capacity=3
// Reserva espacio para 3 elementos pero está vacío
```

### Agregar elementos con `append`

```go
s := make([]string, 0, 3)
s = append(s, "🍕")        // len=1, cap=3
s = append(s, "🍔", "🌮")  // len=3, cap=3
s = append(s, "🍣")        // len=4, cap=6 — ¡se creó un nuevo array subyacente!
```

> **Importante:** cuando `append` excede la capacidad, Go crea un **nuevo array subyacente** (generalmente el doble de capacidad) y copia los datos. El slice anterior ya no comparte memoria con el nuevo.

### Visualización del crecimiento

```
make([]T, 0, 3):  [_, _, _]          len=0, cap=3
append "A":       [A, _, _]          len=1, cap=3
append "B","C":   [A, B, C]          len=3, cap=3
append "D":       [A, B, C, D, _, _] len=4, cap=6  ← nuevo array
```

### Casos de uso en el mundo real

- **Coleccionar resultados:** acumular resultados de una query o API paginada:
  ```go
  var results []Item
  for page := range pages {
      items := fetchPage(page)
      results = append(results, items...)
  }
  ```
- **Pre-alocar para rendimiento:** si sabes cuántos elementos habrá, usar `make` evita reasignaciones:
  ```go
  users := make([]User, 0, len(ids))  // evita múltiples allocations
  for _, id := range ids {
      users = append(users, fetchUser(id))
  }
  ```
- **Queues y stacks:** implementar colas y pilas con `append` y slicing:
  ```go
  // Stack (LIFO)
  stack = append(stack, item)       // push
  item, stack = stack[len(stack)-1], stack[:len(stack)-1]  // pop
  ```

---

## 5. Maps

**Archivo:** [`5_maps/main.go`](5_maps/main.go)

### Concepto

Un map es una colección de pares **clave-valor** (similar a diccionarios en Python o HashMap en Java). Las claves deben ser de un tipo comparable.

### Crear maps

```go
// Con make
music := make(map[string]string)

// Con literal
music := map[string]string{
    "guitar": "🎸",
    "drum":   "🥁",
}
```

### Operaciones fundamentales

| Operación | Sintaxis | Ejemplo |
|-----------|----------|---------|
| Insertar/Actualizar | `m[key] = value` | `music["piano"] = "🎹"` |
| Leer | `value := m[key]` | `v := music["guitar"]` |
| Eliminar | `delete(m, key)` | `delete(music, "drum")` |
| Verificar existencia | `v, ok := m[key]` | Si `ok == false`, la clave no existe |

### El patrón "comma ok"

Acceder a una clave que no existe retorna el **zero value** del tipo (no un error). Para distinguir entre "la clave no existe" y "la clave existe con zero value", usa el patrón comma-ok:

```go
content, ok := music["fake"]
if !ok {
    fmt.Println("La clave no existe")
}
// content == "" (zero value de string), ok == false
```

### Casos de uso en el mundo real

- **Cache en memoria:** almacenar resultados de operaciones costosas:
  ```go
  cache := make(map[string]*Response)
  if cached, ok := cache[url]; ok {
      return cached  // hit
  }
  ```
- **Contadores y frecuencias:** contar ocurrencias de palabras, eventos, errores:
  ```go
  wordCount := make(map[string]int)
  for _, word := range words {
      wordCount[word]++
  }
  ```
- **Lookup tables:** mapear códigos a nombres, IDs a objetos:
  ```go
  statusNames := map[int]string{
      200: "OK",
      404: "Not Found",
      500: "Internal Server Error",
  }
  ```
- **Agrupación de datos:** agrupar items por categoría:
  ```go
  byCategory := make(map[string][]Product)
  for _, p := range products {
      byCategory[p.Category] = append(byCategory[p.Category], p)
  }
  ```
- **Conjuntos (sets):** Go no tiene un tipo set nativo, pero se simula con `map[T]bool`:
  ```go
  visited := make(map[string]bool)
  visited["node-1"] = true
  if visited["node-1"] { ... }
  ```

---

## 6. Estructuras (Structs)

**Archivo:** [`6_estructuras/main.go`](6_estructuras/main.go)

### Concepto

Un struct es un tipo compuesto que agrupa campos con nombre bajo un mismo tipo. Es la forma de Go de modelar entidades del mundo real (similar a clases en otros lenguajes, pero sin herencia).

### Definición e inicialización

```go
// Definición
type Person struct {
    Name        string
    Age         uint8
    HasChildren bool
}

// Inicialización con campos nombrados (recomendado)
camilo := Person{
    Name:        "Camilo",
    Age:         29,
    HasChildren: false,
}

// Inicialización posicional (frágil si el struct cambia)
andrea := Person{"Andrea", 31, true}
```

### Acceso a campos

```go
fmt.Println(camilo.Name)  // "Camilo"
camilo.Age = 30            // modificar campo
```

### Formato de impresión

```go
fmt.Printf("%v\n", camilo)   // {Camilo 29 false}
fmt.Printf("%+v\n", camilo)  // {Name:Camilo Age:29 HasChildren:false}
```

### Casos de uso en el mundo real

- **Modelos de dominio:** representar entidades de negocio:
  ```go
  type Order struct {
      ID        string
      UserID    string
      Items     []OrderItem
      Total     float64
      Status    string
      CreatedAt time.Time
  }
  ```
- **DTOs (Data Transfer Objects):** estructurar datos de APIs:
  ```go
  type APIResponse struct {
      Data    json.RawMessage `json:"data"`
      Error   string          `json:"error,omitempty"`
      Status  int             `json:"status"`
  }
  ```
- **Configuración:** agrupar parámetros de configuración:
  ```go
  type DatabaseConfig struct {
      Host     string
      Port     int
      User     string
      Password string
      DBName   string
  }
  ```
- **Eventos:** modelar eventos en un sistema event-driven:
  ```go
  type PaymentEvent struct {
      OrderID   string
      Amount    float64
      Currency  string
      Timestamp time.Time
  }
  ```

### Struct tags

Aunque no se ven en el ejemplo del curso, las struct tags son metadatos que controlan serialización:

```go
type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"email"`
}
```
