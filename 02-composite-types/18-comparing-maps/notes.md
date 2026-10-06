# Comparing Maps in Go

Maps cannot be compared with each other using `==`, but the standard `maps` package provides content-based comparison. Choose normal equality or a custom function according to the map's value type and intended semantics.

## 1. What I Need to Understand

- `mapA == mapB` is invalid; a map can only be compared directly with `nil`.
- `maps.Equal` checks for the same keys and values.
- Map iteration order does not affect equality.
- `maps.Equal` requires values that support `==`.
- `maps.EqualFunc` delegates value comparison to a supplied function.

## 2. Key Concepts

| Function | Value comparison | Use case |
|---|---|---|
| `maps.Equal` | `==` | Comparable values with normal equality |
| `maps.EqualFunc` | Supplied function | Custom equality or non-comparable values |
| `m == nil` | Nil comparison | Check whether a map is `nil` |

Both comparison helpers care about entries, not literal or iteration order.

## 3. How It Works in Go

```go
left := map[string]int{
    "hello": 5,
    "world": 10,
}
right := map[string]int{
    "world": 10,
    "hello": 5,
}

fmt.Println(maps.Equal(left, right)) // true
```

Custom comparison:

```go
equal := maps.EqualFunc(left, right, func(a, b int) bool {
    return a == b
})
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** map order is not part of map content, so a different literal or iteration order does not make equal entries unequal.

`maps.Equal` first requires the same number of entries and then matching values for each key. It also considers a `nil` map and an initialized empty map equal because both contain no entries; this does not mean their `nil` status is the same.

> ⚠️ **Important correction:** `maps.Equal` is not available for maps whose value type is non-comparable, such as `map[string][]int`. Use `maps.EqualFunc` and define how those values should be compared.

## 5. Quick Cheat Sheet

- `mapA == mapB` → invalid
- `m == nil` → valid
- Same entries → `maps.Equal(a, b)`
- Order does not matter
- `maps.Equal` compares values with `==`
- Values must therefore be comparable
- Custom value equality → `maps.EqualFunc`
- Keys still follow normal map-key equality
- `nil` map and empty map compare equal by content
- Content equality is different from `nil` status

### Test Yourself

1. Why does entry order not affect `maps.Equal`?
2. Why can `maps.Equal` not compare `map[string][]int` values directly?
3. How can two maps be equal by content while only one is `nil`?
