# Arrays in Go

An array is a fixed-length sequence of values of one type. Arrays matter because their length is part of their type and because they provide the backing storage used by slices.

## 1. What I Need to Understand

- An array type has the form `[N]T`; both `N` and `T` define the type.
- Every element starts at the element type's zero value unless a literal supplies another value.
- Array indices start at `0` and end at `len(array)-1`.
- Arrays are directly comparable only when their element type is comparable.
- Arrays are best when an exact, meaningful size is known in advance; slices are more flexible for most sequences.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| `[3]int` | Three-element array of `int` | It is a different type from `[4]int` |
| `[...]int{1, 2}` | Length inferred by the compiler | The resulting type is `[2]int` |
| Indexed literal | Values assigned at selected indices | Unspecified elements keep their zero value |
| `[2][3]int` | Array whose elements are `[3]int` arrays | Go builds nested arrays rather than a separate matrix type |
| `len(array)` | Fixed array length | For an array, `len` is determined by its type |

## 3. How It Works in Go

```go
var counts [3]int                 // [0 0 0]
scores := [...]int{10, 20, 30}    // type [3]int
sparse := [6]int{0: 1, 4: 9}      // [1 0 0 0 9 0]

counts[1] = 7
fmt.Println(counts[1], len(counts)) // 7 3
fmt.Println(scores == [3]int{10, 20, 30}) // true
```

A nested array is an array of arrays:

```go
var grid [2][3]int
grid[1][2] = 5
```

```text
[2][3]int
├── [3]int
└── [3]int
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `[...]int{1, 2, 3}` is an array, while `[]int{1, 2, 3}` is a slice.

`[3]int` cannot be assigned to `[4]int`, and a normal variable cannot determine an array length at runtime. A constant out-of-range index is rejected during compilation; a dynamically computed invalid index causes a runtime `panic`.

> ⚠️ **Important correction:** arrays are not always comparable. `[3]int` supports `==`, but an array such as `[2][]int` does not because slices are not comparable.

Arrays are useful when size is part of the data's meaning, such as a fixed-size digest. Their most common indirect role is serving as backing storage for slices.

## 5. Quick Cheat Sheet

- `[N]T` → array of `N` values of type `T`
- Length is part of the array type
- `var a [3]int` → `[0 0 0]`
- `[...]T{...}` → compiler infers the length
- Indexed literals leave omitted positions at the zero value
- Valid indices → `0` through `len(a)-1`
- `[2][3]int` → array of two `[3]int` arrays
- `==` and `!=` require comparable element types
- `[3]int` and `[4]int` are different types
- Arrays commonly provide backing storage for slices

### Test Yourself

1. Why can a function accepting `[3]int` not also accept `[4]int`?
2. What determines whether two arrays can be compared with `==`?
3. What is the difference between an invalid constant index and an invalid runtime index?
