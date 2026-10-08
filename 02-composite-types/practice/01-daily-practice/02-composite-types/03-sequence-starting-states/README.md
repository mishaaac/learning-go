# Exercise 3 — Sequence Starting States

> **Difficulty:** Foundation
> **Estimated time:** 20 minutes
> **Topics:** Arrays, slice declarations, zero values, `make`, `append`, `len`, `cap`

## Goal

Build a small program that demonstrates how declaration choices give arrays and slices different starting states.

## Scenario

A batch tool needs fixed slots, optional results, known labels, indexable output space, and an initially empty collection with reserved capacity. Each sequence should begin in a state that matches how it will be used.

## What You Need to Do

- [x] Represent four fixed batch slots with an array, initialize only selected positions, and display the zero-valued positions.
- [x] Demonstrate a nil slice, an empty non-nil slice, a slice literal with known labels, and a three-element slice ready for indexed writes.
- [x] Collect two result values in an initially empty slice that reserves capacity for five values, then report its values, length, and capacity.

## Rules and Constraints

- Choose `var`, literals, or `make` according to the required starting state.
- Do not create placeholder elements in the result collection.
- Include a short comment in the completed program explaining why reserved capacity does not create valid indices.

## Expected Result

The output makes every starting state observable, including nil status, zero values, lengths, and capacities. The result collection contains exactly two values and no leading placeholders.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
