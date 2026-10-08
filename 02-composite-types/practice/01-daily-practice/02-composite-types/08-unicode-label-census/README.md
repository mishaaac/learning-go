# Exercise 8 — Unicode Label Census

> **Difficulty:** Applied
> **Estimated time:** 30 minutes
> **Topics:** Slices, strings, bytes, runes, maps as sets, structs, preallocation

## Goal

Build a Unicode-aware label census that combines sequence storage, text representations, structured records, and uniqueness tracking.

## Scenario

A service receives the labels `Go`, `Gopher`, `世界`, `Go`, and `🌞`. Operators need per-label measurements and a deterministic summary without losing duplicate input records.

## What You Need to Do

- [x] Produce one structured record per input label containing the original text, its byte count, and its rune count.
- [x] Preserve all five records in input order while tracking unique label values with a map-backed set.
- [x] Print every record followed by totals for records, unique labels, bytes, and runes.

## Rules and Constraints

- Start the result slice at length zero and reserve capacity for the known input count.
- Use `map[string]struct{}` for uniqueness and presence checks appropriate to that representation.
- Derive text measurements from Go's byte and rune representations; do not hard-code the totals.

## Expected Result

The report contains five ordered records and four unique labels. It visibly shows that the Unicode labels can have different byte and rune counts, and all totals are derived from the input.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
