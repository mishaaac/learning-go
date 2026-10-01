# Problem 2 — Predict Constant Behavior

> **Difficulty:** Intermediate
> **Estimated time:** 25 minutes
> **Topics:** Typed constants, untyped constants, default types, contextual types, explicit conversions

## Goal

Predict the concrete types and values produced when constants meet different declaration contexts.

## Scenario

A short diagnostic relies on constant flexibility, integer division, and explicit conversion. Its output must be reviewed before it is allowed into a larger program.

```go
package main

import "fmt"

const flexible = 255
const fixed int = 3

func main() {
    var octet byte = flexible
    var measurement float64 = flexible
    quotient := 7 / 2
    precise := float64(7) / 2
    converted := float64(fixed)

    fmt.Printf("%T %T %T %T %T\n", octet, measurement, quotient, precise, converted)
    fmt.Println(octet, measurement, quotient, precise, converted)
}
```

## What You Need to Do

- [x] Record the exact two output lines before compiling or running the program.
- [x] Run the program only after the prediction is complete and compare every field.
- [x] Explain the type and value of each expression, including both division results and the explicit conversion.

## Rules and Constraints

Do not modify the snippet before checking the original prediction. The explanation must distinguish contextual typing, default typing, typed constants, and untyped constants.

## Expected Result

The written prediction matches both output lines exactly, and the explanation accounts for every concrete type and numeric value without relying only on observed output.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
