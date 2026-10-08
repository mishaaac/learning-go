# Exercise 6 — Map Lifecycle

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Maps, zero values, comma-ok, updates, `delete`, `clear`, `len`

## Goal

Build an inventory report that handles missing keys, stored zero values, targeted deletion, and complete reuse of a map.

## Scenario

An inventory contains one item with positive stock and one valid item with zero stock. A discontinued item must be removed before the map is cleared for the next reporting cycle.

## What You Need to Do

- [x] Show the distinct lookup results for a stored zero value and an absent key, including presence information.
- [x] Insert or update stock counts, remove one existing entry, and confirm the resulting entry count without relying on map iteration order.
- [x] Clear every entry, prove that the initialized map is still writable, and add one entry for the next cycle.

## Rules and Constraints

- Use comma-ok wherever absence and a stored zero have different meanings.
- Use `delete` for one selected entry and `clear` for the full reset.
- Print explicitly selected keys so the report remains deterministic.

## Expected Result

The output distinguishes zero stock from a missing item, shows the correct lengths after deletion and clearing, and confirms that the cleared map accepts a new entry without reinitialization.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
