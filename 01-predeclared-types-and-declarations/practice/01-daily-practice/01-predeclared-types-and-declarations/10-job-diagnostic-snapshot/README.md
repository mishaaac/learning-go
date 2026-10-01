# Exercise 10 — Job Diagnostic Snapshot

> **Difficulty:** Challenge
> **Estimated time:** 35 minutes
> **Topics:** Declarations, zero values, constants, conversions, numeric operations, strings, runes, booleans

## Goal

Create a complete diagnostic snapshot that integrates the declaration and predeclared-type concepts from this section.

## Scenario

A download job named `Atlas` has completed 7 of 10 units and has used 1 of its 3 allowed attempts. Operators need a compact, trustworthy report without introducing data structures beyond the studied scope.

## What You Need to Do

- [x] Declare fixed job settings as a related constant group, including the job name and maximum attempts.
- [x] Store completed units as `int32`, total units as `int`, and compute a floating-point completion percentage through explicit conversion.
- [x] Report remaining attempts, whether completion has reached an untyped `75` percent threshold, a rune status marker, and at least one deliberately visible zero value.
- [x] Build a readable heading through string concatenation and print every field with a clear label.

## Rules and Constraints

- Include at least one typed constant and one untyped constant.
- Use `var` for a declaration where an explicit type or zero value communicates intent, and use `:=` for at least one straightforward local inference.
- Use explicit conversions only where typed numeric values would otherwise be incompatible or integer division would lose the fraction.
- Add two short comments explaining one type choice and one choice between a constant and a variable.

## Expected Result

The report shows `70` percent completion, `2` remaining attempts, and a false threshold result. It includes the heading, marker, zero-value field, and all labels, and every calculation is derived from the declared source values.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
