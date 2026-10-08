# Exercise 1 — Shadowed Status Prediction

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Blocks, scope, short declarations, variable shadowing

## Goal

Predict how nested declarations and assignments affect values in their own blocks and in enclosing blocks.

## Scenario

A small job tracker reuses familiar names inside nested status checks. Analyze the program before using the compiler to confirm your reasoning.

```go
package main

import "fmt"

func main() {
	status := "queued"
	count := 1

	if count > 0 {
		status, count := "running", count+1
		fmt.Println(status, count)

		if count == 2 {
			status = "checked"
			note := status
			fmt.Println(note, count)
		}

		fmt.Println(status, count)
	}

	fmt.Println(status, count)
}
```

## What You Need to Do

- [x] Write down the exact four output lines before running the snippet.
- [x] Identify which `status` and `count` declarations each print statement observes.
- [x] Explain why the final line may differ from the lines printed inside the outer `if` block.

## Rules and Constraints

- Record your prediction before copying or running the code.
- Distinguish a short declaration in a nested block from assignment to an existing identifier.
- Account for the scope of `note` without trying to use it outside its block.

## Expected Result

Your prediction gives all four lines in order and traces each displayed value to the declaration or assignment that controls it at that point.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
