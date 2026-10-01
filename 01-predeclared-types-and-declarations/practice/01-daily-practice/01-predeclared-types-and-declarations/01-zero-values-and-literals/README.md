# Exercise 1 — Zero Values and Literal Forms

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Predeclared types, zero values, integer literals, runes, strings

## Goal

Build a small program that makes Go's basic zero values and several literal forms visible.

## Scenario

You are preparing a quick reference program for a new Go developer. The program should show that declarations without initializers are predictable and that different integer notations can represent the same value.

## What You Need to Do

- [x] Declare an uninitialized `bool`, numeric value, and `string`, then display their zero values with clear labels.
- [x] Represent the value `42` with decimal, binary, octal, and hexadecimal integer literals, using an underscore where it improves readability.
- [x] Display one `rune` literal and one string literal so their values can be compared visually.

## Rules and Constraints

- Keep the zero-value declarations free of explicit initializers.
- Use only the predeclared basic types and literal forms covered in the study notes.
- Make the empty string visible in the output by placing it between delimiters or including a label.

## Expected Result

The output clearly shows `false`, numeric zero, and an empty string. All four integer notations produce the same numeric value, and the rune and string values are distinguishable in the report.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
