# `if` in Go

Go's `if` selects one branch according to a boolean condition and does not require parentheses around that condition. An optional initial statement can create values scoped to the whole `if`/`else` chain.

## 1. What I Need to Understand

- An `if` condition must produce a `bool`.
- Go writes the condition without surrounding parentheses.
- `else if` and `else` provide mutually exclusive alternative branches.
- Each branch is its own block.
- `if initialStatement; condition` limits declared values to the complete conditional structure.

## 2. Key Concepts

| Form | Scope of a declared variable |
|---|---|
| Declared before `if` | Available before, inside, and after the statement |
| Declared inside one branch | Only that branch |
| Declared in the initial statement | Condition and all `if`/`else if`/`else` branches |

The initial statement is a Go simple statement, commonly a short variable declaration.

## 3. How It Works in Go

```go
if number := rand.Intn(10); number == 0 {
    fmt.Println("too low")
} else if number > 5 {
    fmt.Println("too high", number)
} else {
    fmt.Println("good", number)
}

// number is out of scope here.
```

Execution selects only the first branch whose condition succeeds; `else` handles the remaining case.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a variable declared inside the first branch is not available in `else`; a variable declared before the condition is available to every branch.

The initial statement and condition are separated by a semicolon:

```go
if value, ok := lookup(); ok {
    fmt.Println(value)
}
```

Although other simple statements are legal before the condition, a declaration that supplies the condition is usually the clearest use.

The initial declaration can shadow an outer variable with the same name. Declare the value before the `if` instead when it must remain available afterward.

## 5. Quick Cheat Sheet

- Form → `if condition { ... }`
- Condition must be boolean
- No condition parentheses required
- Alternatives → `else if`, then optional `else`
- Each branch has its own scope
- Initial statement → `if statement; condition`
- Initial variable is visible in the condition
- It is also visible in every branch
- It is out of scope after the full statement
- Initial declaration can shadow an outer name

### Test Yourself

1. Why can a variable from the initial statement be used in `else` but not after the whole `if`?
2. When should a value be declared before the `if` instead of in its initial statement?
3. How can an initial short declaration accidentally cause shadowing?
