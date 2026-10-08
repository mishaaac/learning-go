# Exercise 6 — Candidate Filter

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Nested `for-range`, labeled `continue`, runes, control flow

## Goal

Filter complete text candidates by abandoning an outer iteration as soon as a forbidden rune is found.

## Scenario

A lightweight name check receives `go`, `api!`, `世界`, `bad#tag`, and `loop`. The runes `!` and `#` make an entire candidate unacceptable.

## What You Need to Do

- [x] Inspect the candidates in their given order and examine each candidate rune by rune.
- [x] Reject a candidate immediately when either forbidden rune appears, without inspecting or reporting the rest of that candidate.
- [x] Print only fully accepted candidates and finish with accepted and rejected totals.

## Rules and Constraints

- Use nested `for-range` loops.
- Use a label with `continue` to move directly from the inner scan to the next candidate.
- Print an accepted candidate only after its complete inner scan succeeds.

## Expected Result

Only candidates containing neither forbidden rune are accepted, Unicode text is examined as runes, and each input contributes to exactly one final total.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
