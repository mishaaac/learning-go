# Exercise 5 — Unicode Offset Report

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** String `for-range`, byte offsets, runes, loop counters

## Goal

Build a report that distinguishes UTF-8 byte positions from the number of decoded runes.

## Scenario

A text diagnostic receives the label `Go→Lima🌞`. Operators need to see where each decoded rune begins in the original string.

## What You Need to Do

- [x] Print one line per decoded rune containing its starting byte offset, rune value, and displayed character.
- [x] Count the runes during traversal and report that count after the loop.
- [x] Report the string's total byte length and make the two measurement units clear in the output.

## Rules and Constraints

- Traverse the text with string `for-range`.
- Do not assume successive rune offsets differ by one.
- Do not convert the complete string to `[]rune` to perform the traversal or count.

## Expected Result

The output preserves text order, shows byte-offset jumps where UTF-8 uses multiple bytes, and reports separate byte and rune totals.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
