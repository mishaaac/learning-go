# Debugging 2 — Capacity Without Elements

> **Difficulty:** Hard
> **Topics:** Slice length, slice capacity, `make`, indexed assignment, arrays, structs

## Scenario

A sensor collector reserves enough storage for an incoming fixed batch. The program compiles, but it panics as soon as collection begins and never prints the completed batch.

## What You Need to Do

- [x] Identify the problem.
- [x] Explain why it happens.
- [x] Fix the code.
- [x] Verify the corrected behavior.
- [x] Explain why the fix works.

## Rules and Constraints

Keep the incoming data as a three-element array and preserve all sensor names, values, their order, and the loop. The corrected result must contain exactly three records without leading zero-valued placeholders. Do not hard-code the reported count or replace the collected output with the input array.

## Expected Result

The corrected program completes without a panic, reports an accepted count of `3`, and prints the `alpha`, `beta`, and `gamma` records in order with values `12`, `0`, and `27`.

## Definition of Done

- [x] The corrected program behaves as required.
- [x] I completed `analysis.md`.
- [x] I verified my fix.
- [x] Ready for review.
