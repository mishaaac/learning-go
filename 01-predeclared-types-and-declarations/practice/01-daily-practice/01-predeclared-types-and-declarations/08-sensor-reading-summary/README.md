# Exercise 8 — Sensor Reading Summary

> **Difficulty:** Applied
> **Estimated time:** 25 minutes
> **Topics:** Numeric types, explicit conversions, booleans, constants, strings, runes

## Goal

Build a compact sensor report that combines declarations, arithmetic, conversion, comparison, and text values.

## Scenario

A temperature sensor has collected three readings whose total is `73.5`. The report needs the average, a threshold check, a sensor label, and a unit symbol.

## What You Need to Do

- [x] Store the sample count as an integer and the accumulated temperature as `float64`, then calculate a floating-point average.
- [x] Compare the average with an untyped constant threshold of `24` and store the result as a boolean warning value.
- [x] Print a labeled summary containing the sensor name, sample count, average, warning value, and the rune `℃`.

## Rules and Constraints

- Convert the sample count explicitly before division so the fractional result is preserved.
- Use a named constant for the threshold and an idiomatic multiword name for the sensor label.
- Include a brief comment explaining why the conversion belongs on the sample count rather than on the computed average.

## Expected Result

The report shows an average of `24.5`, a true warning value, and all requested metadata. The calculation must use floating-point division rather than a corrected hard-coded result.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
