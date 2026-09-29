# Predeclared Types and Literals in Go

Go provides basic types with defined initial values and several literal forms for writing constants directly in source code. Understanding their syntax and default behavior prevents type, quoting, and numeric-representation mistakes.

## 1. What I Need to Understand

- **Predeclared types** are available without importing a package. They include `bool`, `string`, integer types, floating-point types, and aliases such as `byte` and `rune`.
- A variable declared without an explicit initial value receives the **zero value** of its type: `false` for `bool`, `0` for numeric types, and `""` for `string`.
- The literal forms studied here write integer, floating-point, `rune`, and `string` constants directly in source code.
- Integer literals can use decimal, binary, octal, or hexadecimal notation. `_` can group digits without changing the value.
- A `rune` literal represents one Unicode code point and uses single quotes. A `string` literal uses either double quotes or backquotes.
- The literal constants covered here begin as **untyped constants**. Context can give them a type; otherwise Go uses their default type.

> ⚠️ **Important correction:** Not every construct called a literal in Go is untyped. The integer, floating-point, `rune`, and `string` literals covered here denote untyped constants, while a composite literal such as `[]int{1, 2}` has a type.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Predeclared type | A type already declared by the language | It can be used without an import or a user declaration |
| Zero value | The value assigned when a variable has no explicit initializer | Every declared variable starts with a defined value |
| Literal | Source-code notation for a constant value | The delimiter or prefix determines how Go reads it |
| Untyped constant | A constant without a fixed type yet | It can adapt to a compatible type required by context |
| Default type | The type chosen when context supplies no other type | Integer → `int`, floating-point → `float64`, `rune` → `rune`, and string → `string` |

Integer bases use these forms:

| Base | Prefix | Example | Typical use |
|---|---|---|---|
| Decimal | None | `1234` | Normal numeric values |
| Binary | `0b` or `0B` | `0b1010` | Visible bit patterns |
| Octal | `0o` or `0O` | `0o777` | Values such as POSIX permissions |
| Hexadecimal | `0x` or `0X` | `0xFF` | Compact bit and byte patterns |

The older leading-zero octal form, such as `0777`, is valid, but `0o777` communicates the base more clearly.

## 3. How It Works in Go

Declaring basic variables without initializers applies their zero values automatically:

```go
var enabled bool
var count int
var price float64
var name string
```

Their values are `false`, `0`, `0`, and `""`, respectively. This avoids uninitialized variables with indeterminate contents.

Underscores may improve the readability of numeric literals:

```go
const population = 1_234_567
const mask = 0b1111_0000
const color = 0xFF_A0_00
```

An underscore may separate digits and may appear immediately after a base prefix, as in `0x_FF`. It cannot finish a literal, appear twice in a row, or separate a digit from a decimal point or exponent marker.

Floating-point literals can use decimal notation and a base-10 exponent with `e` or `E`:

```go
const avogadro = 6.03e23
```

A hexadecimal floating-point literal starts with `0x` or `0X` and requires a `p` or `P` exponent, which represents a power of two:

```go
const hexadecimalFloat = 0x12.34p5
```

Here, `0x12.34p5` equals `582.5` in decimal.

A `rune` literal contains one Unicode code point or one valid escape:

```go
const letter = 'a'
const newLine = '\n'
const tab = '\t'
const quote = '\''
const backslash = '\\'
```

The same `rune` value can be written with numeric escapes:

```go
const letterDirect = 'a'
const letterOctal = '\141'
const letterHex = '\x61'
const letterUnicode16 = '\u0061'
const letterUnicode32 = '\U00000061'
```

Double-quoted strings interpret escapes. Backquoted strings allow direct newlines and give backslashes no special meaning:

```go
const interpreted = "Greetings and\n\"Salutations\""

const raw = `Greetings and
"Salutations"`
```

A raw string cannot contain a backquote. An interpreted string cannot contain a literal newline or an unescaped double quote.

Untyped literal constants receive a default type only when a type is needed and the context does not select another one:

```go
var count = 42
var ratio = 3.5
var letter = 'a'
var message = "hello"

var small int8 = 42
```

The first four variables have the default types `int`, `float64`, `rune` (an alias for `int32`), and `string`. The last literal instead takes the contextual type `int8` because `42` is representable by it.

```text
literal in source
      │
      └── untyped constant
              ├── compatible type supplied by context
              └── default type when context supplies none
```

## 4. Examples, Differences, and Common Mistakes

| Feature | Interpreted string | Raw string |
|---|---|---|
| Delimiter | `"..."` | Backquotes: `` `...` `` |
| Escape sequences | Interpreted | Not interpreted |
| Direct newline | Not allowed | Allowed |
| Double quote inside | Must be escaped | Allowed directly |
| Backslash inside | Starts an escape | Kept as a backslash |
| Backquote inside | Has no delimiter role | Not allowed |

These two literals produce the same text:

```go
const interpretedText = "Greetings and\n\"Salutations\""
```

```go
const rawText = `Greetings and
"Salutations"`
```

> **Do not confuse:** `'a'` is a `rune` literal for one code point; `"a"` is a `string` literal.

In a double-quoted string, a double quote needs `\"`, while a single quote can appear directly. In a `rune` literal, a single quote needs `\'`.

```go
const apostropheText = "It's valid"
const singleQuote = '\''
```

Common numeric mistakes include invalid underscore placement and unclear octal notation:

- `123_` is invalid because the underscore finishes the literal.
- `1__234` is invalid because underscores cannot be consecutive.
- `1_.5` is invalid because an underscore cannot touch the decimal point.
- `0777` is valid octal, but `0o777` is clearer.

Use decimal for ordinary integer and floating-point values. Binary and hexadecimal are useful when bit groupings matter, while octal is useful for values conventionally written in base 8, such as POSIX permissions.

> **Do not confuse:** `e` in a decimal floating-point literal scales by a power of 10; `p` in a hexadecimal floating-point literal scales by a power of 2.

Numeric `rune` escapes are valid, but the visible character is usually clearer unless the escape communicates the intent better.

## 5. Quick Cheat Sheet

- Predeclared basic types can be used without imports.
- Zero values: `bool` → `false`, numeric types → `0`, `string` → `""`.
- Integer prefixes: `0b` binary, `0o` octal, `0x` hexadecimal; ordinary unprefixed digits are decimal, apart from the legacy leading-zero octal form.
- `_` improves numeric readability without changing the value; placement is restricted.
- Decimal floats use `e`; hexadecimal floats use `0x` with `p`.
- A `rune` literal uses single quotes: `'a'`.
- An interpreted `string` uses double quotes and processes escapes.
- A raw `string` uses backquotes, permits direct newlines, and does not process escapes.
- ⚠️ Prefer `0o777` over the less clear leading-zero form `0777`.
- The numeric, `rune`, and `string` literals studied here are untyped constants and use context or a default type.

### Test Yourself

1. What values do uninitialized `bool`, numeric, and `string` variables receive, and why is this behavior useful?
2. When would you choose a `rune`, an interpreted string, or a raw string for the same visible character or text?
3. Why can the literal `42` become `int` in `var count = 42` but `int8` in `var small int8 = 42`?
