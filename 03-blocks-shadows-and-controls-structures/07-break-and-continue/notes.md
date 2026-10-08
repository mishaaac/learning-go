# `break` and `continue` in Go

`break` ends a loop, while `continue` skips the rest of the current iteration and begins the next one. Used carefully, both statements make exit conditions explicit and reduce unnecessary nesting.

## 1. What I Need to Understand

- In a loop, `break` exits the nearest applicable enclosing structure.
- `continue` starts the next iteration of the nearest enclosing `for`.
- In a three-part loop, `continue` still proceeds through the post statement before retesting.
- `for { ...; if !condition { break } }` can model a `do/while` pattern.
- Early `continue` statements can keep independent cases aligned and readable.

## 2. Key Concepts

| Statement | Effect |
|---|---|
| `break` | Leave the current loop completely |
| `continue` | Skip the remaining body for this iteration |
| `return` | Leave the entire function |
| Labeled form | Target an enclosing loop explicitly |

## 3. How It Works in Go

```go
for i := 1; i <= 10; i++ {
    if i == 8 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

This prints odd values below `8`. A body-first repetition can be written as:

```go
for {
    doWork()
    if !shouldContinue() {
        break
    }
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `break` ends the loop; `continue` ends only the current iteration.

When translating `do { ... } while (condition)`, the Go exit test is negated because it asks when to stop:

```go
if !condition {
    break
}
```

Several early `continue` checks can be clearer than deeply nested `if`/`else` branches, but excessive jumps may also fragment the flow. Each statement should make the next execution point obvious.

An unlabeled `break` inside a `switch` nested in a loop exits the `switch`, not the loop; labels are covered separately.

## 5. Quick Cheat Sheet

- `break` → exit the loop
- `continue` → next iteration
- `return` → exit the function
- `continue` skips the remaining body
- Three-part loop `continue` proceeds to `post`
- Body-first pattern → `for { ...; if !condition { break } }`
- Negation expresses when to stop
- Early `continue` can reduce nesting
- Unlabeled control targets the nearest applicable structure
- Use labels when an outer loop must be targeted

### Test Yourself

1. In a three-part loop, what happens after `continue` and before the next condition check?
2. Why is the condition negated when modeling `do/while` with `break`?
3. How can `continue` reduce nesting without changing the result?
