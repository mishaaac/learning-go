# Exercise 4 — Conversion Direction

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** Explicit conversions, integer division, floating-point arithmetic, narrowing conversions

## Goal

Show how the direction of an explicit numeric conversion affects both the type and value of a calculation.

## Scenario

A pricing check combines an item count stored as an integer with a decimal unit price. A separate legacy field stores only one byte and may not preserve a larger integer value.

## What You Need to Do

- [x] Calculate a total by converting the integer count so the multiplication keeps the unit price's fractional part.
- [x] Calculate a second total by converting the unit price to an integer before multiplication, then display both results.
- [x] Convert the integer value `300` to `byte`, display the result, and explain why the original value is not preserved.

## Rules and Constraints

- Start with a typed `int` count and a typed `float64` unit price.
- Make every conversion explicit and choose its direction intentionally.
- Add a brief comment to your completed code explaining the information lost by each narrowing conversion.

## Expected Result

The two totals differ in a way that exposes truncation before arithmetic. The byte conversion also produces a value different from `300`, and the explanation accounts for that change.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
