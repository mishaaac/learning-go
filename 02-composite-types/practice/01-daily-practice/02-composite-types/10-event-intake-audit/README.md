# Exercise 10 — Event Intake Audit

> **Difficulty:** Challenge
> **Estimated time:** 40 minutes
> **Topics:** Arrays, slices, maps as sets, comma-ok, structs, Unicode, `copy`, `clear`

## Goal

Build a deterministic intake audit that validates structured events, rejects duplicates, measures Unicode text, and preserves a final snapshot.

## Scenario

The allowed categories are `build`, `test`, and `deploy`. Process these events in order: `A1/build/Compile`, `A2/test/世界`, `A1/deploy/duplicate`, `A3/unknown/skip`, and `A4/deploy/Ship 🌞`.

## What You Need to Do

- [x] Store the three allowed categories in a fixed-size array and use them to prepare membership checks.
- [x] Accept only events with an allowed category and an ID not previously accepted, preserving accepted events in input order and counting each rejection reason.
- [x] Store byte and rune counts with every accepted event, and maintain a per-category count without depending on map iteration order for output.
- [x] Create an independent slice snapshot of the accepted records, clear the original accepted slice, and demonstrate the different contents while both lengths remain observable.

## Rules and Constraints

- Model incoming items with an anonymous struct shape and accepted items with a named struct type.
- Use map-backed sets and comma-ok where membership must be distinguished from a zero value.
- An ID counts as seen only after its event is accepted.
- Preallocate the accepted slice without creating placeholder records, and use `copy` for the independent snapshot.

## Expected Result

The report shows three accepted events, one duplicate-ID rejection, and one category rejection. Category counts are one each for `build`, `test`, and `deploy`; accepted titles total 22 bytes and 15 runes. Clearing the working slice zeroes its three records without erasing the independent snapshot.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
