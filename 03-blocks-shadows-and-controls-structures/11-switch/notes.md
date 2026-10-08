# `switch` in Go

An expression `switch` selects the first matching `case` and stops automatically after that clause. Go supports multiple case expressions, local initialization, and explicit `fallthrough`, but normal cases do not require `break`.

## 1. What I Need to Understand

- `switch expression` compares one expression with case expressions.
- Cases are considered in source order, and the first match runs.
- Several values can share one case by separating them with commas.
- Cases do not fall through automatically.
- An optional initial statement can declare a value scoped to the entire switch.

## 2. Key Concepts

| Feature | Behavior |
|---|---|
| `case 1, 2, 3:` | Any listed value selects the clause |
| `default:` | Runs when no case matches |
| Empty case | Performs no action |
| `fallthrough` | Continues unconditionally into the next clause body |
| `break` | Ends the nearest enclosing `switch`, `for`, or `select` |

## 3. How It Works in Go

```go
switch size := len(word); size {
case 1, 2, 3, 4:
    fmt.Println("short")
case 5:
    description := "exact"
    fmt.Println(description)
case 6, 7, 8, 9:
    // Intentionally do nothing.
default:
    fmt.Println("long")
}
```

`size` is available in the switch expression and every clause, while `description` belongs only to its case block.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** an empty case does nothing; it does not continue into the next case.

`fallthrough` must be the final non-empty statement in a non-final case. It transfers control to the next clause body without checking whether that next case expression matches, so it should be rare and deliberate.

> ⚠️ **Important correction:** an expression switch is not limited to integers, but its expression and case values must be valid for equality comparison. Channels are comparable; slices, maps, and functions are not.

Inside a `switch` nested in a loop, an unlabeled `break` exits the switch. Use `break loopLabel` to exit an explicitly labeled outer loop.

## 5. Quick Cheat Sheet

- Form → `switch expression { ... }`
- No expression parentheses required
- First matching case runs
- Multiple values → comma-separated case list
- `default` handles no match
- Cases do not fall through automatically
- Empty case → no action
- `fallthrough` explicitly enters the next body
- Initial statement → `switch statement; expression`
- Case declarations stay in their case block
- Unlabeled `break` exits the nearest structure

### Test Yourself

1. Why is `break` normally unnecessary at the end of a Go case?
2. How does `fallthrough` differ from matching the next case normally?
3. Why does `break` inside a switch nested in a loop not necessarily exit the loop?
