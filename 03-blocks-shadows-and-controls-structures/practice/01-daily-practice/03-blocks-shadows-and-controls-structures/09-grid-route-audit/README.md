# Exercise 9 — Grid Route Audit

> **Difficulty:** Applied
> **Estimated time:** 35 minutes
> **Topics:** Nested `for-range`, string byte offsets, runes, `switch`, labeled `break`, `continue`

## Goal

Scan a text grid in reading order, ignore blocked cells, and stop both loops at the first target.

## Scenario

The grid rows are `S.界`, `.#.`, `..X`, and `X..`. A report uses zero-based row numbers and the target's UTF-8 byte offset within its row.

## What You Need to Do

- [x] Traverse rows from top to bottom and each row from left to right as decoded runes.
- [x] Ignore `#`, count every other rune reached before the first `X`, and do not count the target itself.
- [x] At the first `X`, print its row and byte offset, terminate the complete grid scan, and print the final count.

## Rules and Constraints

- Use nested `for-range` loops and treat the inner index as a byte offset.
- Use `continue` for blocked cells and a labeled `break` for the found target.
- Do not inspect or count cells after the first target, including cells in later rows.

## Expected Result

The first target in reading order is reported once, its coordinate uses the requested units, the later target is never processed, and the pre-target count excludes blocked and target cells.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
