# Exercise 5 — Short Declarations and Shadowing

> **Difficulty:** Intermediate
> **Estimated time:** 20 minutes
> **Topics:** `var`, `:=`, assignment, block scope, shadowing

## Goal

Predict which binding a short declaration creates and distinguish redeclaration from assignment.

## Scenario

A counter is changed inside a nested block, but a later report may still see the original value. You need to identify the binding used at each point.

```go
package main

import "fmt"

func main() {
    count := 10

    {
        count, label := 20, "inner"
        fmt.Println(count, label)
    }

    fmt.Println(count)
}
```

## What You Need to Do

- [x] Predict both output lines and identify whether the two occurrences of `count` refer to the same variable.
- [x] Run the original snippet and compare the result with your prediction.
- [x] Create a second valid version in which the nested block changes the outer `count` while `label` remains available for the inner report.

## Rules and Constraints

- Record the prediction before running the original snippet.
- Keep the nested block in the second version.
- Use `:=` only where it declares at least one new non-blank variable, and use `=` only with names already declared in the relevant scope.

## Expected Result

The original version demonstrates shadowing. The second version prints an updated outer count after the block and follows Go's declaration and assignment rules.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
