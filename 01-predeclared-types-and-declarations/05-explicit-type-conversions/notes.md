# Explicit Type Conversions in Go

Go normally requires conversions between typed values of different numeric types to be visible in the code. This makes the chosen result type and any possible change in value explicit, while boolean conditions must be expressed as actual comparisons rather than through truthiness.

## 1. What I Need to Understand

- A conversion uses the form `Type(value)`, such as `float64(x)` or `int(y)`.
- Typed variables with different numeric types are not automatically promoted to a common type.
- The chosen conversion determines the operation's type and can change the resulting value.
- Numeric conversions are not necessarily lossless: fractions, high integer bits, or floating-point precision may be lost.
- Go has no truthiness for numbers, strings, or other non-boolean values; a condition must produce a boolean value.
- Comparisons such as `x == 0` and `text == ""` produce `bool` values without converting their operands to `bool`.

> ⚠️ **Important correction:** The lack of automatic promotion applies to typed values of distinct types. A representable untyped constant can adapt to the required numeric type, and aliases such as `byte` and `uint8` are the same type.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Explicit conversion | `Type(value)` produces a value of another type | The intended type is visible in the code |
| Typed value | A value whose type is already fixed | Distinct typed numeric values usually require conversion before being combined |
| Untyped constant | A constant whose concrete type has not yet been selected | It can be used as a compatible type when its value is representable |
| Narrowing conversion | Conversion to a type with less range or precision | The converted value may change |
| Truthiness | Treating a non-boolean value as true or false | Go does not use this rule for numbers or strings |
| Comparison | An expression such as `x != 0` | It produces the boolean condition explicitly |

Common operations from the source are:

| Goal | Go expression | Effect |
|---|---|---|
| Convert `int` to `float64` | `float64(x)` | Produces a floating-point value |
| Convert `float64` to `int` | `int(y)` | Discards the fractional part toward zero |
| Convert `byte` to `int` | `int(data)` | Produces an `int` with the same value |
| Convert `int` to `byte` | `byte(x)` | Produces a `uint8`; high bits may be discarded |
| Test whether a number is zero | `x == 0` | Produces a `bool` |
| Test whether a string is empty | `text == ""` | Produces a `bool` |

## 3. How It Works in Go

Given an `int` and a `float64`, one operand must be converted before they can be added:

```go
var x int = 10
var y float64 = 30.2

var sumFloat float64 = float64(x) + y
var sumInt int = x + int(y)
```

`sumFloat` is `40.2`. `sumInt` is `40` because converting the non-constant `y` to `int` discards its fractional part before addition.

Float-to-integer conversion truncates toward zero rather than rounding:

```go
var positiveFloat float64 = 30.9
var negativeFloat float64 = -30.9

var positiveInt = int(positiveFloat)
var negativeInt = int(negativeFloat)
```

The results are `30` and `-30`.

Different integer types follow the same explicit rule:

```go
var count int = 10
var data byte = 100

var sumAsInt int = count + int(data)
var sumAsByte byte = byte(count) + data
```

The first addition is performed as `int`; the second is performed as `byte` (`uint8`). The target type is a design choice, not merely syntax.

Converting to a narrower integer type can change the value:

```go
var large int = 300
var reduced byte = byte(large)
```

`byte` has 8 bits, so this conversion keeps the low 8 bits and `reduced` becomes `44`. A conversion does not validate that the original value fits.

Representable untyped constants are an important exception to the usual conversion requirement:

```go
var measurement float64 = 10
var total = measurement + 2
```

The constants `10` and `2` are untyped and representable as `float64`, so no explicit conversion is needed. A typed `int` variable could not replace `2` in the addition without conversion.

Go obtains boolean values by stating conditions explicitly:

```go
var number = 10
var text = "Go"

var isZero = number == 0
var isNonZero = number != 0
var isEmpty = text == ""
```

`isZero`, `isNonZero`, and `isEmpty` have type `bool`. The numeric and string values were compared, not converted to `bool`.

## 4. Examples, Differences, and Common Mistakes

The conversion direction affects both type and value:

| Expression | Operation type | Result from the example |
|---|---|---:|
| `float64(x) + y` | `float64` | `40.2` |
| `x + int(y)` | `int` | `40` |
| `count + int(data)` | `int` | `110` |
| `byte(count) + data` | `byte` | `110` |

> **Do not confuse:** conversion changes or reinterprets a value according to conversion rules; comparison asks a question and returns `true` or `false`.

Go rejects direct arithmetic such as `x + y` when `x` is a typed `int` and `y` is a typed `float64`. Convert the operand that matches the type in which the calculation should occur.

Numeric conversion may lose information:

- Converting a float to an integer discards the fraction toward zero.
- Converting an integer to a smaller integer type discards high bits after extension to infinite precision.
- Converting an integer or float to a floating-point type rounds to the destination's precision.
- Converting a non-constant floating-point value outside the destination integer's range has an implementation-dependent result; validate the range first.

Constant conversions are checked at compile time. For example, `int(30.2)` is invalid because the constant `30.2` is not representable as an `int`, while `int(y)` is permitted when `y` is a `float64` variable and follows the runtime numeric conversion rules.

> **Do not confuse:** Go's lack of truthiness means `if number` and `if text` are invalid. Write the intended condition, such as `number != 0` or `text != ""`.

Numbers and strings cannot be converted directly to `bool`; expressions such as `bool(1)` and `bool("true")` are invalid.

> ⚠️ **Important correction:** A defined type whose underlying type is `bool` can be explicitly converted to `bool`. This is a conversion between boolean types, not truthiness.

```go
type Enabled bool

var custom Enabled = true
var standard bool = bool(custom)
```

The result is still based on an existing boolean value. Go never interprets a numeric or string value as true or false automatically.

## 5. Quick Cheat Sheet

- Conversion syntax → `Type(value)`.
- Typed numeric values of distinct types generally require explicit conversion.
- `float64(x)` → convert an integer value to `float64`.
- `int(y)` → convert a non-constant float to `int`, truncating toward zero.
- `int(data)` and `byte(x)` → convert between integer types.
- ⚠️ A conversion can lose range or precision; it is not validation.
- Representable untyped constants can adapt without an explicit conversion.
- `byte` and `uint8` are aliases, so no conversion is needed between them.
- Go has no numeric or string truthiness.
- `x != 0` and `text != ""` → explicit boolean conditions.
- Numeric/string → `bool` is invalid; compatible boolean types can convert to one another.

### Test Yourself

1. Why do `float64(x) + y` and `x + int(y)` produce different types and values when `x` is `10` and `y` is `30.2`?
2. Why can an untyped constant be combined with a `float64` variable without an explicit conversion while a typed `int` variable cannot?
3. Why are `bool(number)` and `if text` invalid, and how should those conditions be expressed in Go?
