# Maps in Go

A map stores associations between unique keys and values. It is the natural choice when data is retrieved by a meaningful key rather than by a sequential integer position.

## 1. What I Need to Understand

- A map type is written `map[K]V`.
- Keys must be comparable; values may have any type.
- The zero value is `nil`: reads are safe, but writes panic.
- An empty literal or `make` creates an initialized map that can be written.
- Maps grow as entries are added and do not guarantee iteration order.

## 2. Key Concepts

| Form | State | Writable? |
|---|---|---:|
| `var counts map[string]int` | `nil`, length `0` | No |
| `map[string]int{}` | Initialized, length `0` | Yes |
| `map[string]int{"go": 1}` | Initialized with data | Yes |
| `make(map[string]int, 100)` | Empty with an initial size hint | Yes |

The hint passed to `make` is neither the map's length nor a maximum size.

## 3. How It Works in Go

```go
teams := map[string][]string{
    "Orcas": {"Fred", "Ralph"},
    "Lions": {"Sarah"},
}

teams["Kittens"] = []string{"Waldo"}
fmt.Println(len(teams)) // 3

var missing map[string]int
fmt.Println(missing["key"]) // 0
```

Map lookup uses a key rather than a sequence position:

```text
key ──lookup──> value
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** reading `nilMap[key]` is valid, but assigning `nilMap[key] = value` causes a runtime `panic`.

Slices and maps are both flexible and have `nil` zero values, but they model different relationships: a slice uses ordered integer positions, while a map uses unique comparable keys and has no defined iteration order.

> ⚠️ **Important correction:** not every Go value can be a map key. Slices, maps, and functions are not comparable and therefore cannot be key types; arrays and structs are valid only when all their components are comparable.

Two maps cannot be compared directly with `==`, except that a map may be compared with `nil`.

## 5. Quick Cheat Sheet

- Type → `map[K]V`
- Key type → must be comparable
- Value type → may be any type
- Zero value → `nil`
- Read from `nil` map → value type's zero value
- Write to `nil` map → `panic`
- Empty writable map → `map[K]V{}` or `make(map[K]V)`
- `make(map[K]V, n)` → initial size hint, not length or limit
- `len(m)` → number of entries
- Iteration order is not specified
- Map-to-map `==` is invalid; map-to-`nil` is valid

### Test Yourself

1. Why can a struct sometimes be a map key while a slice cannot?
2. What is the critical behavioral difference between a `nil` map and an initialized empty map?
3. When should a map be preferred over a slice?
