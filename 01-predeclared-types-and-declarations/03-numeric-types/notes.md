# Numeric Types in Go

Go's predeclared numeric types cover integers, floating-point values, and complex numbers. Choosing among them requires understanding size, precision, operators, conversions, and the difference between exact and approximate representations.

## 1. What I Need to Understand

- Go has three numeric families: **integers**, **floating-point numbers**, and **complex numbers**.
- Every numeric type has a zero value of `0`; for complex types this means `0 + 0i`.
- Use `int` for general integer work unless an external format requires an exact size or sign.
- Fixed-width integer types remain distinct from `int`, even when they happen to have the same size on a platform.
- Prefer `float64` for general floating-point work, but remember that floating-point values are approximations.
- Use `complex64` or `complex128` when a value needs real and imaginary components.
- Numeric literals begin as untyped constants and receive a type from context or a default type such as `int`, `float64`, or `complex128`.

## 2. Key Concepts

| Family | Predeclared types | Representation |
|---|---|---|
| Signed integers | `int8`, `int16`, `int32`, `int64`, `int` | Whole numbers, including negatives |
| Unsigned integers | `uint8`, `uint16`, `uint32`, `uint64`, `uint`, `uintptr` | Zero and positive whole numbers |
| Floating point | `float32`, `float64` | Approximate real values using IEEE 754 |
| Complex | `complex64`, `complex128` | A real part plus an imaginary part |

The fixed-width integer ranges are:

| Type | Range |
|---|---|
| `int8` | −128 to 127 |
| `int16` | −32,768 to 32,767 |
| `int32` | −2,147,483,648 to 2,147,483,647 |
| `int64` | −9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint8` | 0 to 255 |
| `uint16` | 0 to 65,535 |
| `uint32` | 0 to 4,294,967,295 |
| `uint64` | 0 to 18,446,744,073,709,551,615 |

Special integer names have specific roles:

| Type | Meaning |
|---|---|
| `byte` | Alias for `uint8` |
| `rune` | Alias for `int32`; detailed use is covered later |
| `int` | Signed integer whose size is either 32 or 64 bits |
| `uint` | Unsigned integer with the same size as `int` |
| `uintptr` | Unsigned integer large enough to hold the uninterpreted bits of a pointer value; detailed use is covered later |

`int` and `uint` are commonly 32 bits on 32-bit systems and 64 bits on most 64-bit systems, but the language only guarantees that each is either 32 or 64 bits.

Floating-point and complex types relate by component size:

| Type | Size or components | Practical note |
|---|---|---|
| `float32` | 32 bits; about 6–7 decimal digits of precision | Use when an existing format requires it or measurement shows a real memory need |
| `float64` | 64 bits; default type of floating-point literals | Preferred for general floating-point work |
| `complex64` | Two `float32` components | Lower-precision complex values |
| `complex128` | Two `float64` components; default complex type | Preferred for general complex work |

Useful floating-point range landmarks are:

| Type | Largest finite absolute value, approximately | Smallest positive nonzero value, approximately |
|---|---:|---:|
| `float32` | `3.4e38` | `1.4e-45` |
| `float64` | `1.8e308` | `4.9e-324` |

## 3. How It Works in Go

Integer types support arithmetic, comparison, bitwise operations, and shifts:

| Category | Operators |
|---|---|
| Arithmetic | `+`, `-`, `*`, `/`, `%` |
| Comparison | `==`, `!=`, `<`, `<=`, `>`, `>=` |
| Bitwise | `&`, \|, `^`, `&^` |
| Shifts | `<<`, `>>` |
| Arithmetic assignment | `+=`, `-=`, `*=`, `/=`, `%=` |
| Bitwise assignment | `&=`, \|=, `^=`, `&^=`, `<<=`, `>>=` |

Integer division produces an integer and truncates toward zero:

```go
const positiveQuotient = 5 / 3
const negativeQuotient = -5 / 3
```

The results are `1` and `-1`. Convert the operands before division when a floating-point result is required:

```go
var numerator int = 5
var denominator int = 3
var ratio = float64(numerator) / float64(denominator)
```

`ratio` is approximately `1.6666666666666667`.

Named integer types are not mixed implicitly:

```go
var count int = 10
var count32 int32 = int32(count)
```

The explicit conversion is required because `int` and `int32` are distinct types. `byte` and `uint8` are an exception to this distinction because `byte` is an alias for `uint8`.

Bitwise operators expose integer bit patterns directly:

```go
const mask = 0b1111
const selected = mask & 0b0101
const shifted = selected << 1
```

Floating-point types support `+`, `-`, `*`, `/`, and comparisons, but not `%`. Their large range comes with limited precision, so many decimal fractions are stored as nearby binary approximations.

A simple absolute-tolerance comparison can be written as:

```go
func almostEqual(left, right, tolerance float64) bool {
    difference := left - right
    if difference < 0 {
        difference = -difference
    }
    return difference <= tolerance
}
```

The appropriate tolerance depends on the magnitude of the values and the accuracy required by the problem.

The built-in `complex` function combines real and imaginary components:

```go
var value = complex(20.3, 10.2)

var real32 float32 = 2.5
var imaginary32 float32 = 3.1
var value32 = complex(real32, imaginary32)

var realPart = real(value)
var imaginaryPart = imag(value)
```

`value` receives the default type `complex128`; `value32` has type `complex64`. `real` and `imag` extract the two components: they return `float32` for `complex64` and `float64` for `complex128`.

Complex values support `+`, `-`, `*`, `/`, `==`, and `!=`, but not ordering comparisons. Imaginary literals use the suffix `i`:

```go
var imaginaryValue = 2.5i
var combined = 3 + 2.5i
```

The standard package `math/cmplx` provides additional operations for `complex128`, such as magnitude calculations.

## 4. Examples, Differences, and Common Mistakes

Use the integer type that matches the constraint:

| Situation | Usual choice | Reason |
|---|---|---|
| General integer work | `int` | It is Go's conventional general-purpose integer type |
| Binary format or protocol | Exact fixed-width signed or unsigned type | The external representation determines size and sign |
| Algorithm for multiple integer types | Generics with an appropriate integer constraint | One implementation can accept the intended integer types |

Older code may contain separate implementations for signed and unsigned values. The source notes `strconv.FormatInt` and `strconv.FormatUint` as familiar examples of separate APIs for `int64` and `uint64`.

> **Do not confuse:** equal storage size does not make `int`, `int32`, and `int64` interchangeable. Go still treats them as distinct named types and usually requires an explicit conversion.

`byte` is a true alias, so it can be assigned to, compared with, and used in arithmetic with `uint8` without conversion. The name `byte` is useful when the value represents byte data.

Floating-point selection follows a similar rule: use `float64` normally and `float32` when interoperability, storage, or measured memory requirements justify it.

> ⚠️ **Important correction:** Division by a constant zero is a compile-time error for both integer and floating-point constant expressions. At run time, an integer zero divisor causes a panic; a floating-point zero divisor produces an infinity for a nonzero numerator and `NaN` for `0 / 0`.

```go
func divideInteger(numerator, divisor int) int {
    return numerator / divisor
}

func divideFloat(numerator, divisor float64) float64 {
    return numerator / divisor
}
```

Calling `divideInteger` with a zero divisor panics. Calling `divideFloat` with a runtime zero divisor follows IEEE 754 behavior.

Floating-point values should not represent money or another quantity that requires exact decimal representation. Exact `==` and `!=` comparisons are valid Go, but they do not test approximate equality; use a domain-appropriate tolerance when approximation is intended.

The result of `complex` depends on its arguments:

| Arguments | Result |
|---|---|
| Two untyped numeric constants | Untyped complex constant; default `complex128` when a concrete default type is needed |
| Two `float32` values | `complex64` |
| One `float32` and one representable untyped constant | `complex64` |
| Two `float64` values | `complex128` |
| One `float64` and one representable untyped constant | `complex128` |

> ⚠️ **Important correction:** Two typed arguments of different sizes, such as `float32` and `float64`, do not fall back to `complex128`; they must be made the same type explicitly before calling `complex`.

Complex equality has the same exactness concern as its floating-point components. For an approximate comparison, a tolerance can be applied to the magnitude of the difference, for example with `cmplx.Abs(left - right)`.

Go provides complex types directly, which can be useful for work such as Mandelbrot sets or quadratic equations. Broader numerical facilities such as matrices and linear algebra are outside the language's built-in numeric types; the source mentions Gonum as a third-party option for numerical computing.

## 5. Quick Cheat Sheet

- Numeric families → integers, floating point, and complex numbers.
- Every numeric zero value → `0`; complex zero → `0 + 0i`.
- General integer → `int`; fixed external format → exact-width integer type.
- `byte` = `uint8`; `rune` = `int32`.
- `int` and `uint` are 32 or 64 bits; fixed-width types remain distinct.
- Integer division truncates toward zero; `%` is integer-only.
- Bitwise integer operators → `&`, `|`, `^`, `&^`, `<<`, `>>`.
- General floating point → `float64`; `float32` has about 6–7 decimal digits of precision.
- ⚠️ Floats are approximate: avoid them for exact decimal quantities and use a suitable tolerance for approximate equality.
- Constant division by zero is invalid; runtime integer division by zero panics; runtime float division follows IEEE 754.
- `complex64` uses `float32` parts; `complex128` uses `float64` parts.
- `complex(realPart, imaginaryPart)`, `real(value)`, `imag(value)`; imaginary literals end in `i`.

### Test Yourself

1. When should you choose `int` instead of a fixed-width integer, and why can an `int` not be assigned directly to an `int32`?
2. How do integer and floating-point division differ in result, precision, and behavior when a runtime divisor is zero?
3. What type does `complex` produce from two `float32` values, from two untyped constants, and from typed `float32` and `float64` values?
