# `var` vs. `:=` in Go

Go provides `var` declarations and short variable declarations with `:=`. Choosing between them communicates whether scope, an explicit type, a zero value, or concise local type inference is most important.

## 1. What I Need to Understand

- `var` works at package level and inside functions; `:=` works only inside functions.
- `var` can state a type explicitly, infer it from an initializer, or initialize a declared type with its **zero value**.
- `:=` always requires initializer expressions and infers the variables' types from them.
- `:=` is a declaration, not a shorter spelling of assignment with `=`.
- A short declaration may reuse same-block variables only when it also declares at least one new non-blank variable and preserves the reused variables' types.
- Scope matters: using `:=` in an inner block can create a new variable that shadows a name from an outer block.
- Multiple variables should normally be declared together only when their values are naturally related, such as values returned by the same operation.

> ⚠️ **Important correction:** “At least one new variable” is not the complete redeclaration rule for `:=`. Reused variables must have been declared earlier in the same block—or in the function's parameter list when the short declaration is in its body—must keep the same type, and `_` does not count as a new variable.

## 2. Key Concepts

| Feature | `var` | `:=` |
|---|---|---|
| Allowed scope | Package level or inside functions | Inside functions only |
| Initializer | Optional when a type is present | Required |
| Explicit type | Allowed | Not allowed in the declaration syntax |
| Type inference | Used when the type is omitted | Always used |
| Zero-value declaration | `var count int` | Not available |
| Multiple variables | Allowed | Allowed |
| Grouped declaration | `var (...)` | Not available |
| Reuse of existing names | A `var` declaration cannot redeclare a name in the same block | Possible under the short redeclaration rules |

The main declaration forms are:

| Form | Meaning |
|---|---|
| `var count int = 10` | Explicit type and initial value |
| `var count = 10` | Type inferred from the initializer |
| `var count int` | Explicit type initialized to its zero value |
| `count := 10` | Local short declaration with an inferred type |
| `var (...)` | A group of separate `var` declarations |

`var count = 10` and `count := 10` produce the same type and value when they appear in a function and `count` is new. Their syntax and the contexts where they are permitted are still different.

## 3. How It Works in Go

The three common `var` forms express different intentions:

```go
func declareValues() (int, int, bool) {
    var count int = 10
    var limit = 20
    var ready bool

    return count, limit, ready
}
```

`count` has an explicit type and value, `limit` is inferred as `int`, and `ready` receives the zero value `false`.

A declaration can introduce several related variables:

```go
func coordinates() (int, string) {
    var x, label = 10, "start"
    return x, label
}
```

Their types are inferred separately: `x` is `int` and `label` is `string`.

A grouped `var` declaration combines separate declaration specifications:

```go
var (
    maxRetries int = 3
    serviceName     = "learner"
    enabled    bool
)
```

This syntax is valid at package level and inside functions. At package level, `var` is required because `:=` is not permitted there.

Inside a function, `:=` is the usual concise form when the inferred type is the intended type:

```go
func message() string {
    text := "hello"
    return text
}
```

A short declaration can mix one existing same-block variable with a new variable:

```go
func values() (int, string) {
    count := 10
    count, label := 30, "thirty"
    return count, label
}
```

The second short declaration assigns `30` to the existing `count` and declares `label`. It is valid because `label` is new and `count` remains an `int`.

Declaration and assignment are separate operations:

```go
func updateCount() int {
    count := 10
    count = 20
    return count
}
```

The first statement declares `count`; the second assigns a new value to the variable that already exists.

## 4. Examples, Differences, and Common Mistakes

| Situation | Valid? | Reason |
|---|---:|---|
| Package level: `var count = 10` | ✅ | `var` is allowed at package level |
| Package level: `count := 10` | ❌ | Short declarations are restricted to functions |
| Local: `var count int` | ✅ | `count` starts with the zero value `0` |
| Local: `count := 10` | ✅ | The initializer gives `count` the inferred type `int` |
| After `count := 10`: `count := 20` | ❌ | The short declaration introduces no new variable |
| After `count := 10`: `count, label := 20, "twenty"` | ✅ | `label` is a new non-blank variable |
| After `count := 10`: `count, _ := 20, "ignored"` | ❌ | `_` creates no binding and does not count as new |

> **Do not confuse:** `:=` declares at least one variable; `=` only assigns to variables that are already declared.

An outer variable is not redeclared by `:=` in an inner block. A new variable is created and shadows the outer one:

```go
func shadowExample() int {
    count := 10

    if count > 0 {
        count, label := 20, "inner"
        _, _ = count, label
    }

    return count
}
```

The function returns `10`. The inner `count` is a different variable whose scope ends with the `if` block.

> **Do not confuse:** same spelling does not imply the same variable. The block in which an identifier is declared determines its binding and scope.

Both of these local declarations are valid:

```go
var data byte = 20
otherData := byte(20)
```

The first form emphasizes the desired type directly; the second uses an explicit conversion before inference. Preferring `var data byte = 20` when the non-default type is part of the intent is a style choice, not a language requirement.

Grouped declarations should communicate a relationship rather than merely save lines. They are especially natural when receiving multiple results or using the comma-ok idiom; those mechanisms are studied separately.

Package variables are mutable unless another rule prevents mutation. Keeping them rare or effectively unchanged is design guidance that makes data flow easier to follow; `var` itself does not enforce immutability.

## 5. Quick Cheat Sheet

- `var` → allowed at package level and inside functions.
- `:=` → allowed only inside functions.
- `var count int = 10` → explicit type and value.
- `var count = 10` → inferred type.
- `var count int` → zero value of `int`.
- `count := 10` → concise local declaration with inference.
- `var (...)` → groups separate declarations.
- `:=` may reuse same-block variables only if at least one non-blank variable is new and reused types stay unchanged.
- An outer name used with inner `:=` can be shadowed by a new variable.
- `_` never introduces a binding and does not satisfy the “new variable” rule.
- `:=` declares; `=` assigns.
- Declare multiple variables together when they are related, and keep mutable package variables uncommon.

### Test Yourself

1. When does `var` communicate an intention that `:=` cannot express directly?
2. Why does the outer `count` remain unchanged when an inner block executes `count, label := 20, "inner"`?
3. After `count := 10`, why is `count, label := 20, "twenty"` valid while `count, _ := 20, "ignored"` is not?
