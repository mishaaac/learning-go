# Choosing Between `if` and `switch` in Go

Both an `if` chain and an expressionless switch can evaluate boolean conditions. The choice is mainly communicative: a switch presents related alternatives as one decision, while `if` is often clearer for separate checks or asymmetric flow.

## 1. What I Need to Understand

- `if`/`else if` and `switch {}` can express equivalent condition sequences.
- A switch signals that its cases are related alternatives.
- An `if` is natural for one condition, guard clauses, or less-related checks.
- Both forms choose the first successful branch when written as one chain.
- Refactor when neither form can present the logic coherently.

## 2. Key Concepts

| Situation | Usually clearer |
|---|---|
| One primary condition | `if` |
| Early exit or guard | `if` |
| Several related alternatives | `switch` |
| Clear remaining case | `default` in `switch` |
| Unrelated responsibilities | Separate `if` statements or refactoring |

## 3. How It Works in Go

```go
switch {
case value%3 == 0 && value%5 == 0:
    fmt.Println("FizzBuzz")
case value%3 == 0:
    fmt.Println("Fizz")
case value%5 == 0:
    fmt.Println("Buzz")
default:
    fmt.Println(value)
}
```

These cases answer one related question about the same value, so the switch makes the alternatives visible without repeated `continue` statements.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a switch is not automatically better merely because there are several conditions; the cases should belong to one conceptual decision.

An `if` guard is often more direct:

```go
if err != nil {
    return err
}
```

Case order matters when predicates overlap. In FizzBuzz, divisibility by both values must be checked before divisibility by either one alone.

Independent `if` statements may all run; an `if`/`else if` chain or switch selects only one branch. Choose according to the required behavior, not only visual style.

## 5. Quick Cheat Sheet

- One condition → usually `if`
- Guard clause → `if`
- Related alternatives → `switch`
- Expressionless switch handles boolean predicates
- First matching switch case runs
- `default` makes the remainder explicit
- Overlapping cases require careful order
- Separate `if` statements can all execute
- `if`/`else if` selects one branch
- Unrelated cases suggest `if` or refactoring

### Test Yourself

1. Why does FizzBuzz communicate well as an expressionless switch?
2. How do separate `if` statements differ from an `if`/`else if` chain?
3. What indicates that a proposed switch should instead be refactored?
