# Problem 3 — Route Control Audit

> **Difficulty:** Hard
> **Estimated time:** 45 minutes
> **Topics:** Nested `for-range`, string runes, labels, `break`, `continue`, `goto` design considerations

## Goal

Build a route audit whose inner scan can reject one route or terminate the entire audit without processing forbidden trailing data.

## Scenario

Audit the route labels `north`, `api!tail`, `世界`, `runXstop`, and `later` in order. The rune `!` rejects only its route, while `X` triggers an audit-wide shutdown.

## What You Need to Do

- [x] Inspect labels rune by rune and count ordinary runes that are actually reached.
- [x] Report completed routes, reject a route immediately at `!`, and stop the complete audit immediately at `X`.
- [x] Print final totals for completed routes, rejected routes, and reached ordinary runes.
- [x] Explain why your chosen control transfer is preferable here to the available alternatives, including `goto`.

## Rules and Constraints

Neither marker counts as an ordinary rune. Text after a marker in the same label must remain unprocessed, and every label after the shutdown marker must remain untouched. Each reached rune may be inspected only once.

## Expected Result

The event report preserves route order, keeps route rejection separate from global shutdown, handles the Unicode label by decoded rune, and produces totals based only on work reached before shutdown.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
