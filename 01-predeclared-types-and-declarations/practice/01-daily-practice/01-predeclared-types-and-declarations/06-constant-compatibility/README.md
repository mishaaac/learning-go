# Exercise 6 — Constant Compatibility

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** Constants, typed constants, untyped constants, representability, constant expressions

## Goal

Build a program that demonstrates the flexibility of an untyped constant and the fixed type of a typed constant.

## Scenario

A protocol limit must be usable in byte-oriented code and in general calculations. A retry count, however, is intentionally tied to the program's ordinary integer type.

## What You Need to Do

- [x] Declare an untyped constant with value `255` and use it to initialize `byte`, `int`, and `float64` variables.
- [x] Declare a typed `int` constant and use it in a `float64` calculation through an explicit conversion.
- [x] Add a compile-time constant expression and display all resulting values with clear labels.

## Rules and Constraints

- The final program must contain both a typed and an untyped named constant.
- Do not replace the named constants with mutable variables.
- Briefly explain why changing the untyped value from `255` to `256` would invalidate only the byte destination.

## Expected Result

The same untyped value initializes three distinct concrete types, the typed integer participates in floating-point work only after conversion, and the constant expression is evaluated successfully.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
