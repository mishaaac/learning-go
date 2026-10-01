# Debugging 1 — Incompatible Average

> **Difficulty:** Intermediate
> **Topics:** Numeric types, explicit conversions, floating-point arithmetic

## Scenario

A temperature summary should calculate an average from four samples totaling `97.0`. A recent type change caused the program to stop compiling.

## What You Need to Do

- [x] Identify the problem.
- [x] Explain why it happens.
- [x] Fix the code.
- [x] Verify the corrected behavior.
- [x] Explain why the fix works.

## Rules and Constraints

Keep the sample count as an `int` and the accumulated temperature as a `float64`. Preserve the source values, retain the fractional part of the average, and do not hard-code the expected result.

## Expected Result

The corrected program compiles and reports an average temperature of `24.25` derived from the supplied values.

## Definition of Done

- [x] The corrected program behaves as required.
- [x] I completed `analysis.md`.
- [x] I verified my fix.
- [x] Ready for review.
