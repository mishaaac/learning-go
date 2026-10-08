# `goto` in Go

`goto` transfers control to a label within the same function. Go keeps the feature but restricts jumps that would violate scope, so it should remain an exceptional tool rather than ordinary control flow.

## 1. What I Need to Understand

- The syntax is `goto label`, with `label:` marking the destination.
- The destination label must be in the same function.
- A jump cannot skip declarations whose variables would be in scope at the destination.
- A jump cannot enter a block from outside it.
- Labeled `break` and `continue` are usually clearer for nested loops.

## 2. Key Concepts

| Jump | Allowed? |
|---|---:|
| To a valid label in the same function | Sometimes |
| Across code without violating scope | Possibly |
| Over a variable declaration into its scope | No |
| From outside into an inner block | No |
| To another function | No |

## 3. How It Works in Go

```go
value := readValue()
if value < 0 {
    goto done
}

process(value)

done:
fmt.Println("finished")
```

The following shape is illegal because the jump skips a declaration that would be in scope at the label:

```go
goto done
result := compute()
done:
fmt.Println(result)
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a label does not remove normal scope rules; it only names a possible control-flow destination.

Jumping from outside an `if` or other block to a label inside that block is invalid. A `goto` also cannot cross function boundaries.

An exceptional use can converge several complex paths on one common cleanup or final-processing section, avoiding artificial flags or substantial duplication. First consider returning, extracting a function, or using labeled `break`/`continue`.

> ⚠️ **Important correction:** “avoid `goto`” is a design guideline, not a language ban. The compiler accepts legal jumps, and the standard library contains rare cases where a constrained jump makes control flow simpler.

## 5. Quick Cheat Sheet

- Syntax → `goto label`
- Destination → `label:`
- Label must be in the same function
- Cannot jump over a declaration into its scope
- Cannot jump into an inner block
- Cannot jump to another function
- For nested loops, prefer labeled `break`/`continue`
- Consider `return` or function extraction first
- Rare use → converge on common final logic
- Legal does not automatically mean readable

### Test Yourself

1. Why is jumping over a declaration illegal when the variable is in scope at the destination?
2. When is labeled `break` clearer than `goto`?
3. What characteristics might justify a rare jump to common final logic?
