# Problem 3 — Repair Composite Boundaries

> **Difficulty:** Hard
> **Estimated time:** 45 minutes
> **Topics:** Nil maps, slice length and capacity, `copy`, struct comparability, slice equality

## Goal

Diagnose every compile-time, runtime, and incorrect-result defect in a program that crosses several composite-type boundaries, then restore its intended behavior.

## Scenario

The following report is intended to store three readings, preserve an independent snapshot, and compare two reports by content. It currently cannot satisfy that contract.

```go
package main

import "fmt"

type Report struct {
	Name   string
	Values []int
}

func main() {
	var totals map[string]int
	totals["ready"] = 1

	values := make([]int, 0, 3)
	values[0] = 10
	values = append(values, 20, 30)

	snapshot := make([]int, 0, len(values))
	copied := copy(snapshot, values)

	current := Report{Name: "batch", Values: values}
	archived := Report{Name: "batch", Values: snapshot}

	fmt.Println(current == archived)
	values[0] = 99
	fmt.Println(totals["ready"], copied)
	fmt.Println(current.Values)
	fmt.Println(archived.Values)
}
```

## What You Need to Do

- [x] List all independent defects before editing and classify each one as a compile-time failure, runtime failure, or incorrect-result defect.
- [x] Repair the program so it stores the intended readings, preserves an independent snapshot, and determines whether the complete report contents were equal before the later mutation.
- [x] Explain why the original report comparison is invalid and justify the final content-comparison behavior.

## Rules and Constraints

Preserve the `Report` fields, the values `10`, `20`, `30`, and `99`, the `ready` count, and the final mutation. Do not remove required data, comment out failing statements, use the blank identifier to hide a result, or hard-code the copied count or equality result.

## Expected Result

The repaired program reports equal content before the mutation, a ready count of `1`, and a copied count of `3`. Its final current values are `[99 20 30]`, while the archived values remain `[10 20 30]`. The written diagnosis covers defects that compilation alone cannot reveal.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
