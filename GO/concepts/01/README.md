# Módulo 01 — Fundamentos de Go

Este módulo cubre los conceptos básicos del lenguaje Go: desde la declaración de variables hasta los operadores lógicos y aritméticos.

---

## Contenido

| # | Tema | Carpeta |
|---|------|---------|
| 1 | [Variables](#1-variables) | `1_variables/` |
| 2 | [Constantes](#2-constantes) | `2_constantes/` |
| 3 | [Comentarios](#3-comentarios) | `3_comentarios/` |
| 4 | [Tipos de datos básicos](#4-tipos-de-datos-básicos) | `4_tiposdatosbasicos/` |
| 5 | [Operadores aritméticos](#5-operadores-aritméticos) | `5_operadoresaritmeticos/` |
| 6 | [Operadores lógicos](#6-operadores-lógicos) | `6_operadoreslogicos/` |

---

## 1. Variables

**Archivo:** [`1_variables/main.go`](1_variables/main.go)

### Concepto

Las variables son espacios en memoria que almacenan valores. En Go existen dos formas principales de declararlas:

- **Declaración explícita:** `var nombre string = "valor"`
- **Declaración corta:** `nombre := "valor"` (infiere el tipo automáticamente)

Go también permite la **asignación múltiple** en una sola línea:

```go
apple, banana, orange := "apple", "banana", "orange"
```

### Reglas importantes

- Toda variable declarada **debe usarse**, de lo contrario el compilador lanza un error.
- La declaración corta `:=` solo se puede usar dentro de funciones.
- `var` se puede usar a nivel de paquete (fuera de funciones).

### Casos de uso en el mundo real

- **Configuración de servicios:** almacenar host, puerto, timeouts al iniciar una aplicación.
- **Procesamiento de datos:** guardar valores intermedios en cálculos o transformaciones.
- **Asignación múltiple:** cuando una función retorna múltiples valores (patrón muy común en Go):
  ```go
  result, err := http.Get("https://api.example.com/data")
  ```

---

## 2. Constantes

**Archivo:** [`2_constantes/main.go`](2_constantes/main.go)

### Concepto

Las constantes son valores que **no cambian** durante la ejecución del programa. Se declaran con `const`.

```go
const Pi = 3.14159

const (
    Jan = iota + 1  // 1
    Feb              // 2
    Mar              // 3
)
```

### El poder de `iota`

`iota` es un contador que se incrementa automáticamente dentro de un bloque `const`. Empieza en 0 y se incrementa en 1 por cada línea. Es ideal para crear enumeraciones.

```go
const (
    Admin   = iota  // 0
    Editor          // 1
    Viewer          // 2
)
```

### Casos de uso en el mundo real

- **Códigos de estado:** definir estados de un pedido (Pendiente, Enviado, Entregado).
- **Roles de usuario:** Admin, Editor, Viewer usando `iota`.
- **Configuración inmutable:** valores como tasas de impuesto, límites de reintentos, URLs base de APIs.
- **Días de la semana / meses:** enumeraciones naturales con `iota`.
  ```go
  const (
      Lunes = iota + 1
      Martes
      Miercoles
      // ...
  )
  ```

---

## 3. Comentarios

**Archivo:** [`3_comentarios/main.go`](3_comentarios/main.go)

### Concepto

Los comentarios documentan el código y son ignorados por el compilador.

| Tipo | Sintaxis | Uso |
|------|----------|-----|
| Línea | `// comentario` | Explicaciones breves |
| Bloque | `/* comentario */` | Explicaciones largas o deshabilitar código |

### Comentarios de documentación (godoc)

En Go, los comentarios que preceden a una declaración pública sirven como **documentación oficial**. La herramienta `go doc` los extrae automáticamente.

```go
// Sum returns the sum of two integers.
func Sum(a, b int) int {
    return a + b
}
```

### Casos de uso en el mundo real

- **Documentación de API pública:** cada función, tipo o paquete exportado debería tener un comentario `godoc`.
- **TODO/FIXME:** marcar deuda técnica (`// TODO: optimizar esta query`).
- **Explicar decisiones:** documentar el "por qué" de una decisión, no el "qué" (el código ya dice el qué).
- **Deshabilitar código temporalmente:** usar `/* ... */` para comentar bloques durante debugging.

---

## 4. Tipos de datos básicos

**Archivo:** [`4_tiposdatosbasicos/main.go`](4_tiposdatosbasicos/main.go)

### Concepto

Go es un lenguaje de **tipado estático**: cada variable tiene un tipo definido en tiempo de compilación.

| Categoría | Tipos | Ejemplo |
|-----------|-------|---------|
| Booleanos | `bool` | `true`, `false` |
| Cadenas | `string` | `"hola mundo"` |
| Enteros sin signo | `uint8`, `uint16`, `uint32`, `uint64` | `255`, `65535` |
| Enteros con signo | `int8`, `int16`, `int32`, `int64` | `-128` a `127` (int8) |
| Punto flotante | `float32`, `float64` | `3.14` |
| Complejos | `complex64`, `complex128` | `3+4i` |
| Alias | `byte` (uint8), `rune` (int32) | `'A'`, `'ñ'` |

### Conversión de tipos (type casting)

Go **no convierte tipos implícitamente**. Debes hacerlo de forma explícita:

```go
var a int16 = 50
var b int32 = int32(a)  // conversión explícita
```

### El identificador en blanco `_`

Se usa para **ignorar valores** que no necesitas:

```go
_, err := someFunction()  // ignora el primer retorno
```

### Casos de uso en el mundo real

- **`uint` para IDs y contadores:** nunca serán negativos, `uint64` para IDs de base de datos.
- **`float64` para cálculos financieros:** precios, porcentajes (aunque para precisión extrema se usa `math/big`).
- **`rune` para texto Unicode:** procesar caracteres especiales como emojis o acentos.
- **`byte` para I/O:** leer/escribir archivos, streams de red, encoding.
- **`bool` para feature flags:** activar/desactivar funcionalidades en tiempo de ejecución.

---

## 5. Operadores aritméticos

**Archivo:** [`5_operadoresaritmeticos/main.go`](5_operadoresaritmeticos/main.go)

### Concepto

Los operadores aritméticos realizan cálculos matemáticos sobre valores numéricos.

#### Operadores básicos

| Operador | Operación | Ejemplo |
|----------|-----------|---------|
| `+` | Suma | `5 + 3 = 8` |
| `-` | Resta | `5 - 3 = 2` |
| `*` | Multiplicación | `5 * 3 = 15` |
| `/` | División | `10 / 3 = 3` (entera) |
| `%` | Módulo (residuo) | `10 % 3 = 1` |

#### Operadores de asignación compuesta

| Operador | Equivale a |
|----------|------------|
| `+=` | `x = x + valor` |
| `-=` | `x = x - valor` |
| `*=` | `x = x * valor` |
| `/=` | `x = x / valor` |
| `%=` | `x = x % valor` |

#### Incremento y decremento

```go
x++  // x = x + 1
x--  // x = x - 1
```

> **Nota importante:** En Go, `++` y `--` son **sentencias**, no expresiones. No se puede hacer `y = x++`.

### Precedencia de operadores

Los paréntesis `()` controlan el orden de evaluación, igual que en matemáticas:

```go
result := (2 + 3) * 4  // 20, no 14
```

### Casos de uso en el mundo real

- **Cálculo de precios:** subtotales, impuestos, descuentos (`precio * cantidad * (1 - descuento)`).
- **Paginación:** calcular offset (`(pagina - 1) * tamaño`).
- **Módulo para distribución:** distribuir tareas entre workers (`taskID % numWorkers`).
- **Contadores:** tracking de requests, métricas, reintentos con `++`.

---

## 6. Operadores lógicos

**Archivo:** [`6_operadoreslogicos/main.go`](6_operadoreslogicos/main.go)

### Concepto

Los operadores lógicos y de comparación permiten evaluar condiciones y tomar decisiones.

#### Operadores de comparación

| Operador | Significado | Ejemplo |
|----------|-------------|---------|
| `>` | Mayor que | `5 > 3` → `true` |
| `<` | Menor que | `3 < 5` → `true` |
| `==` | Igual a | `5 == 5` → `true` |
| `!=` | Diferente de | `5 != 3` → `true` |
| `>=` | Mayor o igual | `5 >= 5` → `true` |
| `<=` | Menor o igual | `3 <= 5` → `true` |

#### Operadores lógicos

| Operador | Significado | Comportamiento |
|----------|-------------|----------------|
| `&&` | AND | `true` solo si **ambos** son `true` |
| `\|\|` | OR | `true` si **al menos uno** es `true` |
| `!` | NOT | Invierte el valor booleano |

### Evaluación en cortocircuito

Go evalúa de izquierda a derecha y **se detiene cuando el resultado es seguro**:

- `false && ...` → no evalúa el lado derecho (ya es `false`).
- `true || ...` → no evalúa el lado derecho (ya es `true`).

### Casos de uso en el mundo real

- **Validación de rangos:** `age >= 18 && age <= 65` para verificar si un usuario es adulto en edad laboral.
- **Control de acceso:** `isAdmin || isOwner` para permitir acciones.
- **Validación de formularios:** `email != "" && isValidEmail(email)` para validar campos.
- **Guards en APIs:**
  ```go
  if user == nil || !user.IsActive {
      return errors.New("unauthorized")
  }
  ```
- **Filtros de búsqueda:** combinar múltiples condiciones para filtrar resultados de queries.
