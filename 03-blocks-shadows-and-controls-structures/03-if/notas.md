# `if` en Go

El `if` de Go selecciona una rama según una condición booleana y no necesita paréntesis alrededor de esa condición. Una sentencia inicial opcional puede crear valores cuyo scope cubre toda la cadena `if`/`else`.

## 1. Qué debo entender

- La condición de un `if` debe producir un `bool`.
- Go escribe la condición sin paréntesis alrededor.
- `else if` y `else` proporcionan ramas alternativas mutuamente excluyentes.
- Cada rama es su propio bloque.
- `if initialStatement; condition` limita los valores declarados a toda la estructura condicional.

## 2. Conceptos clave

| Forma | Scope de una variable declarada |
|---|---|
| Declarada antes del `if` | Disponible antes, dentro y después de la sentencia |
| Declarada dentro de una rama | Solo esa rama |
| Declarada en la sentencia inicial | Condición y todas las ramas `if`/`else if`/`else` |

La sentencia inicial es una simple statement de Go, normalmente una short variable declaration.

## 3. Cómo funciona en Go

```go
if number := rand.Intn(10); number == 0 {
    fmt.Println("too low")
} else if number > 5 {
    fmt.Println("too high", number)
} else {
    fmt.Println("good", number)
}

// number está fuera de ámbito aquí.
```

La ejecución selecciona únicamente la primera rama cuya condición se cumpla; `else` maneja el caso restante.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** una variable declarada dentro de la primera rama no está disponible en `else`; una variable declarada antes de la condición está disponible en todas las ramas.

La sentencia inicial y la condición se separan mediante un punto y coma:

```go
if value, ok := lookup(); ok {
    fmt.Println(value)
}
```

Aunque otras simple statements son legales antes de la condición, una declaración que proporciona la condición suele ser el uso más claro.

La declaración inicial puede hacer shadowing de una variable exterior con el mismo nombre. Declara el valor antes del `if` cuando deba seguir disponible después.

## 5. Chuleta rápida

- Forma → `if condition { ... }`
- La condición debe ser booleana
- No se necesitan paréntesis en la condición
- Alternativas → `else if` y después `else` opcional
- Cada rama tiene su propio scope
- Sentencia inicial → `if statement; condition`
- La variable inicial es visible en la condición
- También es visible en todas las ramas
- Queda fuera de scope después de la sentencia completa
- La declaración inicial puede ocultar un nombre exterior

### Comprueba que realmente lo sabes

1. ¿Por qué una variable de la sentencia inicial puede usarse en `else`, pero no después de todo el `if`?
2. ¿Cuándo debería declararse un valor antes del `if` en vez de en su sentencia inicial?
3. ¿Cómo puede una short declaration inicial causar shadowing accidentalmente?
