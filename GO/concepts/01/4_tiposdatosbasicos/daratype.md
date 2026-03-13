# Tipos de datos básicos en Go

## Enteros sin signo (`uint`)

| Tipo   | Bits | Rango                                      |
|--------|------|--------------------------------------------|
| `uint8`  | 8   | 0 a 255                                    |
| `uint16` | 16  | 0 a 65.535                                 |
| `uint32` | 32  | 0 a 4.294.967.295                          |
| `uint64` | 64  | 0 a 18.446.744.073.709.551.615             |
| `uint`   | 32/64 | Depende de la arquitectura (mismo tamaño que `uintptr`) |

## Enteros con signo (`int`)

| Tipo   | Bits | Rango                                      |
|--------|------|--------------------------------------------|
| `int8`  | 8   | -128 a 127                                 |
| `int16` | 16  | -32.768 a 32.767                           |
| `int32` | 32  | -2.147.483.648 a 2.147.483.647             |
| `int64` | 64  | -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807 |
| `int`   | 32/64 | Depende de la arquitectura (tipo más usado) |

## Aliases útiles

| Tipo  | Equivalente | Uso principal                          |
|-------|-------------|----------------------------------------|
| `byte` | `uint8`    | Datos binarios, slices de bytes        |
| `rune` | `int32`    | Puntos de código Unicode (caracteres)  |

## Punto flotante

| Tipo     | Bits | Precisión aproximada |
|----------|------|----------------------|
| `float32` | 32  | ~6-7 dígitos decimales |
| `float64` | 64  | ~15-16 dígitos decimales (recomendado) |

## Números complejos

| Tipo        | Componentes      |
|-------------|------------------|
| `complex64`  | `float32` + `float32` (parte real e imaginaria) |
| `complex128` | `float64` + `float64` |

## Otros tipos básicos

| Tipo    | Descripción                    |
|---------|--------------------------------|
| `bool`  | `true` o `false`               |
| `string`| Secuencia inmutable de bytes (UTF-8) |

---

## Notas prácticas

- **`int`** y **`uint`**: Usar por defecto para enteros; el compilador elige el tamaño según la plataforma.
- **`float64`**: Preferir sobre `float32` salvo que el ahorro de memoria sea crítico.
- **`byte`**: Ideal para I/O, redes, archivos binarios.
- **`rune`**: Para iterar sobre caracteres Unicode en strings (`for _, r := range "Hola"`).
