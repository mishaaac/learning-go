# Booleans in Go

The predeclared type `bool` represents truth values in Go. Its two possible values and defined zero value make boolean variables predictable even when they have no explicit initializer.

## 1. What I Need to Understand

- A variable of type `bool` can hold only `true` or `false`.
- `true` and `false` are Go's predeclared boolean constants.
- The **zero value** of `bool` is `false`.
- `var flag bool` creates a `bool` whose initial value is automatically `false`.
- In `var isAwesome = true`, Go infers the variable's type as `bool` from its initializer.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| `bool` | Go's predeclared boolean type | It represents a condition with two possible states |
| `true` | The true boolean constant | It represents a condition that holds |
| `false` | The false boolean constant | It represents a condition that does not hold |
| Zero value | The value used when a `bool` has no explicit initializer | A declared `bool` always starts with a defined value |
| Type inference | Go determines a variable's type from its initializer | `var isAwesome = true` produces a variable of type `bool` |

The two values are values of the same type; they do not create separate types.

## 3. How It Works in Go

An explicit type without an initializer uses the zero value:

```go
var flag bool
```

`flag` has type `bool` and starts as `false`. Writing `= false` would produce the same initial value:

```go
var firstFlag bool
var secondFlag bool = false
```

Both variables are `false`. The shorter first form is enough when the zero value expresses the intended initial state.

When an initializer is present but the type is omitted, Go determines the type from that value:

```go
var isAwesome = true
var isReady = false
```

Both variables have type `bool`. The constants `true` and `false` have the default type `bool` when a concrete type is required and the context does not provide another boolean type.

```text
var flag bool
     │    │
     │    └── explicit type
     └─────── receives false as its zero value

var isAwesome = true
     │           │
     │           └── initializer
     └────────────── type inferred as bool
```

## 4. Examples, Differences, and Common Mistakes

| Declaration | Type | Initial value | How it is determined |
|---|---|---|---|
| `var flag bool` | `bool` | `false` | The type is explicit; the value is the zero value |
| `var isAwesome = true` | `bool` | `true` | The value is explicit; the type is inferred |
| `var isReady bool = true` | `bool` | `true` | Both the type and value are explicit |

> **Do not confuse:** `false` is a boolean value; `"false"` is a `string` containing text.

```go
var flag = false
var label = "false"
```

`flag` has type `bool`, while `label` has type `string`.

A declaration without an initializer is not an uninitialized variable:

```go
var enabled bool
```

`enabled` is already valid and has the value `false`. Adding `= false` is optional and should be done only when it makes the intent clearer.

The distinction between `var` and the short declaration operator `:=` belongs to the later topic on variable declarations; it is not needed to understand the zero value of `bool`.

## 5. Quick Cheat Sheet

- `bool` → Go's predeclared boolean type.
- Possible values → `true` and `false`.
- Zero value of `bool` → `false`.
- `var flag bool` → type `bool`, initial value `false`.
- `var isAwesome = true` → type inferred as `bool`, value `true`.
- `= false` is unnecessary when the zero value is the intended initial value.
- ⚠️ `false` is a `bool`; `"false"` is a `string`.
- The details of `var` versus `:=` are a separate declaration topic.

### Test Yourself

1. What value does `var flag bool` give to `flag`, and where does that value come from?
2. How does Go determine the type and value of `var isAwesome = true`?
3. Why are `false` and `"false"` not interchangeable even though they look similar?
