# Exercise 2 — Nested Control Prediction

> **Difficulty:** Foundation
> **Estimated time:** 20 minutes
> **Topics:** Three-part `for`, `for-range`, `switch`, `break`, `continue`, loop labels

## Goal

Predict control flow when a `switch` and labeled control statements are nested inside two loops.

## Scenario

A scoring pass treats some values as special while processing three rows. Determine which additions and report lines are reached before running the code.

```go
package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	total := 0

outer:
	for row := 1; row <= 3; row++ {
		for _, value := range values {
			switch {
			case value == 2:
				continue
			case row == 2 && value == 3:
				continue outer
			case value == 4:
				break
			default:
				total += row * value
			}
		}
		fmt.Println("row", row, total)
	}

	fmt.Println("total", total)
}
```

## What You Need to Do

- [x] Predict every output line and its order before running the snippet.
- [x] Trace which statement receives control after each `continue`, `continue outer`, and `break`.
- [x] Explain why the `break` in the `switch` does not terminate either loop.

## Rules and Constraints

- Record the full prediction before using the compiler.
- Treat the label as targeting the outer `for`, not the inner loop or the `switch`.
- Include any row report that remains reachable after the inner loop finishes.

## Expected Result

Your prediction identifies the rows that print, the additions that contribute to the final total, and the different targets of the three control statements.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
