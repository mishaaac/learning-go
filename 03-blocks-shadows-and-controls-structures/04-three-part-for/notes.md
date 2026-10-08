# The Three-Part `for` Loop in Go

Go's three-part `for` makes initialization, continuation, and update rules visible in one header. It is useful when an iteration has clear numeric or positional boundaries.

## 1. What I Need to Understand

- The form is `for init; condition; post { ... }`.
- `init` runs once before the first condition check.
- `condition` is checked before every iteration and must be boolean.
- `post` runs after each completed iteration before the next condition check.
- Variables declared in `init` are scoped to the loop, including its body.

## 2. Key Concepts

| Part | Example | Timing |
|---|---|---|
| Init | `i := 0` | Once, before looping |
| Condition | `i < 10` | Before every iteration |
| Post | `i++` | After every completed iteration |

One or more parts may be empty, but the semicolons remain when using a `for` clause.

## 3. How It Works in Go

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

```text
init → condition ─false→ finish
          │true
          ▼
         body
          ▼
         post ──────────┘
```

Initialization may be outside the loop:

```go
i := 0
for ; i < 3; i++ {
    fmt.Println(i)
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `init` executes once; `post` executes repeatedly after iterations.

`init` and `post` are simple statements. A `var` declaration is not valid directly in `init`; use a short declaration or declare the variable beforehand.

> ⚠️ **Important correction:** the `post` statement cannot be a short variable declaration. `i++`, assignments, and function calls may be valid, but `j := i + 1` is not allowed there.

If the update logic is complex, omit `post` and update inside the body. Be certain every path still makes progress, or the loop may never terminate.

## 5. Quick Cheat Sheet

- Form → `for init; condition; post`
- No surrounding parentheses
- Init runs once
- Condition runs before each iteration
- Post runs after each completed iteration
- Condition must produce `bool`
- Init variable lives for the whole loop
- Parts may be omitted
- Semicolons remain in a three-part clause
- `var` is not an init simple statement
- Post cannot use `:=`

### Test Yourself

1. In what order do condition, body, and post execute?
2. Why is `j := i + 1` invalid as the post statement?
3. What risk appears when update logic is moved into the loop body?
