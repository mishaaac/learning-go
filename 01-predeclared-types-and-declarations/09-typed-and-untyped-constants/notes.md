# Typed and Untyped Constants in Go

Go constants can be typed or untyped. This distinction controls whether a constant can adapt to a compatible context or already has a fixed type that follows ordinary assignment rules.

## 1. What I Need to Understand

- `const count = 10` declares an **untyped constant**; `const count int = 10` declares a typed constant.
- An untyped constant has a kind, such as untyped integer or untyped string, and also has a default type used when no other type is supplied.
- Context may give an untyped constant a compatible type, but its value must be representable by that type.
- If no context chooses a type, Go uses the constant's default type.
- A typed constant already has a fixed type and does not implicitly adapt to a distinct numeric type.
- Variables are always typed: initializing a variable from an untyped constant eventually selects a concrete type.
- Untyped constants are usually more reusable; typed constants are useful when the type is part of the intended meaning.

> ⚠️ **Important correction:** Saying that a typed constant can be assigned “only to the corresponding type” is shorthand. It follows Go's normal assignability rules, but a typed `int` constant does not implicitly become `float64` or `byte`; an explicit conversion is required for those distinct numeric types.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Untyped constant | A constant without a fixed concrete type | It can adapt to a compatible contextual type |
| Constant kind | The untyped category of a constant, such as integer, floating point, or string | It limits which contexts are compatible even before a concrete type is selected |
| Default type | The concrete type selected when context provides no other type | It turns an untyped constant into a typed value when needed |
| Typed constant | A constant whose type is already fixed | It follows the assignment and operation rules of that type |
| Representability | Whether a constant value can be represented by a destination type | An out-of-range or incompatible assignment fails at compile time |
| Explicit conversion | A form such as `float64(value)` | It intentionally creates a constant of another compatible type |

The default types are:

| Untyped constant kind | Default type |
|---|---|
| Boolean | `bool` |
| `rune` | `rune` |
| Integer | `int` |
| Floating point | `float64` |
| Complex | `complex128` |
| String | `string` |

The central comparison is:

| Feature | Untyped constant | Typed constant |
|---|---|---|
| Fixed concrete type | No | Yes |
| Has a default type | Yes | Not needed; its type is already known |
| Can adapt to compatible contextual types | Yes, when representable | No implicit change to a distinct numeric type |
| Useful for | Reusable mathematical or general constant values | Communicating or enforcing a particular type |

## 3. How It Works in Go

An untyped constant can adapt to several compatible destinations:

```go
const count = 10

var integerCount int = count
var floatingCount float64 = count
var smallCount byte = count
```

The value `10` is representable by all three destination types, so the compiler uses the type required by each declaration.

When no destination type is written, the default type is selected:

```go
const count = 10
const ratio = 2.5
const letter = 'A'

var inferredCount = count // int
var inferredRatio = ratio // float64
var inferredLetter = letter // rune
```

A typed constant keeps its declared type:

```go
const typedCount int = 10

var exactCount int = typedCount
var convertedCount float64 = float64(typedCount)
```

`exactCount` receives the typed `int` constant directly. `convertedCount` requires the explicit conversion because `float64` is a distinct numeric type.

A conversion can also make the constant expression itself typed:

```go
const converted = float64(10)
```

`converted` is a typed `float64` constant even though its declaration does not contain a separate type before `=`.

When a typed and an untyped constant participate in a compatible expression, the untyped operand can adopt the typed operand's type:

```go
const base int = 10
const total = base + 2
```

`2` can be represented as `int`, so `total` is also a typed `int` constant.

```text
untyped constant
       │
       ├── compatible context exists
       │       ├── value representable → use contextual type
       │       └── value not representable → compile-time error
       │
       └── no contextual type → use default type
```

Typed constants are useful when a specific type carries meaning, including later patterns that define related values with `iota`. Detailed `iota` usage is outside this topic.

## 4. Examples, Differences, and Common Mistakes

| Declaration | Valid? | Result or reason |
|---|---:|---|
| `const value = 10; var x byte = value` | ✅ | The untyped value is representable by `byte` |
| `const value = 10; var x float64 = value` | ✅ | The untyped value adapts to `float64` |
| `const value = 10; var x = value` | ✅ | `x` receives the default type `int` |
| `const value int = 10; var x int = value` | ✅ | Source and destination use `int` |
| `const value int = 10; var x float64 = value` | ❌ | A typed `int` does not implicitly become `float64` |
| `const value int = 10; var x float64 = float64(value)` | ✅ | The explicit conversion produces a `float64` constant |
| `const value = 1000; var x byte = value` | ❌ | `1000` is outside the range of `byte` |
| `const value = 10.5; var x int = value` | ❌ | `10.5` is not representable as an `int` |

> **Do not confuse:** untyped does not mean universally compatible. The constant's kind and exact value must still be valid for the destination type.

The same untyped constant may work in one context and fail in another:

```go
const limit = 255

var byteLimit byte = limit
var integerLimit int = limit
```

Both declarations compile. If `limit` were `256`, the `byte` declaration would fail while the `int` declaration would remain valid.

> **Do not confuse:** a named untyped constant remains a constant; it does not become a variable merely because different uses select different types.

A variable inferred from an untyped constant is no longer untyped:

```go
const source = 10
var value = source
```

`value` has the concrete default type `int`. It cannot later adapt as though it were still an untyped constant.

Both typed and untyped constants are immutable constant values. Their difference concerns type selection and compatibility, not whether they can be reassigned.

## 5. Quick Cheat Sheet

- `const count = 10` → untyped integer constant.
- `const count int = 10` → typed `int` constant.
- Untyped constants have a kind and a default type.
- Defaults → `bool`, `rune`, `int`, `float64`, `complex128`, `string`.
- Context may select another compatible type for an untyped constant.
- The constant value must be representable by the selected type.
- `var count = untypedInteger` → default type `int`.
- Typed constants do not implicitly change to distinct numeric types.
- `float64(typedCount)` → explicit conversion to a typed `float64` constant.
- An untyped operand can adapt to a compatible typed operand in a constant expression.
- Variables always have concrete types and do not retain untyped flexibility.
- Prefer untyped constants for flexibility; use typed constants when the type communicates intent.

### Test Yourself

1. Why can the same untyped constant `10` initialize variables of type `int`, `float64`, and `byte`, while a typed `int` constant cannot initialize all three directly?
2. When does Go use an untyped constant's default type, and what type does each constant kind receive?
3. Why can an untyped constant with value `255` initialize a `byte`, while the same declaration fails after changing the value to `256`?
