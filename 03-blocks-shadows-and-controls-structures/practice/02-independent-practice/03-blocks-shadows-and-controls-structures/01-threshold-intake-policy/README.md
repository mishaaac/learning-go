# Problem 1 — Threshold Intake Policy

> **Difficulty:** Intermediate
> **Estimated time:** 30 minutes
> **Topics:** Choosing a `for` form, choosing `if` or `switch`, mutually exclusive branches, early exit

## Goal

Build a deterministic intake report that applies several policies in order and stops at a critical reading.

## Scenario

An intake stream contains the readings `-2`, `0`, `4`, `9`, `15`, and `6`. Negative readings are rejected, zero readings are ignored, values from `1` through `7` are standard, values from `8` through `14` are elevated, and values of `15` or more are critical.

## What You Need to Do

- [x] Process readings in their original order and report every rejection, accepted classification, and critical event that is reached.
- [x] Stop the intake immediately after reporting the first critical reading.
- [x] Print final totals for rejected, standard, elevated, and critical readings without including unreached input.
- [x] Be prepared to explain why your loop and branching choices fit this policy.

## Rules and Constraints

Each reached reading must produce at most one policy outcome. Ignored readings must not affect any total, and no reading after the critical event may affect output or state.

## Expected Result

The report follows input order, distinguishes all policy categories, terminates at the critical boundary, and ends with totals consistent with only the portion of the stream that was reached.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
