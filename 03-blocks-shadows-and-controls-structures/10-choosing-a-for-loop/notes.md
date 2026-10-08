# Choosing the Right `for` Loop in Go

Go uses one keyword for several iteration patterns. The clearest form is the one that directly expresses whether the code traverses content, follows explicit bounds, repeats on a condition, or exits from inside the body.

## 1. What I Need to Understand

- Prefer `range` when traversing all elements of a supported value.
- Use a three-part `for` when start, end, and update boundaries matter.
- Use `for condition` when repetition follows a changing condition.
- Use `for {}` when exit decisions naturally occur inside the body.
- For strings, prefer `range` when processing Unicode code points rather than bytes.

## 2. Key Concepts

| Need | Natural form |
|---|---|
| Traverse all slice/map elements | `for ... := range value` |
| Decode string runes | `for offset, r := range text` |
| Visit a selected index interval | `for i := start; i < end; i++` |
| Repeat while a condition holds | `for condition` |
| Execute first, decide exit inside | `for { ... break }` |

## 3. How It Works in Go

Traverse everything:

```go
for index, value := range values {
    fmt.Println(index, value)
}
```

Traverse only the middle of a slice:

```go
for index := 1; index < len(values)-1; index++ {
    fmt.Println(index, values[index])
}
```

Decode text safely by rune:

```go
for _, r := range text {
    fmt.Printf("%c", r)
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** indexing a string iterates bytes; ranging over a string decodes UTF-8 runes.

Using `range` plus several `continue` and `break` checks to express simple numeric bounds can obscure the intended interval. Conversely, manually indexing an entire slice adds bookkeeping that `range` already expresses.

An infinite loop should usually expose a deliberate `break`, `return`, cancellation, or other exit path. If no form reads clearly, the loop body may be doing too many jobs and could benefit from refactoring.

## 5. Quick Cheat Sheet

- All elements → `range`
- String runes → `range`
- Explicit start/end/update → three-part `for`
- Condition-driven repetition → `for condition`
- Exit chosen inside body → `for {}`
- Slice positions are element indices
- String indices are byte positions
- Avoid manual indexing when `range` states the intent
- Avoid `range` gymnastics for simple numeric bounds
- Infinite loops need a deliberate exit design

### Test Yourself

1. Why is a three-part loop clearer for visiting only the middle of a slice?
2. Why should Unicode string traversal normally use `range`?
3. What design signal does an infinite loop without an obvious exit provide?
