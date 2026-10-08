# Exercise 1 — Array and Slice View

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Arrays, indexed literals, slices, `len`, `cap`, shared storage

## Goal

Predict how an indexed array literal and a slice view determine their lengths, capacities, and visible values.

## Scenario

A compact diagnostic creates a partially initialized array and exposes part of it as a slice. The slice is then used to update the shared storage.

```go
package main

import "fmt"

func main() {
	readings := [...]int{2: 7, 4: 9}
	window := readings[1:4]

	window[1] = 8

	fmt.Println(len(readings), len(window), cap(window))
	fmt.Println(readings)
	fmt.Println(window)
}
```

## What You Need to Do

- [x] Predict the exact three output lines before running the snippet.
- [x] Explain how the highest explicit index determines the inferred array length.
- [x] Explain why changing the slice also changes one element in the array.

## Rules and Constraints

- Record the complete prediction before copying or running the code.
- Base valid indices on length, not capacity.
- Do not change the slice bounds until the original prediction has been checked.

## Expected Result

Your prediction gives the correct array length, slice length, slice capacity, and values after the update. You can identify which storage is shared by the two values.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
