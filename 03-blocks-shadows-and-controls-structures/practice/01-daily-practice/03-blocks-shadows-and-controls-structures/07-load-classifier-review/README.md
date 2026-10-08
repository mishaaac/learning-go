# Exercise 7 — Load Classifier Review

> **Difficulty:** Intermediate
> **Estimated time:** 30 minutes
> **Topics:** Choosing `if` or `switch`, expressionless `switch`, `for-range`, ordered conditions

## Goal

Implement one clear classification decision per load reading and defend the control structures you select.

## Scenario

A service reports load readings of `0`, `3`, `7`, `10`, and `14`. Each reading must be classified as `idle`, `normal`, `high`, or `critical` using ordered thresholds.

## What You Need to Do

- [x] Classify `0` as `idle`, values from `1` through `6` as `normal`, values from `7` through `9` as `high`, and values of `10` or more as `critical`.
- [x] Print exactly one classification for every reading, preserving input order.
- [x] Explain during review why your chosen loop form and your chosen branching form communicate this decision clearly.

## Rules and Constraints

- Express classification as one mutually exclusive decision, not as independent checks that can print multiple categories.
- Arrange overlapping conditions so a broader rule cannot hide a more specific result.
- Use a traversal form suited to processing every element.

## Expected Result

The report contains five ordered lines with one correct category per reading, and you can compare your design with the reasonable alternative of an `if` chain or expressionless `switch`.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
