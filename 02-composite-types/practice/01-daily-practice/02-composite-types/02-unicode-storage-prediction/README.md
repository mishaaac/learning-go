# Exercise 2 — Unicode Storage Prediction

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Strings, bytes, runes, UTF-8, `len`, string slicing

## Goal

Predict the difference between the byte representation and the rune representation of Unicode text.

## Scenario

A text probe inspects the same string through byte-based and code-point-based operations. Analyze every result before using the compiler to confirm it.

```go
package main

import "fmt"

func main() {
	text := "Go, 世界"
	bytes := []byte(text)
	runes := []rune(text)
	window := text[4:7]

	fmt.Println(len(text), len(bytes), len(runes))
	fmt.Printf("%q %d %c\n", window, text[4], runes[4])
}
```

## What You Need to Do

- [x] Predict the exact two output lines before running the snippet.
- [x] Account for the UTF-8 width of each non-ASCII code point in the length values.
- [x] Explain why `text[4]` and `runes[4]` observe different units.

## Rules and Constraints

- Record the complete prediction before copying or running the code.
- Treat string indexing and slicing as byte-based operations.
- Do not replace the text or slice bounds until the original result has been checked.

## Expected Result

Your prediction distinguishes byte count from rune count, identifies a valid byte-aligned substring, and explains the numeric byte and displayed rune without treating them as interchangeable.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
