# Structs en Go

Un struct agrupa un conjunto fijo de campos relacionados y cada campo puede tener su propio tipo. Expresa datos con una forma conocida con mayor claridad que un map de keys dinámicas.

## 1. Qué debo entender

- Un tipo struct declara nombres de campos y sus tipos.
- El zero value de un struct contiene el zero value de cada campo.
- Los campos se leen y escriben mediante dot notation.
- Los struct literals pueden ser posicionales o usar nombres de campos.
- Los literals con nombres suelen ser más claros, permiten omitir campos y resisten cambios de orden.

## 2. Conceptos clave

| Concepto | Significado |
|---|---|
| `type Person struct { ... }` | Define un tipo struct con nombre |
| `var person Person` | Struct con zero value |
| `Person{}` | Struct literal con todos los campos en sus zero values |
| `Person{"Ana", 30, "cat"}` | Literal posicional; importan todos los campos y el orden |
| `Person{Name: "Ana"}` | Literal con nombres; los campos omitidos conservan zero values |

## 3. Cómo funciona en Go

```go
type Person struct {
    Name string
    Age  int
    Pet  string
}

var first Person
first.Name = "Bob"

second := Person{
    Name: "Beth",
    Age:  30,
}

fmt.Println(first.Age)  // 0
fmt.Printf("%q\n", second.Pet) // ""
```

Un tipo declarado dentro de una función o bloque está limitado a ese scope; una declaración en el nivel del package puede reutilizarse en todo el package.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un map tiene keys dinámicas de un solo tipo de key y un solo tipo de value; un struct tiene un conjunto fijo de campos que pueden tener tipos diferentes.

Un literal posicional debe proporcionar todos los campos en el orden de declaración. Un literal con nombres puede cambiar el orden y omitir campos:

```go
person := Person{Age: 30, Name: "Beth"}
```

No mezcles elementos posicionales y con nombres en el mismo struct literal. Prefiere los campos con nombres salvo en estructuras muy pequeñas, estables y evidentes.

> ⚠️ **Corrección importante:** un struct de Go no es una clase. Go puede definir métodos sobre tipos con nombre, pero no utiliza un modelo tradicional de herencia de clases.

## 5. Chuleta rápida

- Struct → grupo fijo de campos relacionados
- Los campos pueden tener tipos diferentes
- Definir → `type Person struct { ... }`
- Zero value → cada campo tiene su propio zero value
- Literal vacío → `Person{}`
- Leer/escribir campo → `person.Name`
- Literal posicional → todos los campos, orden de declaración
- Literal con nombres → `Field: value`
- Campos con nombre omitidos → zero values
- No mezcles estilos de literal
- Prefiere literals con nombres por claridad

### Comprueba que realmente lo sabes

1. ¿Por qué un struct puede modelar datos mixtos con mayor precisión que `map[string]string`?
2. ¿Qué ocurre con los campos omitidos en un struct literal con nombres?
3. ¿Por qué los literals con nombres suelen ser más fáciles de mantener que los posicionales?
