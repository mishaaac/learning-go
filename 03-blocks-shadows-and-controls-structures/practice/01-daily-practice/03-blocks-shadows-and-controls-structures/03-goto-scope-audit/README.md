# Exercise 3 — Goto Scope Audit

> **Difficulty:** Foundation
> **Estimated time:** 20 minutes
> **Topics:** `goto`, labels, blocks, declaration scope, compile-time restrictions

## Goal

Decide whether each jump is legal in Go and justify the decision using block and declaration scope.

## Scenario

A code review contains three proposed uses of `goto`. Review each program independently before asking the compiler for evidence.

### Program A

```go
package main

import "fmt"

func main() {
	value := -2
	if value < 0 {
		goto done
	}

	fmt.Println("processed", value)

done:
	fmt.Println("finished")
}
```

### Program B

```go
package main

import "fmt"

func main() {
	goto done
	message := "ready"

done:
	fmt.Println(message)
}
```

### Program C

```go
package main

import "fmt"

func main() {
	goto inside

	if true {
	inside:
		fmt.Println("inside")
	}
}
```

## What You Need to Do

- [x] Predict which programs compile and which the compiler rejects.
- [x] For every rejected program, identify the scope rule violated by the jump.
- [x] For every valid program, predict its exact output and assess whether a simpler control structure would communicate the intent better.

## Rules and Constraints

- Evaluate each program as a separate source file.
- Make all predictions before compiling any program.
- Base legality on jump and scope rules, not only on whether execution would appear safe at runtime.

## Expected Result

Your review classifies all three programs, explains every invalid jump precisely, and distinguishes language legality from design quality.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
