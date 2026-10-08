# Exercise 5 — Copy or Share

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** `copy`, array-to-slice views, slice-to-array conversion, array pointers, shared storage

## Goal

Build a storage experiment that clearly distinguishes an independent copy from a shared view.

## Scenario

A calibration program receives four measurements. Some consumers must observe live edits, while an audit snapshot must preserve the original values.

## What You Need to Do

- [x] Start with four measurements in a slice and create an independent slice snapshot with enough destination length for `copy`.
- [x] Create both an independent array value and a storage-sharing array pointer from the measurement slice.
- [x] Mutate the source and the derived values, then print a labeled report that makes every sharing relationship observable.

## Rules and Constraints

- Use all three relevant mechanisms: `copy`, slice-to-array value conversion, and slice-to-array pointer conversion.
- The target array length must not exceed the source slice length.
- Add a short comment explaining why destination length, rather than spare capacity, controls how many elements `copy` can write.

## Expected Result

The report demonstrates that the copied slice and array value can change independently, while edits through the array pointer and source slice remain visible through their shared storage.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
