# Strings, Runes, and Bytes in Go

A Go string is an immutable sequence of bytes, while a `rune` represents a Unicode code point. This distinction explains the behavior of indexing, `len`, slicing, and conversions for UTF-8 text.

## 1. What I Need to Understand

- A string stores bytes and is not required to contain valid UTF-8.
- Go source text is UTF-8, so ordinary string literals usually contain UTF-8-encoded text.
- `s[i]`, `s[a:b]`, and `len(s)` operate on bytes.
- A `rune` is an alias for `int32` and represents a Unicode code point.
- UTF-8 uses one to four bytes per code point, so byte boundaries are not always text boundaries.

## 2. Key Concepts

| Form | Represents |
|---|---|
| `byte` | Alias for `uint8`; one byte |
| `rune` | Alias for `int32`; one Unicode code point |
| `s[i]` | Byte at index `i` |
| `len(s)` | Number of bytes |
| `[]byte(s)` | Copy of the string's bytes |
| `[]rune(s)` | Decoded Unicode code points |

## 3. How It Works in Go

```go
text := "Hello 🌞"

fmt.Println(len(text))   // 10 bytes
fmt.Println(text[0])     // 72, the byte for 'H'
fmt.Println([]byte(text))
fmt.Println([]rune(text))
```

```text
"Hello " → 6 bytes
"🌞"      → 4 bytes
total     → 10 bytes, 7 code points
```

Use `range` to decode UTF-8 code points:

```go
for _, r := range text {
    fmt.Printf("%c ", r)
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a byte index is not necessarily a character or code-point index.

Slicing in the middle of a multibyte UTF-8 encoding can produce an invalid UTF-8 string. Strings remain immutable, but byte-based slicing still needs valid boundaries when the result is meant to be text.

> ⚠️ **Important correction:** `string(65)` converts the integer as Unicode code point U+0041 and yields `"A"`; it does not format the decimal digits as `"65"`. Numeric formatting requires an appropriate formatting or conversion function.

Converting `string` to `[]byte` preserves raw bytes; converting to `[]rune` decodes UTF-8 and substitutes `utf8.RuneError` for invalid encodings.

## 5. Quick Cheat Sheet

- String → immutable byte sequence
- A string may contain invalid UTF-8
- `byte` → alias for `uint8`
- `rune` → alias for `int32`, Unicode code point
- `s[i]` → one byte
- `s[a:b]` → byte range
- `len(s)` → byte count
- UTF-8 → one to four bytes per code point
- `[]byte(s)` → raw bytes
- `[]rune(s)` → decoded code points
- `string(65)` → `"A"`, not `"65"`

### Test Yourself

1. Why can slicing a valid UTF-8 string produce invalid UTF-8?
2. What different information do `[]byte(text)` and `[]rune(text)` expose?
3. Why does `string(65)` not produce the digits `"65"`?
