# Problem 2 — Predict Storage Effects

> **Difficulty:** Intermediate
> **Estimated time:** 30 minutes
> **Topics:** Arrays, sub-slices, capacity limits, `append`, `copy`, `clear`, backing storage

## Goal

Predict the complete effect of shared storage, capacity-restricted growth, copying, and clearing.

## Scenario

A storage experiment creates two views of the same array, protects one view from in-place growth, and preserves an independent copy before later changes.

```go
package main

import "fmt"

func main() {
	values := [5]int{10, 20, 30, 40, 50}
	shared := values[1:3]
	guarded := values[1:3:3]
	clone := make([]int, len(shared))
	copy(clone, shared)

	shared[0] = 99
	shared = append(shared, 77)
	guarded = append(guarded, 88)
	clear(shared[:2])

	fmt.Println(values)
	fmt.Println(shared)
	fmt.Println(guarded)
	fmt.Println(clone)
}
```

## What You Need to Do

- [x] Record the exact four output lines before compiling or running the program.
- [x] Identify which final values still share storage and which have independent storage.
- [x] Explain every visible mutation, including the different effects of the two append operations and the final clear.

## Rules and Constraints

Do not modify or run the snippet until the complete prediction is recorded. Do not assume any unspecified capacity growth factor. The explanation must account for the state of every printed element, not only the final collections as a whole.

## Expected Result

The written prediction matches all four output lines exactly. The explanation correctly relates each result to slice length, capacity, backing storage, copying, and zeroing.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
