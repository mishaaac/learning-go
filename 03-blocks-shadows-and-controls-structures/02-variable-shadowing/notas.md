# Shadowing de variables en Go

El shadowing ocurre cuando una declaración interior usa el mismo nombre que un identificador de un scope contenedor. El identificador exterior sigue existiendo, pero ese nombre se resuelve como la declaración interior hasta que termina el scope interno.

## 1. Qué debo entender

- El shadowing crea un identificador nuevo; no modifica el exterior.
- `:=` puede hacer shadowing accidental de un nombre de un bloque contenedor.
- En una short declaration, los nombres existentes solo se reutilizan cuando fueron declarados en el mismo bloque.
- Las declaraciones locales pueden ocultar nombres de packages importados e identificadores predeclarados.
- El shadowing es legal, pero puede volver engañoso el código y ocultar bugs.

## 2. Conceptos clave

| Situación | Resultado |
|---|---|
| `x := 5` interior con una `x` exterior | La nueva `x` interior oculta a la exterior |
| `x, y := ...` en el mismo bloque que `x` | Se reutiliza `x`; se declara la nueva `y` |
| `x, y := ...` en un bloque interior | Tanto `x` como `y` pueden ser nuevas |
| `fmt := "text"` | El nombre local oculta al package importado `fmt` |
| `true := 10` | El nombre local oculta al `true` predeclarado |

## 3. Cómo funciona en Go

```go
x := 10

if x > 5 {
    fmt.Println(x) // 10
    x := 5
    fmt.Println(x) // 5
}

fmt.Println(x) // 10
```

Con una short declaration múltiple en el bloque interior:

```go
x := 10
if x > 5 {
    x, y := 5, 20
    fmt.Println(x, y) // 5 20
}
fmt.Println(x) // 10
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `x = 5` asigna a una variable existente visible; `x := 5` declara una variable en el bloque actual.

Los nombres predeclarados como `int`, `string`, `true`, `make` y `nil` son identificadores en vez de keywords, por lo que Go permite hacerles shadowing. Hacerlo provoca que nombres conocidos del lenguaje signifiquen algo inesperado.

> ⚠️ **Corrección importante:** la regla de “al menos una variable nueva” de `:=` se aplica dentro del bloque actual. Un nombre coincidente que solo se encuentra en un bloque contenedor no cuenta como variable existente para esa short declaration.

Evita ocultar nombres de packages o valores cuya versión exterior todavía se necesite posteriormente en el mismo scope.

## 5. Chuleta rápida

- Mismo nombre en un scope interior → shadowing
- El identificador exterior sigue existiendo
- El nombre interior prevalece hasta que termina su scope
- `=` → asignación
- `:=` → short declaration
- `:=` solo reutiliza nombres del bloque actual
- Al menos un nombre de la izquierda debe ser nuevo en ese bloque
- Los nombres de packages pueden sufrir shadowing
- Los identificadores predeclarados pueden sufrir shadowing
- Que sea legal no significa que sea claro o recomendable

### Comprueba que realmente lo sabes

1. ¿Por qué la `x` exterior permanece en `10` después de una `x := 5` interior?
2. ¿Cómo afecta la regla del bloque actual a `x, y := ...` dentro de un `if`?
3. ¿Por qué resulta especialmente confuso ocultar el nombre de un package importado?
