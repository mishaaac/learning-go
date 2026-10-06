# Structs in Go

A struct groups a fixed set of related fields, and each field may have its own type. It expresses data with a known shape more clearly than a map of dynamic keys.

## 1. What I Need to Understand

- A struct type declares field names and their types.
- The zero value of a struct contains the zero value of every field.
- Fields are read and written with dot notation.
- Struct literals may be positional or keyed by field name.
- Keyed literals are usually clearer, allow omitted fields, and resist field-order changes.

## 2. Key Concepts

| Concept | Meaning |
|---|---|
| `type Person struct { ... }` | Defines a named struct type |
| `var person Person` | Zero-valued struct |
| `Person{}` | Struct literal with all fields at zero values |
| `Person{"Ana", 30, "cat"}` | Positional literal; all fields and order matter |
| `Person{Name: "Ana"}` | Keyed literal; omitted fields keep zero values |

## 3. How It Works in Go

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

A type declared inside a function or block is limited to that scope; a package-level declaration can be reused throughout the package.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a map has dynamic keys of one key type and one value type; a struct has a fixed field set whose fields may have different types.

A positional literal must provide every field in declaration order. A keyed literal can change order and omit fields:

```go
person := Person{Age: 30, Name: "Beth"}
```

Do not mix positional and keyed elements in the same struct literal. Prefer keyed fields except for very small, stable, obvious structures.

> ⚠️ **Important correction:** a Go struct is not a class. Go can define methods on named types, but it does not use a traditional class-inheritance model.

## 5. Quick Cheat Sheet

- Struct → fixed group of related fields
- Fields may have different types
- Define → `type Person struct { ... }`
- Zero value → every field has its own zero value
- Empty literal → `Person{}`
- Read/write field → `person.Name`
- Positional literal → all fields, declaration order
- Keyed literal → `Field: value`
- Omitted keyed fields → zero values
- Do not mix literal styles
- Prefer keyed literals for clarity

### Test Yourself

1. Why can a struct model mixed data more precisely than `map[string]string`?
2. What happens to fields omitted from a keyed struct literal?
3. Why are keyed literals usually easier to maintain than positional literals?
