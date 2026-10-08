# Exercise 9 — Inventory Snapshot

> **Difficulty:** Applied
> **Estimated time:** 30 minutes
> **Topics:** Struct slices, `copy`, maps, comma-ok, independent storage, comparison

## Goal

Create an editable inventory working set while preserving an independent ordered snapshot for auditing.

## Scenario

The original catalog contains three products identified by SKU, including one with zero stock. Incoming adjustments may refer to either an existing SKU or a missing one.

## What You Need to Do

- [x] Model each product as a named struct and preserve the three source records in an independent snapshot before applying changes.
- [x] Build a SKU lookup map for the editable records, apply adjustments only when the SKU exists, and report whether each requested adjustment was accepted.
- [x] Print the source snapshot and final working set in stable order, then report whether their contents are still equal.

## Rules and Constraints

- Use a destination slice with sufficient length and `copy` to create the snapshot.
- Use comma-ok to distinguish a missing SKU from a valid product whose stock is zero.
- Do not derive report order from ranging over a map.
- Add a short comment explaining why copying the slice is sufficient for the field types in your product struct.

## Expected Result

Existing products receive their requested adjustments, a missing SKU leaves the inventory unchanged, and the audit snapshot retains its original values. The equality report reflects whether the two ordered slices still contain the same records.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
