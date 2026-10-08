# Exercise 4 — Append Storage Boundaries

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** Sub-slices, backing arrays, full slice expressions, `append`, capacity

## Goal

Predict when appending to a sub-slice reuses shared storage and when a capacity limit forces separate storage.

## Scenario

Two views begin over the same prefix. One keeps spare capacity while the other deliberately restricts its growth boundary.

```go
package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40}
	shared := original[:2]
	limited := original[:2:2]

	shared = append(shared, 99)
	limited = append(limited, 77)
	limited[0] = 5

	fmt.Println(original)
	fmt.Println(shared)
	fmt.Println(limited)
}
```

## What You Need to Do

- [x] Predict the exact three output lines before running the snippet.
- [x] Identify which append can reuse the original backing array and the visible consequence of that reuse.
- [x] Explain why the later indexed update through `limited` does or does not affect `original`.

## Rules and Constraints

- Record the complete prediction before copying or running the code.
- Do not rely on a particular capacity growth factor after allocation.
- Distinguish sharing existing elements from the storage choice made during append.

## Expected Result

Your prediction correctly traces both appends and the final indexed update. Your explanation uses length, capacity, and backing storage rather than treating all slices as independent containers.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
