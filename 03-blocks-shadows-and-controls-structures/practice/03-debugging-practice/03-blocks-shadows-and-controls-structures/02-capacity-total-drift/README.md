# Debugging 2 — Capacity Total Drift

> **Difficulty:** Intermediate
> **Topics:** Variable shadowing, short declarations, assignment, `for-range`, `if`, `continue`

## Scenario

A reservation service accepts requests only while their cumulative amount remains within a capacity of `10`. The program runs, but its decisions and final total do not reflect the accepted reservations.

## What You Need to Do

- [x] Identify the problem.
- [x] Explain why it happens.
- [x] Fix the code.
- [x] Verify the corrected behavior.
- [x] Explain why the fix works.

## Rules and Constraints

Keep the capacity at `10` and the request sequence `4`, `7`, `3`. Process every request in order, do not hardcode the report, and never allow the accepted total to exceed capacity.

## Expected Result

The request for `4` is accepted with a running total of `4`, the request for `7` is rejected, and the request for `3` is accepted with a running total of `7`. The final reported total is `7`.

## Definition of Done

- [x] The corrected program behaves as required.
- [x] I completed `analysis.md`.
- [x] I verified my fix.
- [x] Ready for review.
