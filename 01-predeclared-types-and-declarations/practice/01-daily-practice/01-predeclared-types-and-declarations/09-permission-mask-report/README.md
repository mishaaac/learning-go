# Exercise 9 — Permission Mask Report

> **Difficulty:** Applied
> **Estimated time:** 25 minutes
> **Topics:** Integer literals, bitwise operators, aliases, booleans, naming

## Goal

Use binary masks and bitwise operations to produce a readable permission report.

## Scenario

A compact permission value uses three low-order bits for read, write, and execute access. The current value grants read and execute access but not write access.

## What You Need to Do

- [x] Define named binary constants for the read, write, and execute bits, then combine the required bits into one `byte` permission value.
- [x] Derive one boolean for each permission by testing the combined value against its mask.
- [x] Print the permission value in a labeled report together with the three boolean results.

## Rules and Constraints

- Use bitwise operators for both combining and testing permissions.
- Use `byte` when the value represents byte-sized data, and choose idiomatic mixed-capitalization names.
- Also represent the same combined numeric value once with an octal or hexadecimal literal and verify equality in the output.

## Expected Result

The report shows read and execute as enabled, write as disabled, and confirms that the alternate-base literal represents the same numeric value as the combined mask.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
