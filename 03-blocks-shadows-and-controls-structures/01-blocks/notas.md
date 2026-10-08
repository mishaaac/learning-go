# Bloques en Go

Los bloques definen dónde son visibles y utilizables los identificadores declarados. Entender los scopes anidados es esencial porque las estructuras de control de Go introducen bloques y las declaraciones internas pueden ocultar nombres externos.

## 1. Qué debo entender

- El universe block contiene los identificadores predeclarados de Go.
- Las declaraciones de nivel de package pertenecen al package block; los nombres de packages importados pertenecen a un file block.
- El cuerpo de una función es un bloque y sus parámetros son visibles en todo ese cuerpo.
- Las llaves explícitas y las estructuras de control introducen bloques anidados.
- Un bloque interior puede usar nombres de bloques contenedores salvo que declare el mismo nombre.

## 2. Conceptos clave

| Bloque | Contenido habitual | Alcance |
|---|---|---|
| Universe | `int`, `true`, `nil`, `make` | Todo el código Go |
| Package | Variables, constantes, tipos y funciones de nivel de package | Package |
| File | Nombres introducidos mediante `import` | Un archivo fuente |
| Cuerpo de función | Parámetros y declaraciones locales | Cuerpo de la función |
| Bloque anidado/de control | Declaraciones dentro de llaves o cláusulas | Ese scope interior |

## 3. Cómo funciona en Go

```go
var packageCount int

func printValue(input int) {
    value := input

    if value > 0 {
        message := "positive"
        fmt.Println(value, message)
    }

    // message está fuera de ámbito aquí.
}
```

```text
universe
└── package
    └── imports del archivo
        └── función
            └── bloque de control
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** la visibilidad fluye desde un bloque contenedor hacia uno interior, no desde un bloque interior hacia fuera.

Las llaves del cuerpo de una función forman un bloque explícito. Go también define bloques implícitos para estructuras como cada `if`, `for` y `switch`, y para las cláusulas de los `switch`.

> ⚠️ **Corrección importante:** los imports no pertenecen al package block. Un nombre importado tiene como scope el archivo que contiene ese import, aunque las declaraciones de nivel de package sean visibles entre archivos del mismo package.

Declarar el mismo nombre en un bloque interior crea un identificador diferente y hace shadowing del exterior durante ese scope.

## 5. Chuleta rápida

- Los bloques controlan el scope de los identificadores
- Universe block → identificadores predeclarados
- Package block → declaraciones de nivel de package
- File block → nombres importados
- Los parámetros son visibles en el cuerpo de la función
- `{}` crea un bloque explícito
- Las estructuras de control también definen bloques implícitos
- El código interior puede ver declaraciones contenedoras
- El código exterior no puede ver declaraciones internas
- Mismo nombre interior → shadowing

### Comprueba que realmente lo sabes

1. ¿Por qué dos archivos de un package pueden usar la misma declaración de nivel de package, pero no automáticamente el mismo nombre importado?
2. ¿En qué dirección funciona la visibilidad entre bloques anidados?
3. ¿Qué ocurre cuando un bloque interior declara un nombre que ya usa un bloque exterior?
