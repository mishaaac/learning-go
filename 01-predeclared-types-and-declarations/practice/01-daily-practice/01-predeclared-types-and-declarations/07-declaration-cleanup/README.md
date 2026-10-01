# Exercise 7 — Declaration Cleanup

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** Unused variables, assignment, short declarations, identifier naming, constants

## Goal

Repair declaration errors and improve non-idiomatic names without hiding unused values.

## Scenario

A small status program no longer compiles after an unfinished edit. Restore its behavior and leave its declarations clear enough for another Go developer to review.

```go
package main

import "fmt"

const MAX_RETRIES = 3

func main() {
    count := 10
    count := 20
    status_message := "ready"

    fmt.Println(count)
}
```

## What You Need to Do

- [x] Identify every declaration-related compile failure in the snippet before editing it.
- [x] Make the final program print the retry limit, current count, and status message.
- [x] Rename non-idiomatic identifiers according to Go's mixed-capitalization convention.

## Rules and Constraints

- Preserve the final count value of `20` and the status text `ready`.
- Do not use the blank identifier solely to silence an unused-variable error.
- Do not remove a required value just to make the compiler succeed.

## Expected Result

The program compiles without unused local variables, prints all three required values, and uses idiomatic names for both variables and constants.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
