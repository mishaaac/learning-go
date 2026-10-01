# Problem 1 — Inventory State Report

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Predeclared types, zero values, declarations, numeric operations, strings, runes

## Goal

Create a warehouse intake report from a set of typed values and derived results.

## Scenario

The North Dock received shipment `A` with 1,250 units. Eighteen units were damaged, and each accepted unit has an average weight of 2.75 kilograms. Inspection has not started, and no optional note has been entered.

## What You Need to Do

- [x] Report the facility name, shipment marker, delivered count, damaged count, inspection state, and optional note.
- [x] Derive and report the accepted-unit count and its total weight.
- [x] Make the false inspection state and empty optional note unambiguous in the output.

## Rules and Constraints

Use `int` for unit counts, `float64` for weight, `rune` for the shipment marker, and zero-value declarations for the inspection state and optional note. Include both a `var` declaration and a short declaration. Derived results must come from the source values rather than from hard-coded answers.

## Expected Result

The report contains all requested fields, shows 1,232 accepted units and their calculated total weight, and makes both zero values visible.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
