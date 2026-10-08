# Debugging 1 — Skipped Job Summary

> **Difficulty:** Intermediate
> **Topics:** `goto`, labels, blocks, declaration scope, `if`

## Scenario

A dispatcher must print one final summary for every job. A pending job should be reported as skipped, while any other job should be reported as processed, but the current program does not compile.

## What You Need to Do

- [x] Identify the problem.
- [x] Explain why it happens.
- [x] Fix the code.
- [x] Verify the corrected behavior.
- [x] Explain why the fix works.

## Rules and Constraints

Keep the two possible outcomes and a single final summary line. Do not hardcode the complete expected output, and ensure the corrected program also behaves correctly if `jobStatus` changes from `pending` to `active`.

## Expected Result

With the provided status, the program prints `pending skipped`. With the status changed to `active`, it prints `active processed`.

## Definition of Done

- [x] The corrected program behaves as required.
- [x] I completed `analysis.md`.
- [x] I verified my fix.
- [x] Ready for review.
