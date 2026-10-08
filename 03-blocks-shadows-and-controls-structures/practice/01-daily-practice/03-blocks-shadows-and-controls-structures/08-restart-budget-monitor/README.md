# Exercise 8 — Restart Budget Monitor

> **Difficulty:** Applied
> **Estimated time:** 35 minutes
> **Topics:** `for-range`, expression `switch`, labeled `break`, `continue`, nested control structures

## Goal

Build a command monitor that stops an enclosing loop from inside a `switch` when its restart budget is exhausted.

## Scenario

Process these commands in order: `ok`, `retry`, `ignored`, `retry`, `retry`, `ok`, `fatal`. The monitor permits at most three retry commands in one run.

## What You Need to Do

- [x] Count processed `ok` and `retry` commands while skipping `ignored` without further work for that iteration.
- [x] Stop the entire command scan immediately when the third retry is counted or when `fatal` is encountered.
- [x] Print the stop reason and a final summary, ensuring commands after the stopping command have no effect.

## Rules and Constraints

- Use an expression `switch` to dispatch command values.
- Use labeled control to exit the loop from inside the switch; an unlabeled `break` is not sufficient.
- Count the command that exhausts the retry budget before stopping.

## Expected Result

The third retry ends this data set before the later `ok` and `fatal` commands are processed. The summary reflects only commands reached through the stopping point and names the budget as the reason.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
