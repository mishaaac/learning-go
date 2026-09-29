# Constants with `const` in Go

`const` gives a name to a value defined by Go's compile-time constant rules. It prevents reassignment, but it does not turn an arbitrary runtime variable or data structure into an immutable value.

## 1. What I Need to Understand

- A `const` declaration associates an identifier with a **constant expression** evaluated at compile time.
- Constants may be declared at package level or inside a function, and related constants can be grouped with `const (...)`.
- A constant may have an explicit type or remain untyped; the detailed distinction is studied separately.
- Go constant values are limited to boolean, numeric, `rune`, and `string` kinds.
- Constant expressions may combine constant operands, permitted operators, conversions, and certain built-in function calls.
- A variable or ordinary function result is not a constant, even when its runtime value seems predictable.
- `const` is not a general mechanism for making arrays, slices, maps, structs, fields, or runtime-computed variables immutable.

> ⚠️ **Important correction:** “The compiler can know the value” is useful but incomplete. A value must also belong to one of Go's permitted constant kinds and be produced by a valid constant expression; a fully written array, slice, map, or struct value still cannot be a constant.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Constant declaration | A declaration beginning with `const` | It gives a name to a value that follows Go's constant rules |
| Constant expression | An expression evaluated at compile time from permitted constant operations | It may contain more than one literal, such as `20 * 10` |
| Typed constant | A constant declared or converted with an explicit constant type | Its value must be representable by that type |
| Untyped constant | A constant without a fixed concrete type yet | It can adapt to a compatible context |
| Runtime value | A value produced from variables or ordinary function execution | It cannot initialize a constant |
| Immutability | The inability to modify a value after creation | Go constants are not a general-purpose immutable-variable feature |

The permitted constant kinds are:

| Kind | Example |
|---|---|
| Boolean | `true` |
| Integer | `10` |
| Floating point | `3.14` |
| Complex | `2 + 3i` |
| `rune` | `'A'` |
| `string` | `"hello"` |

The central comparison is:

| Feature | `const` | `var` |
|---|---|---|
| Package or function scope | Yes | Yes |
| Reassignment | No | Yes |
| Runtime-computed initializer | No | Yes |
| Composite values such as slices or structs | No | Yes |
| Grouped declaration with `(...)` | Yes | Yes |

## 3. How It Works in Go

A constant can have an explicit type or no explicit type:

```go
const maxCount int64 = 10
const greeting = "hello"
const total = 20 * 10
```

`total` is valid because both operands and the multiplication form a constant expression, so its value is computed as `200` during compilation.

Constants are allowed both at package level and inside functions:

```go
const packageLabel = "learning-go"

func localLabel() string {
    const prefix = "item"
    return prefix
}
```

Related constants can be grouped:

```go
const (
    idKey   = "id"
    nameKey = "name"
)
```

Certain predeclared functions can participate in constant expressions under the conditions defined by the language:

```go
const textLength = len("Go")
const arrayCapacity = cap([4]int{})
const complexValue = complex(2, 3)
const realPart = real(complexValue)
const imaginaryPart = imag(complexValue)
```

`textLength` is `2`, `arrayCapacity` is `4`, and the complex-number operations remain constant because their operands meet the constant-expression rules. `len` and `cap` do not produce constants for every possible argument.

Variables break constant expressions:

```go
func calculateTotal() int {
    x := 5
    y := 10
    return x + y
}
```

The result is an `int` value computed when the function runs. Neither `x + y` nor a call to `calculateTotal()` can initialize a `const`.

```text
permitted constant operands
            │
            └── permitted constant operations
                         │
                         └── constant expression → may initialize const

variable or ordinary function call
            │
            └── runtime value → cannot initialize const
```

`iota` is a predeclared integer constant available within constant declarations. Its detailed use is deferred to the later topic on defining related constant values.

## 4. Examples, Differences, and Common Mistakes

| Declaration | Valid? | Reason |
|---|---:|---|
| `const answer = 6 * 7` | ✅ | It uses only constant operands and an allowed operation |
| `const title = "Go"` | ✅ | A string may be a constant |
| `const length = len("Go")` | ✅ | The length of a constant string is constant |
| `const size = len([4]int{})` | ✅ | This array length meets the constant `len` rules |
| `const total = x + y` when `x` and `y` are variables | ❌ | Variable values are not constant operands |
| `const total = calculateTotal()` | ❌ | An ordinary function call is not a constant expression |
| `const items = []int{1, 2}` | ❌ | A slice is not a permitted constant value |
| `const point = struct{ X int }{X: 1}` | ❌ | A struct is not a permitted constant value |

> **Do not confuse:** a constant expression is determined by Go's language rules, not merely by whether a human can predict its result.

A declared constant cannot be assigned another value:

```go
const maxAttempts = 3
// maxAttempts = 4 // cannot assign to a constant
```

The invalid assignment is commented out so that the example remains compilable. Increment and decrement statements such as `maxAttempts++` are invalid for the same reason.

> **Do not confuse:** a named constant is not a variable marked `readonly`. It is a name for a constant value and does not provide immutable runtime storage.

A runtime result must be stored in a variable:

```go
var currentTotal = calculateTotal()
```

Code can choose not to reassign `currentTotal`, but Go does not enforce that choice through `const`. The same limitation applies to arrays, slices, maps, structs, and individual struct fields.

Calls to `complex`, `real`, and `imag` produce constants only when their operands satisfy the constant rules. Similarly, `len` and `cap` are constant only for specific arguments, such as a constant string or qualifying array expression; they are not automatically constant for runtime slices or other runtime values.

## 5. Quick Cheat Sheet

- `const` → names a value allowed by Go's compile-time constant rules.
- Allowed kinds → boolean, integer, floating point, complex, `rune`, and `string`.
- `const maxCount int64 = 10` → typed constant.
- `const total = 20 * 10` → untyped constant expression with value `200`.
- `const (...)` → groups related constants.
- Constants may appear at package level or inside functions.
- A constant cannot be reassigned, incremented, or decremented.
- Variables and ordinary function calls cannot be constant operands.
- Some uses of `complex`, `real`, `imag`, `len`, and `cap` produce constants.
- Arrays, slices, maps, structs, and struct fields cannot be declared constant.
- `iota` is available in constant declarations and is studied separately.
- ⚠️ `const` is not a general immutable-variable or `readonly` feature.

### Test Yourself

1. Why is `const total = 20 * 10` valid while `const total = x + y` is invalid when `x` and `y` are variables initialized with `20` and `10`?
2. Why can neither a slice literal nor a predictable ordinary function result be made immutable simply by placing it in a `const` declaration?
3. Under what kind of conditions can calls to `len`, `cap`, `complex`, `real`, or `imag` participate in a constant expression?
