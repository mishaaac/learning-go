# Problem 3 — Repair Type Boundaries

> **Difficulty:** Hard
> **Estimated time:** 35 minutes
> **Topics:** Numeric types, conversions, booleans, short declarations, representability, unused variables

## Goal

Diagnose every declaration and type error in a broken program, then restore its intended behavior.

## Scenario

A batch-progress report contains several independent compile failures. The repaired program must preserve the intended data instead of merely removing problematic declarations.

```go
package main

import "fmt"

const batchSize int = 10

func main() {
    var completed int32 = 7
    ratio := completed / batchSize
    var enabled bool = 1
    var code byte = 300
    label := "ready"
    label := "complete"
    unusedMessage := "batch processed"

    fmt.Println(completed, ratio, enabled, code, label)
}
```

## What You Need to Do

- [x] List every compile failure and explain the language rule involved before changing the code.
- [x] Produce a valid report containing the completed count, a fractional progress ratio, a boolean that reports whether the completed count is nonzero, the exact code value `300`, the final label `complete`, and the message `batch processed`.
- [x] Explain each final type choice and every explicit conversion used by the repaired program.

## Rules and Constraints

Do not change the numeric values `7`, `10`, or `300`. Do not comment out required data, use the blank identifier to suppress an error, or replace a derived result with a hard-coded answer. The progress ratio must retain its fractional part.

## Expected Result

The final program compiles, reports all required values, preserves `300` exactly, produces a fractional ratio for seven completed items out of ten, and reports the nonzero-count condition as true. The written diagnosis covers all original failures.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
