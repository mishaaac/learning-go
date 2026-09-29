# Untyped Literals in Go

The basic literals studied here denote constants that do not initially have a fixed Go type. This lets a literal adapt to a compatible context while still requiring its value to be representable by the destination type.

## 1. What I Need to Understand

- An integer, floating-point, `rune`, or `string` literal initially denotes an **untyped constant**.
- Context can give an untyped constant a concrete type without an explicit conversion.
- If no context supplies a type, Go uses the literal kind's default type, such as `int` for an integer literal or `float64` for a floating-point literal.
- A literal can adapt only when its kind is compatible with the required type and its value is representable by that type.
- A variable already has a concrete type, so two typed numeric variables of different types normally require an explicit conversion before they can be combined.
- A compatible literal can also initialize a defined type whose underlying type accepts that value.

> ⚠️ **Important correction:** The statement “literals are untyped” applies to the basic literal forms studied here. A composite literal such as `[]int{1, 2}` includes its type and is not an untyped constant.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Untyped constant | A constant whose concrete Go type has not yet been selected | It can adapt to a compatible type required by context |
| Contextual type | A type required by an assignment, declaration, or expression | It can determine the type used for an untyped constant |
| Default type | The type selected when no other context supplies one | An integer literal defaults to `int`; a floating-point literal defaults to `float64` |
| Representability | Whether a constant value can be stored as the required type | An out-of-range or incompatible value causes a compile-time error |
| Typed value | A value whose type is already fixed | It does not gain the same flexibility as an untyped constant |
| Defined type | A new type declared from an underlying type | A compatible, representable literal can initialize it directly |

The central comparison is:

| Situation | Result |
|---|---|
| Untyped constant + compatible typed value | The constant can take the typed value's type |
| Two typed values of different numeric types | An explicit conversion is normally required |
| Literal incompatible with the required type | Compile-time error |
| Compatible numeric literal outside the destination range | Compile-time error |

For integer destinations, the constant must be an exact integer within range. For floating-point destinations, the constant must not overflow; its exact value may be rounded to the destination type's precision.

## 3. How It Works in Go

An explicit destination type supplies the context for a literal:

```go
var distance float64 = 10
```

The literal `10` remains untyped until the declaration requires `float64`. Because its value is representable as `float64`, no `float64(10)` conversion is needed.

An untyped constant can also adapt to a typed operand in an expression:

```go
var unitPrice float64 = 200.3
var total = unitPrice * 5
```

`unitPrice` fixes the operation's type as `float64`, and the untyped constant `5` is representable by that type. `total` therefore has type `float64`.

If the whole expression is constant, it can remain untyped until the assignment supplies a type:

```go
var calculated float64 = 200.3 * 5
```

Without a destination type, Go selects default types:

```go
var count = 10       // int
var ratio = 200.3    // float64
```

Literals can initialize compatible defined types as well:

```go
type Score int16

var score Score = 100
```

The compiler applies this decision process:

```text
basic literal
     │
     └── untyped constant
             │
             ├── compatible context exists
             │       ├── value representable → use the contextual type
             │       └── value not representable → compile-time error
             │
             └── no contextual type → use the default type
```

Numeric constants have exact values while they remain constants. When a value becomes a concrete floating-point value, it is represented with that type's limited precision.

## 4. Examples, Differences, and Common Mistakes

| Declaration or expression | Valid? | Reason |
|---|---:|---|
| `var value float64 = 10` | ✅ | `10` is representable as `float64` |
| `var value = 10` | ✅ | With no other context, `10` gets the default type `int` |
| `var small byte = 255` | ✅ | `255` is within the range of `byte` |
| `var small byte = 1000` | ❌ | `1000` is outside the range of `byte` |
| `var whole int = 10.0` | ✅ | The exact constant value is the integer `10` |
| `var whole int = 10.5` | ❌ | `10.5` is not representable as an `int` |
| `var label string = 10` | ❌ | A numeric constant is not compatible with `string` |

> ⚠️ **Important correction:** A floating-point literal is not rejected by an integer destination merely because it contains a decimal point. `10.0` has an exact integral value and can initialize an `int`; `10.5` cannot.

> **Do not confuse:** an untyped constant is flexible at compile time; it is not a dynamically typed value that changes type at runtime.

A typed variable does not adapt in the same way as a literal:

```go
var count int = 10
var measurement float64 = 30.2

var valid = measurement + 10
// var invalid = measurement + count // mismatched types: float64 and int
```

The untyped `10` can be used as `float64`, but the typed variable `count` remains an `int`. Combining `measurement` and `count` requires choosing and writing an explicit conversion.

> **Do not confuse:** compatibility and range are separate checks. `1000` is numeric and therefore compatible in kind with `byte`, but its value is not representable by `byte`.

Assignments from untyped constants are checked at compile time. This prevents an incompatible or out-of-range constant from silently becoming a concrete value.

## 5. Quick Cheat Sheet

- Basic integer, floating-point, `rune`, and `string` literals denote untyped constants.
- Context can supply a compatible concrete type.
- No context → use the default type; integer → `int`, floating point → `float64`.
- A literal's value must be representable by the destination type.
- `var value float64 = 10` → valid without an explicit conversion.
- `measurement * 5` → `5` can adapt to the type of `measurement`.
- Typed values of different numeric types normally require an explicit conversion.
- Defined types can be initialized by compatible, representable literals.
- ⚠️ `var whole int = 10.0` is valid; `var whole int = 10.5` is not.
- ⚠️ `var small byte = 1000` fails because the value is out of range.
- Numeric and `string` constants are not interchangeable.
- `untyped` means compile-time flexibility, not absence of type-system rules.

### Test Yourself

1. Why does `measurement + 5` compile when `measurement` is a `float64`, while adding a variable of type `int` normally does not?
2. Why can `10.0` initialize an `int` but `10.5` cannot?
3. What happens when an untyped literal has no contextual type, and how does this differ from assigning it to a defined type such as `Score`?
