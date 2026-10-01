# Debugging 2 — Scoped Status Update

> **Difficulty:** Hard
> **Topics:** Short declarations, assignment, block scope, shadowing

## Scenario

A job tracker prints the correct values while an update is in progress, but its stored state appears unchanged immediately afterward. The program compiles without errors.

## What You Need to Do

- [x] Identify the problem.
- [x] Explain why it happens.
- [x] Fix the code.
- [x] Verify the corrected behavior.
- [x] Explain why the fix works.

## Rules and Constraints

Keep the nested block, the initial state, and the update values. The update inside the block must affect the variables reported after the block, and the final output must be derived from those variables rather than from replacement literals.

## Expected Result

Both report lines show the state `running` and the attempt count `1`. The values printed after the nested block come from the same bindings that were declared before it.

## Definition of Done

- [x] The corrected program behaves as required.
- [x] I completed `analysis.md`.
- [x] I verified my fix.
- [x] Ready for review.
