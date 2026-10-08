# Expressionless `switch` in Go

An expressionless switch places a boolean condition in each `case` instead of comparing one value with several alternatives. It behaves like switching on `true` and is useful for related ranges or predicates.

## 1. What I Need to Understand

- The syntax is `switch { case condition: ... }`.
- Each case expression must be boolean.
- Cases are checked in order, and only the first true case runs.
- An initial statement may calculate a value used by every case.
- Use an expression switch instead when every case is equality against the same value.

## 2. Key Concepts

| Question being expressed | Better form |
|---|---|
| “Which value equals `x`?” | `switch x` |
| “Which related condition is true?” | `switch {}` |
| “No previous condition matched” | `default` |

An expressionless switch is conceptually equivalent to `switch true`.

## 3. How It Works in Go

```go
switch length := len(word); {
case length < 5:
    fmt.Println("short")
case length > 10:
    fmt.Println("long")
default:
    fmt.Println("medium")
}
```

`length` is computed once and remains available in all clauses of that switch.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `switch { case x == 2: ... }` works, but `switch x { case 2: ... }` expresses repeated equality more directly.

Order matters when conditions overlap. Put a more specific condition before a broader condition that would also match:

```go
switch {
case value%3 == 0 && value%5 == 0:
    fmt.Println("FizzBuzz")
case value%3 == 0:
    fmt.Println("Fizz")
}
```

An expressionless switch is clearest when its conditions are different predicates about the same decision, not unrelated checks collected merely because the syntax allows it.

## 5. Quick Cheat Sheet

- Form → `switch { ... }`
- Conceptually → `switch true`
- Cases contain boolean conditions
- First true case runs
- Order matters for overlapping conditions
- `default` handles the remainder
- Initial statement is allowed
- Different related predicates → expressionless switch
- Same expression compared with values → expression switch
- Avoid grouping unrelated conditions

### Test Yourself

1. Why must the most specific overlapping condition come first?
2. When does `switch x` communicate intent better than `switch {}`?
3. What relationship should the conditions in an expressionless switch have?
