# Exercise 4 — Selective Counter

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** Three-part `for`, `break`, `continue`, loop progress

## Goal

Build a bounded number scan that skips excluded candidates and stops at a defined upper rule.

## Scenario

A batch accepts odd ticket numbers from a numeric window. The scan must preserve predictable progress even when an iteration is skipped.

## What You Need to Do

- [x] Inspect ticket numbers from `1` upward, without inspecting any number above `12`.
- [x] Skip every even ticket and print each accepted odd ticket in encounter order.
- [x] Stop the scan before accepting any ticket greater than `9`, then print the accepted count and their total.

## Rules and Constraints

- Use a three-part `for` for the numeric window.
- Use `continue` for excluded tickets and `break` for the early stopping rule.
- Keep the update behavior in the loop header so every continued iteration still makes progress.

## Expected Result

The report contains only eligible odd tickets in ascending order, stops when the first out-of-policy odd ticket is encountered, and finishes with a correct count and sum.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
