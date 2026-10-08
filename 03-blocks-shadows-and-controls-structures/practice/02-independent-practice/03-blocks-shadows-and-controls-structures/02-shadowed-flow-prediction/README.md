# Problem 2 — Shadowed Flow Prediction

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Blocks, variable shadowing, `if` initial statements, three-part `for`, `switch`, `break`, `continue`

## Goal

Predict the observable behavior of nested scopes and control statements without executing the program first.

## Scenario

The same identifier is used with different types in separate blocks, while loop control changes which print statements are reached. Review the complete program as one unit.

```go
package main

import "fmt"

func main() {
	status := "ready"
	total := 0

	for index := 1; index <= 4; index++ {
		if status := index % 2; status == 0 {
			total += index
			continue
		}

		switch index {
		case 3:
			status := "paused"
			fmt.Println(status, total)
			break
		default:
			total += index
		}

		fmt.Println(status, total)
	}

	fmt.Println(status, total)
}
```

## What You Need to Do

- [x] Record the exact output, including line order, before running the program.
- [x] Explain which declaration each use of `status` resolves to.
- [x] Account for every change to `total` and every skipped print statement.
- [x] Explain the target and effect of both `break` and `continue` in this program.

## Rules and Constraints

Do not compile, run, or alter the program until the complete prediction and explanation have been written. Evaluate scope according to the block containing each declaration.

## Expected Result

The prediction includes every reachable output line and a consistent scope and control-flow explanation for each value displayed.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
