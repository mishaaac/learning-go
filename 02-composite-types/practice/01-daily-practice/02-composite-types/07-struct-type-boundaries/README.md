# Exercise 7 — Struct Type Boundaries

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Named structs, anonymous structs, keyed literals, comparison, conversion, assignability

## Goal

Demonstrate how identical-looking struct shapes behave across named and anonymous types.

## Scenario

An internal job record and an export record carry the same comparable fields but represent different domain roles. A one-off local snapshot uses the same shape without introducing another reusable name.

## What You Need to Do

- [x] Define two distinct named struct types with matching `string`, `int`, and `bool` fields, and initialize values with keyed literals.
- [x] Convert explicitly between the named types, assign a compatible named value to an anonymous struct, and report valid equality results.
- [x] Explain why the two named types deserve separate names and why direct assignment between them is not allowed.

## Rules and Constraints

- Keep every compared field type comparable.
- Do not attempt an invalid direct comparison between values of different named types.
- Use the anonymous struct only for a local, one-off value and include a brief comment justifying that choice.

## Expected Result

The program compiles and shows equal field content after the required conversion and assignment. The explanation separates structural similarity, type identity, and the decision to name a reusable domain shape.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
