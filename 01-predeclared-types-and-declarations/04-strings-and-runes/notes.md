# Strings and Runes in Go

Go uses `string` for immutable text data and `rune` to express the intent of working with a Unicode code point. Their literals, zero values, aliases, and byte-level behavior are important when code contains more than simple ASCII text.

## 1. What I Need to Understand

- `string` is a predeclared type whose zero value is the empty string `""`.
- A string is an **immutable sequence of bytes**; a variable can receive a new string, but an existing string's contents cannot be changed.
- Strings support equality, ordering comparisons, and concatenation with `+`.
- `rune` is an alias for `int32` and is the conventional name for a Unicode code point.
- `byte` is an alias for `uint8`.
- A `rune` literal uses single quotes and defaults to `rune`; a `string` literal uses double quotes or backquotes and defaults to `string`.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| `string` | An immutable sequence of bytes | It stores text and other byte data but cannot be edited in place |
| Empty string | `""`, the zero value of `string` | An uninitialized `string` variable has a valid, defined value |
| `rune` | Alias for `int32`, conventionally used for a Unicode code point | The name communicates character-processing intent |
| `byte` | Alias for `uint8` | It communicates that a numeric value represents a byte |
| String literal | Text written with `"..."` or backquotes | Its default type is `string` |
| Rune literal | One code point or escape written with single quotes | Its default type is `rune` |

The main string operators are:

| Purpose | Operators | Behavior |
|---|---|---|
| Equality | `==`, `!=` | Compare complete string values |
| Ordering | `<`, `<=`, `>`, `>=` | Compare strings lexicographically, byte by byte |
| Concatenation | `+` | Produce a new string from two strings |

The ordering operators do not perform locale-aware or human-language sorting.

## 3. How It Works in Go

Variables declared without initializers receive their zero values:

```go
var text string
var symbol rune
```

`text` starts as `""`, while `symbol` starts as `0`, the numeric zero value of `int32` and therefore of its alias `rune`.

Go source can contain Unicode text and code points directly:

```go
var greeting = "Hello"
var world = "世界"
var initial rune = 'J'
```

When no other type is supplied, `"世界"` has the default type `string` and `'J'` has the default type `rune`.

Strings can be compared and concatenated:

```go
const fullMessage = "Hello, " + "Go"
const same = "Go" == "Go"
const comesFirst = "apple" < "banana"
```

`fullMessage` is a newly formed string. Concatenation does not modify either operand.

String variables can be reassigned:

```go
func replaceMessage() string {
    message := "first value"
    message = "another value"
    return message
}
```

The variable changes which string value it holds. The original string value itself is not modified.

Because `rune` and `byte` are aliases, values can be used with their underlying types without conversion:

```go
var initial rune = 'J'
var codePoint int32 = initial

var data byte = 65
var number uint8 = data
```

The alias names communicate intent: `rune` suggests a code point, while `byte` suggests byte data.

Unicode text in a source string literal is encoded as UTF-8. Indexing a string still reads one byte, not one `rune`:

```go
var word = "世界"
var firstByte = word[0]
```

`firstByte` has type `byte` (`uint8`) and is only the first byte of the UTF-8 encoding. Detailed conversion and iteration between strings, bytes, and runes belongs to the later text-processing topic.

## 4. Examples, Differences, and Common Mistakes

| Aspect | `string` | `rune` |
|---|---|---|
| Intended use | Text or byte data | One Unicode code point |
| Literal example | `"Go"` | `'G'` |
| Literal delimiters | Double quotes or backquotes | Single quotes |
| Default literal type | `string` | `rune` |
| Underlying relationship | Predeclared defined type | Alias for `int32` |
| Zero value | `""` | `0` |
| Mutability | Immutable | A numeric value, not a container |

> **Do not confuse:** reassigning a `string` variable changes the value stored in the variable; it does not mutate the previous string.

This is valid reassignment:

```go
func rename() string {
    name := "Gopher"
    name = "Go developer"
    return name
}
```

An assignment such as `name[0] = 'g'` is invalid because string elements are read-only. Creating a different string and assigning it to `name` is valid.

> **Do not confuse:** `"G"` is a `string`; `'G'` is a `rune`. Double and single quotes are not interchangeable.

Use the alias that communicates the value's purpose:

```go
var firstInitial rune = 'J'
var lastInitial rune = 'B'
```

Using `int32` would compile, but `rune` makes the code's intent clearer. Because `rune` is an alias rather than a validation type, an arbitrary `int32` value is not necessarily a valid Unicode code point.

> **Do not confuse:** indexing a string returns a byte. A Unicode code point may occupy more than one UTF-8 byte, and a visible character can sometimes contain more than one code point.

String ordering is useful for deterministic lexical comparison, but it should not be treated as language-aware alphabetical order.

## 5. Quick Cheat Sheet

- `string` → immutable sequence of bytes.
- Zero value of `string` → `""`.
- String comparisons → `==`, `!=`, `<`, `<=`, `>`, `>=`.
- String ordering → lexicographical byte comparison, not locale-aware sorting.
- String concatenation → `+`, producing a new string.
- Reassigning a variable is allowed; modifying a string element is not.
- `rune` = `int32`; use it to communicate Unicode code point intent.
- `byte` = `uint8`; use it to communicate byte-data intent.
- Rune literal → `'G'`; default type `rune`.
- String literal → `"Go"`; default type `string`.
- ⚠️ String indexing yields a byte, not necessarily a complete `rune` or visible character.

### Test Yourself

1. Why can a `string` variable be reassigned even though strings are immutable?
2. What do `rune` and `byte` alias, and what intent does each name communicate?
3. Why might `word[0]` fail to represent the first visible character of a Unicode string?
