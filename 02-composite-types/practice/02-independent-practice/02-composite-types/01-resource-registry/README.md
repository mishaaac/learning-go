# Problem 1 — Resource Registry

> **Difficulty:** Intermediate
> **Estimated time:** 30 minutes
> **Topics:** Structs, slices, maps, maps as sets, comma-ok, Unicode strings

## Goal

Design and build an ordered resource registry that rejects duplicate identifiers and distinguishes missing records from valid zero-valued data.

## Scenario

Process these resources in order: `A1 / Go / 3`, `B2 / 世界 / 0`, `A1 / duplicate / 9`, and `C3 / 🌞 / 2`. The first accepted resource for an identifier is authoritative.

## What You Need to Do

- [x] Produce an ordered catalog containing only accepted resources and report the total input, accepted, and duplicate counts.
- [x] Report the lookup result and presence state for identifiers `B2` and `Z9` without confusing zero priority with absence.
- [x] Explain the data-structure choices used for ordered records, keyed retrieval, and uniqueness.

## Rules and Constraints

The resource model must be a named struct with identifier, label, and priority fields. Preserve input order in the accepted catalog, never derive report order from map iteration, and derive all counts from the supplied data. The final explanation must address why a slice and a map serve different roles in the design.

## Expected Result

The catalog contains three resources in their original accepted order and reports one duplicate. The `B2` lookup is present with priority zero, while `Z9` is absent even though its lookup value contains zero values.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
