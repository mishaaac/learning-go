# Variable Shadowing in Go

Shadowing occurs when an inner declaration uses the same name as an identifier from an enclosing scope. The outer identifier still exists, but that name resolves to the inner declaration until the inner scope ends.

## 1. What I Need to Understand

- Shadowing creates a new identifier; it does not modify the outer one.
- `:=` can accidentally shadow a name from an enclosing block.
- In a short declaration, existing names are reused only when they were declared in the same block.
- Local declarations can shadow imported package names and predeclared identifiers.
- Shadowing is legal, but it can make code misleading and hide bugs.

## 2. Key Concepts

| Situation | Result |
|---|---|
| Inner `x := 5` with outer `x` | New inner `x` shadows outer `x` |
| `x, y := ...` in the same block as `x` | Existing `x` reused; new `y` declared |
| `x, y := ...` in an inner block | Inner `x` and `y` may both be new |
| `fmt := "text"` | Local name hides imported package `fmt` |
| `true := 10` | Local name hides the predeclared `true` |

## 3. How It Works in Go

```go
x := 10

if x > 5 {
    fmt.Println(x) // 10
    x := 5
    fmt.Println(x) // 5
}

fmt.Println(x) // 10
```

With a multiple short declaration in the inner block:

```go
x := 10
if x > 5 {
    x, y := 5, 20
    fmt.Println(x, y) // 5 20
}
fmt.Println(x) // 10
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `x = 5` assigns to a visible existing variable; `x := 5` declares a variable in the current block.

Predeclared names such as `int`, `string`, `true`, `make`, and `nil` are identifiers rather than keywords, so Go allows them to be shadowed. Doing so makes familiar language names mean something unexpected.

> ⚠️ **Important correction:** the “at least one new variable” rule for `:=` applies within the current block. A matching name found only in an enclosing block does not count as an existing variable for that short declaration.

Avoid shadowing package names or values whose outer version is still needed later in the same scope.

## 5. Quick Cheat Sheet

- Same name in inner scope → shadowing
- Outer identifier still exists
- Inner name wins until its scope ends
- `=` → assignment
- `:=` → short declaration
- `:=` reuses names only from the current block
- At least one left-side name must be new in that block
- Package names can be shadowed
- Predeclared identifiers can be shadowed
- Legal does not mean clear or advisable

### Test Yourself

1. Why does the outer `x` remain `10` after an inner `x := 5`?
2. How does the current-block rule affect `x, y := ...` inside an `if`?
3. Why is shadowing an imported package name especially confusing?
