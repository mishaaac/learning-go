# Problem 4 — Dispatch Control System

> **Difficulty:** Integrative
> **Estimated time:** 60 minutes
> **Topics:** Block scope, nested loops, loop selection, labels, `break`, `continue`, expression and expressionless `switch`

## Goal

Design and build a multi-route dispatch audit with local route completion, global emergency shutdown, ordered classifications, and reliable final state.

## Scenario

Four named routes carry control codes in order: `North: [2, 0, 8, -1, 99]`, `East: [5, 11, 4]`, `South: [7, -9, 3]`, and `West: [1]`. Zero is ignored, `-1` closes the current route successfully, and `-9` shuts down the entire dispatch system.

Positive codes below `5` are light, codes from `5` through `9` are standard, and codes of `10` or more are heavy. A route that reaches the end of its codes also completes successfully.

## What You Need to Do

- [x] Process routes and their codes in the stated order, reporting every reached positive code with its route and classification.
- [x] Enforce route closure and system shutdown so trailing codes and routes cannot affect the report.
- [x] Report completed routes and global light, standard, and heavy totals after processing ends.
- [x] Explain the scope of route-local and audit-wide state and justify the control structures used for traversal, classification, closure, and shutdown.

## Rules and Constraints

Every reached code must have exactly one effect. Control codes must not enter weight totals, route completion must be counted once per completed route, and all audit-wide totals must remain available after processing stops.

## Expected Result

The report reflects two completed routes before shutdown, excludes `North`'s trailing code, stops during `South`, leaves `West` untouched, and preserves accurate classification totals for all positive codes reached before shutdown.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
