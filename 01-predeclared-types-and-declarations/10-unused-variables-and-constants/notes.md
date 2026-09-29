# Unused Variables and Constants in Go

Go rejects variables declared inside a function body when they are never used, which keeps accidental leftovers out of compiling code. This check does not prove that every assigned value matters, and unused package variables, parameters, and constants follow different rules.

## 1. What I Need to Understand

- A variable declared inside a function body must be used; otherwise, compilation fails with a `declared and not used` error.
- Merely assigning another value to that variable does not count as using its value.
- Once the variable is used somewhere, the compiler does not perform complete data-flow analysis to ensure that every assigned value is later read.
- Function parameters and receivers may remain unused even though variables declared in the function body may not.
- Package-level variables may remain unused, and both local and package-level constants may remain unused.
- The blank identifier `_` can deliberately discard a value, but it should not be used to hide code that ought to be removed.
- For the example studied here, `go vet` does not report overwritten or final unread assignments; specialized analyzers can perform deeper checks.

> ⚠️ **Important correction:** The strict unused-variable rule applies to variables declared inside a function body. Unused parameters and receivers are allowed, so “every local variable must be used” is too broad without this distinction.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| Unused local variable | A variable declared in a function body but never used | It causes a compile-time error |
| Variable use | An operation that reads or otherwise uses the declared variable's value | It satisfies the compiler's unused-variable check |
| Unread assignment | A value assigned to a used variable but overwritten or abandoned before being read | It may compile even though the assignment is unnecessary |
| Package variable | A variable declared outside function bodies | The compiler permits it to remain unused |
| Unused constant | A declared constant that is never referenced | It is permitted because constant evaluation has no runtime side effects |
| Blank identifier | `_`, a placeholder that discards a value without creating a binding | It is useful when a produced value is intentionally unwanted |

The key comparison is:

| Situation | Compiler result |
|---|---|
| Variable declared in a function body and never used | ❌ Error |
| Function parameter or receiver never used | ✅ Allowed |
| Package-level variable never used | ✅ Allowed |
| Local or package-level constant never used | ✅ Allowed |
| Particular assignment overwritten before a later read | ✅ May be allowed |
| Particular final assignment never read | ✅ May be allowed |

The compiler's rule is mainly about whether the variable binding is used, not whether every value that flows through it contributes to the program's result.

## 3. How It Works in Go

This declaration would fail if uncommented because `value` is never used:

```go
func unusedLocal() {
    // value := 10 // compile-time error: declared and not used
}
```

Assigning to the variable again would not fix the problem if its value were still never read:

```go
func stillUnused() {
    // value := 10
    // value = 20 // assignment alone is not a meaningful use
}
```

The example from the source compiles because `fmt.Println(x)` uses `x`:

```go
import "fmt"

func demonstrateAssignments() {
    x := 10
    x = 20
    fmt.Println(x)
    x = 30
}
```

The individual values have different outcomes:

```text
x := 10       → overwritten before it is read
x = 20        → read by fmt.Println
x = 30        → never read afterward
```

The compiler sees a valid use of the variable `x`; it does not reject the first and last assignments as dead stores.

Parameters, package variables, and constants may remain unused:

```go
var packageValue = 10

const packageLimit = 20

func process(unusedParameter int) {
    const localLimit = 30
}
```

All three unused declarations are permitted. This does not mean that keeping unnecessary names is good design; it only describes what the compiler accepts.

The blank identifier can discard a value explicitly:

```go
func discardValue() {
    value := 10
    _ = value
}
```

`_ = value` evaluates `value` and prevents the unused-variable error. When a declaration is genuinely unnecessary, removing it is usually clearer than adding a blank assignment only to silence the compiler.

```text
variable declared in function body
               │
               ├── never used → compile-time error
               │
               └── used at least once → compiles
                                      │
                                      └── individual unread assignments may remain
```

## 4. Examples, Differences, and Common Mistakes

| Example | Valid? | Reason |
|---|---:|---|
| `func run() { value := 10 }` | ❌ | A variable declared in the function body is never used |
| `func run(value int) {}` | ✅ | An unused function parameter is allowed |
| `var packageValue = 10` | ✅ | An unused package-level variable is allowed |
| `func run() { const limit = 10 }` | ✅ | An unused local constant is allowed |
| `value := 10; _ = value` inside a function | ✅ | The blank assignment explicitly discards the value |
| Declare `x`, print it once, then assign one final unread value | ✅ | The binding is used even though the final assignment is not |

> **Do not confuse:** an unused variable is not the same as an unread assignment. The first is rejected in a function body; the second may survive after the variable has otherwise been used.

For the studied assignment sequence, both the compiler and standard `go vet` accept the code. This should not be generalized to mean that `go vet` has no assignment checks; it simply does not provide general dead-store detection for this case. More specialized static-analysis tools can report values assigned but never read.

> ⚠️ **Important correction:** Go does not guarantee the exact contents of the generated binary. An unused constant has no runtime evaluation or side effects and may be discarded, but saying that it will definitely be absent from every compiled binary goes beyond the language rules.

Unused package variables are legal, but they can still make the program harder to understand by suggesting state that has no purpose. The compiler's permission is not a recommendation to keep them.

> **Do not confuse:** `_ = value` intentionally suppresses the unused-variable error; it does not prove that the original declaration contributes useful behavior.

## 5. Quick Cheat Sheet

- Variable declared in a function body and never used → compile-time error.
- A later assignment alone does not count as reading the variable's value.
- One valid use of a variable does not make every assignment useful.
- Overwritten or final unread assignments may still compile.
- Unused parameters and receivers are allowed.
- Unused package-level variables are allowed.
- Unused local and package-level constants are allowed.
- `_` discards a value without introducing a binding.
- Prefer removing dead code over adding `_ = value` only to silence the compiler.
- The studied dead assignments are not reported by standard `go vet`.
- Specialized static analyzers can detect more unread assignments.
- ⚠️ Binary contents are an implementation detail; only the lack of runtime constant evaluation and side effects is dependable here.

### Test Yourself

1. Why does `x := 10; x = 20` still fail when `x` is never read, while the longer example containing `fmt.Println(x)` compiles?
2. In the sequence `x := 10; x = 20; fmt.Println(x); x = 30`, which assigned values are read, and what does the compiler actually verify?
3. Why may an unused parameter, package-level variable, or constant compile even though an unused variable declared inside a function body does not?
