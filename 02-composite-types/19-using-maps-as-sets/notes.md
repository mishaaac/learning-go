# Using Maps as Sets in Go

Go has no built-in `set` type, but a map can model unique membership by storing set elements as keys. The value is commonly either `bool` for convenient lookup or `struct{}` when only presence matters.

## 1. What I Need to Understand

- Set elements become map keys, so the element type must be comparable.
- Map keys are unique, so repeated insertion does not create duplicates.
- `map[T]bool` uses `true` for membership and the zero value `false` for absence.
- `map[T]struct{}` stores no meaningful value and checks membership with `comma ok`.
- `len(set)` reports the number of unique keys.

## 2. Key Concepts

| Representation | Insert | Membership test |
|---|---|---|
| `map[T]bool` | `set[value] = true` | `if set[value]` |
| `map[T]struct{}` | `set[value] = struct{}{}` | `_, ok := set[value]` |

Both forms support insertion, membership, deletion with `delete`, and counting with `len`.

## 3. How It Works in Go

```go
values := []int{5, 10, 2, 5, 2}
set := map[int]bool{}

for _, value := range values {
    set[value] = true
}

fmt.Println(len(values)) // 5
fmt.Println(len(set))    // 3
fmt.Println(set[5])      // true
fmt.Println(set[99])     // false
```

Using an empty struct:

```go
compact := map[int]struct{}{5: {}}
_, present := compact[5]
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** in `map[T]bool`, a present key stored with `false` looks absent in `if set[value]`. A bool-backed set should consistently store `true`.

`map[T]struct{}` makes the “keys only” intent explicit but requires `comma ok` for membership. `map[T]bool` is often simpler to read and can represent enabled/disabled states if that is actually desired.

> ⚠️ **Important correction:** `struct{}` has size zero and `bool` has a nonzero size, but total per-entry map memory is an implementation detail. Do not assume a guaranteed one-byte saving for every entry; choose based on clarity unless measurement shows memory matters.

Union, intersection, and difference must be implemented from these basic operations or provided by another abstraction.

## 5. Quick Cheat Sheet

- Go has no built-in `set` type
- Set element → map key
- Element type must be comparable
- Duplicate insertion keeps one key
- `len(set)` → unique element count
- Simple form → `map[T]bool`
- Insert bool form → `set[x] = true`
- Test bool form → `if set[x]`
- Keys-only form → `map[T]struct{}`
- Test struct form → `_, ok := set[x]`
- Remove member → `delete(set, x)`

### Test Yourself

1. Why do duplicate source values produce only one set member?
2. What invariant must a `map[T]bool` set maintain for `if set[x]` to mean membership?
3. What tradeoff distinguishes `map[T]bool` from `map[T]struct{}`?
