# Anonymous Structs in Go

An anonymous struct uses a struct type literal directly instead of introducing a reusable type name. It is useful for local, one-off data shapes and small tables of uniformly shaped values.

## 1. What I Need to Understand

- `struct { ... }` can be used directly as a type without a `type` declaration.
- The variable has a concrete struct type even though that type has no declared name.
- Fields use the same zero-value and dot-notation rules as named structs.
- An anonymous struct literal defines the type and initializes a value together.
- Common uses include temporary encoding shapes and table-driven tests.

## 2. Key Concepts

| Form | Purpose |
|---|---|
| `var value struct { Name string }` | Declare a zero-valued anonymous struct |
| `struct { Name string }{Name: "Fido"}` | Define and initialize immediately |
| `[]struct { Input int; Want int }` | Collection of anonymous test-case structs |
| Named struct | Better when the type should be reused or have a domain name |

## 3. How It Works in Go

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

A small table-driven-test shape can be written as:

```go
tests := []struct {
    Input int
    Want  int
}{
    {Input: 2, Want: 4},
    {Input: 3, Want: 6},
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** in `var person struct { Name string }`, `person` names the variable, not the type.

Anonymous does not mean dynamic: the fields and their types are still known at compile time. Repeating the same long anonymous struct in several places reduces clarity; introduce a named type when the shape represents a reusable concept.

Temporary anonymous structs can help marshal or unmarshal a small external representation without creating a permanent domain type. The encoding rules themselves depend on the encoder being used and are separate from anonymous-struct syntax.

## 5. Quick Cheat Sheet

- Anonymous struct → struct type without a declared name
- Type literal → `struct { ... }`
- Fields remain statically typed
- Zero-value rules are unchanged
- Field access still uses `.`
- Define and initialize → `struct { ... }{ ... }`
- One-off local shape → good candidate
- Reused domain concept → prefer a named type
- Table-driven tests often use `[]struct { ... }`
- Temporary marshaling shapes are another common use

### Test Yourself

1. Why is an anonymous struct still statically typed?
2. When should a repeated anonymous shape become a named type?
3. How does `[]struct { ... }` support table-driven tests?
