# Exercise 3 — Predict Text Representation

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Strings, runes, bytes, raw strings, interpreted strings

## Goal

Predict how Go represents Unicode text, rune values, and escape sequences.

## Scenario

A text diagnostic mixes Unicode, string indexing, and two forms of string literal. Analyze what each operation observes before you execute it.

```go
package main

import "fmt"

func main() {
    text := "世界"
    var symbol rune = '世'
    interpreted := "line1\nline2"
    raw := `line1\nline2`

    fmt.Printf("%T %T\n", text[0], symbol)
    fmt.Println(len(text) > 2)
    fmt.Println(interpreted == raw)
}
```

## What You Need to Do

- [x] Predict the exact three output lines before running the snippet.
- [x] Explain why indexing `text` and declaring `symbol` produce values with different types.
- [x] Explain why the interpreted and raw strings compare as equal or unequal.

## Rules and Constraints

- Record your prediction before copying the snippet into `main.go`.
- Treat a string as an immutable sequence of bytes, not as a container of runes.
- Do not rewrite either string literal until the original prediction has been checked.

## Expected Result

Your prediction correctly identifies the two printed types and both boolean results. Your explanation distinguishes bytes, runes, and the handling of escape sequences.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
