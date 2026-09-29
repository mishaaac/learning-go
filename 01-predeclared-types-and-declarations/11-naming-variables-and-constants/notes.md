# Naming Variables and Constants in Go

Go separates the language rules that make an identifier valid from the conventions that make a name idiomatic and readable. Good names communicate a value's purpose, fit its scope, and use capitalization consistently with Go's export rules.

## 1. What I Need to Understand

- An identifier begins with a Unicode letter or `_`; later characters may also be Unicode digits.
- Reserved keywords such as `var`, `func`, and `type` cannot be used as identifiers.
- Unicode identifiers are legal, but visually similar code points can create different names that are difficult to distinguish.
- `_` alone is the special **blank identifier**; names such as `_0` and `__` are ordinary identifiers.
- Idiomatic multiword names use `mixedCaps` or `MixedCaps`, not `snake_case`; constants follow the same convention.
- Initial capitalization communicates export visibility in the contexts defined by Go, not whether a name represents a constant.
- Short names fit small, obvious scopes; wider scopes usually need more descriptive names, without repeating the value's type unnecessarily.

> ⚠️ **Important correction:** An uppercase first letter does not make every identifier exported. The name must also be declared in the package block or be a field or method name; a local identifier remains local even when it starts with an uppercase letter.

## 2. Key Concepts

| Concept | Language rule or convention | Why it matters |
|---|---|---|
| Identifier syntax | First character: Unicode letter or `_`; remaining characters: Unicode letters, Unicode digits, or `_` | Determines whether the compiler accepts the token as an identifier |
| Reserved keyword | A word such as `const`, `func`, or `range` that cannot be an identifier | Valid identifier characters do not make a keyword reusable as a name |
| Blank identifier | `_` discards a value and creates no binding | It is different from ordinary names that merely contain `_` |
| Exported identifier | A qualifying name that begins with a Unicode uppercase letter | Capitalization controls access from other packages in the defined contexts |
| Idiomatic name | A clear name that follows Go naming conventions | Valid code can still be unnecessarily confusing |
| Scope-sensitive name | A name whose detail matches the distance over which it must be understood | Small scopes need less context than package-level declarations |

Common choices are:

| Situation | Typical naming style |
|---|---|
| Very small local scope | Short, clear name |
| Loop index | `i`, `j` |
| Key and value in `range` | `k`, `v` when their meaning is obvious |
| Multiword unexported name | `indexCounter` |
| Multiword exported name | `IndexCounter` |
| Package-level variable or constant | More descriptive name when its wider scope requires it |
| Constant | Same `mixedCaps` or `MixedCaps` convention; not automatically `UPPER_SNAKE_CASE` |

The name should explain what the value represents. Go's type system already communicates whether that value is an `int`, `string`, or another type in most contexts.

## 3. How It Works in Go

Identifiers may use Unicode letters, digits after the first character, and underscores:

```go
func identifierValues() (int, int, string, string) {
    _0 := 0
    π := 3
    ａ := "fullwidth"
    a := "ASCII"

    return _0, π, ａ, a
}
```

All four names are valid. The fullwidth `ａ` and ASCII `a` are different Unicode code points and therefore different identifiers, despite their similar appearance.

The identifier `_` behaves differently from `_0` or `__`:

```go
func firstValue() int {
    first, _ := twoValues()
    return first
}

func twoValues() (int, int) {
    return 10, 20
}
```

`_` discards the second result and does not introduce a variable. A name such as `_0` would create a normal variable instead.

Multiword identifiers use capitalization rather than separators:

```go
const maxRetries = 3
var packageRequestCount int
```

The names are unexported because they begin with lowercase letters. If a package-level declaration is intentionally part of the package's public API, its name begins with an uppercase letter:

```go
const DefaultRetries = 3
```

The capitalization describes export visibility; it does not distinguish constants from variables.

Short local names work when their meaning stays obvious:

```go
func sum(values []int) int {
    total := 0
    for _, v := range values {
        total += v
    }
    return total
}
```

`v` is readable because its scope is only the small loop body. `total` is slightly more descriptive because it is used throughout the function.

```text
small and obvious scope
        └── short name may be enough

wider or less obvious scope
        └── add the context needed to explain the value's purpose
```

Receiver names are also commonly short and related to the receiver type. Their detailed use belongs to the later topic on methods.

## 4. Examples, Differences, and Common Mistakes

| Name | Valid? | Usually idiomatic here? | Reason |
|---|---:|---:|---|
| `_0` | ✅ | ❌ | Legal, but it communicates little and resembles the blank identifier |
| `π` | ✅ | ❌ | Legal Unicode, but it may be harder to type or search consistently |
| `ａ` | ✅ | ❌ | It can be confused visually with ASCII `a` |
| `index_counter` | ✅ | ❌ | Go normally prefers `indexCounter` |
| `indexCounter` | ✅ | ✅ | Idiomatic unexported multiword name |
| `INDEX_COUNTER` | ✅ | ❌ | Constants do not require `UPPER_SNAKE_CASE` in Go |
| `maxRetries` | ✅ | ✅ | Idiomatic unexported constant or variable name |
| `MaxRetries` | ✅ | ✅ when export is intended | Initial capitalization communicates export visibility in a qualifying declaration |
| `i` | ✅ | ✅ for an obvious loop index | Its small scope supplies the missing context |
| `k`, `v` | ✅ | ✅ for obvious key/value roles | Their conventional meaning is clear in a small `range` loop |

> **Do not confuse:** `_` is the blank identifier; `_0` and `__` are normal identifiers that introduce bindings.

> **Do not confuse:** capitalization does not mark constants. Both a `var` and a `const` can use lowercase or uppercase names; in qualifying declarations, capitalization instead determines export visibility.

Keywords are not identifiers even though they consist of valid letters. For example, `var`, `range`, and `type` cannot be chosen as variable or constant names.

Avoid encoding the type in names such as `userString` or `countInt` when the type adds no useful meaning. Prefer a name that describes the role, such as `user` or `count`; include type-related wording only when it distinguishes the concept itself.

Short names are not automatically better. If several brief names become hard to track, the block may need clearer names or may be doing too much work. Conversely, a very long name inside a two-line loop can repeat context the reader already has.

Unusual Unicode can be appropriate in a specialized domain, but visually confusable characters are risky because two distinct identifiers may look almost identical in review.

## 5. Quick Cheat Sheet

- Identifier start → Unicode letter or `_`.
- Later characters → Unicode letters, Unicode digits, or `_`.
- Reserved keywords cannot be identifiers.
- `_` alone → blank identifier; `_0` and `__` are ordinary names.
- Visually similar Unicode code points can create different identifiers.
- Multiword names → `mixedCaps` or `MixedCaps`, not `snake_case`.
- Constants do not require `UPPER_SNAKE_CASE`.
- Uppercase initial → export meaning only in qualifying package, field, or method declarations.
- Small, obvious scope → short name can be clear.
- Common loop names → `i`, `j`; obvious `range` roles → `k`, `v`.
- Wider scope → add enough description to communicate purpose.
- Usually describe the value's role, not its Go type.

### Test Yourself

1. Why can `ａ` and `a` name two different variables even though they look similar, and what risk does that create?
2. Why do `maxRetries` and `MaxRetries` say something about visibility rather than whether either name belongs to a constant?
3. How should a variable's scope and role influence whether you choose a one-letter name, a short word, or a more descriptive multiword name?
