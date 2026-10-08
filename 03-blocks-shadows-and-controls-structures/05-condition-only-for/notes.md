# Condition-Only `for` Loops in Go

A `for` with only a condition is Go's equivalent of a conventional `while` loop. Initialization and updates live outside or inside the body, while the header states only whether iteration should continue.

## 1. What I Need to Understand

- The syntax is `for condition { ... }`.
- The condition is checked before every iteration.
- The condition must produce a `bool`.
- No semicolons are written when only the condition remains.
- The body must normally change state so the condition can eventually become false.

## 2. Key Concepts

| Three-part loop | Condition-only loop |
|---|---|
| `for init; condition; post` | `for condition` |
| Init usually in header | Initialization occurs beforehand |
| Update usually in header | Update occurs in the body |
| Uses semicolons | Uses no semicolons |

## 3. How It Works in Go

```go
value := 1

for value < 100 {
    fmt.Println(value)
    value *= 2
}
```

The values printed are `1`, `2`, `4`, `8`, `16`, `32`, and `64`. After the update produces `128`, the next condition check is false.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `for condition` checks before the body, so the body may execute zero times.

This form is clearer than an empty three-part clause:

```go
for value < limit {
    value = nextValue(value)
}
```

If the body never changes anything that influences the condition, the loop may be infinite. Updates hidden in several branches also make termination harder to verify.

Use a three-part `for` when initialization and a regular post step are naturally part of the loop header.

## 5. Quick Cheat Sheet

- Syntax → `for condition { ... }`
- Equivalent role to `while`
- Condition checked before the body
- Body may run zero times
- Condition must be boolean
- No header semicolons
- Initialize before the loop
- Update normally occurs in the body
- State must progress toward termination
- Prefer three-part `for` for regular counters

### Test Yourself

1. Why can a condition-only loop execute zero times?
2. What responsibility moves into the body when `init` and `post` are removed?
3. When is a three-part `for` clearer than a condition-only form?
