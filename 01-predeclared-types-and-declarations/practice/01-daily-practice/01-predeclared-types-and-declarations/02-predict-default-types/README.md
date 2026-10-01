# Exercise 2 — Predict Default Types

> **Difficulty:** Foundation
> **Estimated time:** 15 minutes
> **Topics:** Untyped literals, default types, contextual types, type inference

## Goal

Predict when Go uses a literal's default type and when a declaration supplies a different contextual type.

## Scenario

A teammate wants to know the concrete types produced by several declarations. You must reason about the code before asking the compiler to confirm your answer.

```go
package main

import "fmt"

func main() {
    var count = 42
    var ratio = 3.5
    var letter = 'G'
    var message = "Go"
    var small int8 = 42

    fmt.Printf("%T %T %T %T %T\n", count, ratio, letter, message, small)
    fmt.Println(count, ratio, letter, message, small)
}
```

## What You Need to Do

- [x] Before running the code, write down the exact type sequence expected on the first output line.
- [x] Predict the values on the second output line, including how the rune will be displayed.
- [x] Run the snippet, compare it with your prediction, and explain why `small` differs from the default integer type.

## Rules and Constraints

- Record the complete prediction before copying or running the snippet.
- Do not change the declarations until after the prediction has been checked.
- Base the explanation on default types and contextual typing, not on memorized output alone.

## Expected Result

Your prediction identifies all five concrete types and values. After running the program, you can explain which declarations used default types and which one used an explicit contextual type.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
