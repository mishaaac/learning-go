# Structs anónimos en Go

Un struct anónimo usa directamente un literal de tipo struct en vez de introducir un nombre de tipo reutilizable. Es útil para formas de datos locales y puntuales, así como para pequeñas tablas de valores con la misma estructura.

## 1. Qué debo entender

- `struct { ... }` puede usarse directamente como tipo sin una declaración `type`.
- La variable tiene un tipo struct concreto aunque ese tipo no tenga un nombre declarado.
- Los campos siguen las mismas reglas de zero values y dot notation que los structs con nombre.
- Un literal de struct anónimo define el tipo e inicializa un valor a la vez.
- Sus usos habituales incluyen formas temporales para codificación y table-driven tests.

## 2. Conceptos clave

| Forma | Propósito |
|---|---|
| `var value struct { Name string }` | Declara un struct anónimo con zero value |
| `struct { Name string }{Name: "Fido"}` | Define e inicializa inmediatamente |
| `[]struct { Input int; Want int }` | Colección de structs anónimos para casos de prueba |
| Struct con nombre | Mejor cuando el tipo debe reutilizarse o tener un nombre de dominio |

## 3. Cómo funciona en Go

```go
pet := struct {
    Name string
    Kind string
}{
    Name: "Fido",
    Kind: "dog",
}

fmt.Println(pet.Name) // Fido
```

Una pequeña estructura para table-driven tests puede escribirse así:

```go
tests := []struct {
    Input int
    Want  int
}{
    {Input: 2, Want: 4},
    {Input: 3, Want: 6},
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** en `var person struct { Name string }`, `person` es el nombre de la variable, no del tipo.

Anónimo no significa dinámico: los campos y sus tipos siguen conociéndose durante la compilación. Repetir el mismo anonymous struct largo en varios lugares reduce la claridad; introduce un tipo con nombre cuando la forma represente un concepto reutilizable.

Los structs anónimos temporales pueden ayudar a hacer marshaling o unmarshaling de una representación externa pequeña sin crear un tipo de dominio permanente. Las reglas de codificación dependen del codificador utilizado y están separadas de la sintaxis del struct anónimo.

## 5. Chuleta rápida

- Struct anónimo → tipo struct sin nombre declarado
- Literal de tipo → `struct { ... }`
- Los campos siguen siendo estáticamente tipados
- Las reglas de zero value no cambian
- El acceso a campos sigue usando `.`
- Definir e inicializar → `struct { ... }{ ... }`
- Forma local de un solo uso → buen candidato
- Concepto de dominio reutilizado → prefiere un tipo con nombre
- Los table-driven tests suelen usar `[]struct { ... }`
- Las formas temporales para marshaling son otro uso habitual

### Comprueba que realmente lo sabes

1. ¿Por qué un anonymous struct sigue estando tipado estáticamente?
2. ¿Cuándo debería una forma anónima repetida convertirse en un tipo con nombre?
3. ¿Cómo ayuda `[]struct { ... }` a crear table-driven tests?
