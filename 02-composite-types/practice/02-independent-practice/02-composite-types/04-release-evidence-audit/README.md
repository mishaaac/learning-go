# Problem 4 — Release Evidence Audit

> **Difficulty:** Integrative
> **Estimated time:** 60 minutes
> **Topics:** Arrays, slices, structs, maps as sets, comma-ok, bytes, runes, `copy`, `clear`

## Goal

Design a complete release-evidence audit that validates ordered records, measures Unicode text correctly, reports stage coverage, and preserves an independent archive.

## Scenario

The required stages are `plan`, `build`, `test`, and `deploy`. Process these entries in order: `K1 / plan / Draft`, `K2 / build / Café`, `K2 / test / duplicate`, `K3 / review / skip`, `K4 / test / 世界`, and `K5 / deploy / 🚀`.

## What You Need to Do

- [x] Accept an entry only when its stage is required and its identifier has not already been accepted; report duplicate-identifier and unknown-stage rejections separately.
- [x] Preserve accepted entries in input order, attach byte and rune counts to each one, and report deterministic counts for every required stage.
- [x] Report whether all required stages have accepted evidence and justify the composite types chosen for fixed stages, ordered records, uniqueness, and keyed counts.
- [x] Preserve an independent archive of the accepted records, clear the working records, and demonstrate the resulting working and archived states.

## Rules and Constraints

Represent the required stage collection with an array and each accepted entry with a named struct. Use at least one slice, one map-backed set, one keyed count map, comma-ok presence checks, `copy`, and `clear`. An identifier becomes used only when its entry is accepted. Do not rely on map iteration order or hard-code any derived count, coverage result, or text measurement.

## Expected Result

The audit reports four accepted entries, one duplicate-identifier rejection, and one unknown-stage rejection. Each required stage has one accepted entry, so coverage is complete. Accepted descriptions total 20 bytes and 12 runes. After the working records are cleared, their length remains four and their fields contain zero values, while the independent archive retains the four accepted records in order.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
