# Comparing and Converting Structs in Go

Struct equality depends on the comparability of every field, while conversion depends on the relationship between the source and target struct types. Named and anonymous structs follow different assignability rules even when their visible fields look the same.

## 1. What I Need to Understand

- A struct supports `==` and `!=` only when every field is comparable.
- Equality compares corresponding fields.
- Independently defined named struct types remain different types.
- Compatible named struct types may require an explicit conversion.
- A named and an unnamed struct may be directly assignable when their underlying types are identical and the assignability rules are satisfied.

## 2. Key Concepts

| Situation | Direct assignment/comparison |
|---|---|
| Same comparable named type | Allowed |
| Struct contains slice, map, or function | Equality not allowed |
| Two distinct named types with matching structure | Explicit conversion required |
| Named and anonymous types with identical underlying structure | Direct assignment may be allowed |
| Custom equality semantics | Write a function; `==` cannot be redefined |

## 3. How It Works in Go

```go
type FirstPerson struct {
    Name string
    Age  int
}

type SecondPerson struct {
    Name string
    Age  int
}

first := FirstPerson{Name: "Bob", Age: 50}
second := SecondPerson(first) // explicit conversion

var anonymous struct {
    Name string
    Age  int
}
anonymous = first

fmt.Println(first == anonymous) // true
fmt.Println(second)
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** two named structs with matching fields are still distinct named types; structural similarity alone does not permit direct assignment or comparison.

Field names, order, types, and embedded status determine structural identity; tags also matter for type identity and direct assignability. For explicit struct conversion, Go permits otherwise identical underlying struct types while ignoring tags.

> ⚠️ **Important correction:** channels are comparable, so a channel field does not by itself make a struct non-comparable. Slices, maps, and functions are the common field types that prevent struct comparison.

Go does not offer a method that overrides `==`. If equality should ignore or transform fields, express that rule in a separate function.

## 5. Quick Cheat Sheet

- Struct comparable → every field comparable
- Equality → compare corresponding fields
- Slice/map/function field → struct not comparable
- Channel field → still comparable
- Distinct named types remain distinct
- Matching named structs → explicit conversion may be possible
- Conversion needs matching field structure and types
- Struct tags are ignored for explicit conversion
- Named ↔ compatible anonymous may assign directly
- `==` cannot be customized
- Custom semantics → write a comparison function

### Test Yourself

1. Why does a slice field prevent struct comparison while a channel field does not?
2. Why is an explicit conversion needed between two separately named but structurally matching structs?
3. How do struct tags affect type identity versus explicit conversion?
