# Infinite `for` Loops in Go

`for { ... }` is Go's explicit infinite-loop form. With no condition to become false, execution repeats until something inside the body exits or external execution is stopped.

## 1. What I Need to Understand

- The syntax is `for { ... }`.
- Omitting the condition is equivalent to a condition that is always true.
- This form has no init, condition, post, or semicolons.
- A practical infinite loop normally contains an exit such as `break` or `return`.
- Without an internal exit, the loop ends only through an external event or program termination.

## 2. Key Concepts

| Form | Header control |
|---|---|
| `for init; condition; post` | Initialization, test, update |
| `for condition` | Test only |
| `for {}` | No header test; repeats indefinitely |

An external timeout or `Ctrl-C` stops the program, not the loop's own logic.

## 3. How It Works in Go

```go
for {
    fmt.Println("Hello")
}
```

A more practical version decides when to leave:

```go
for {
    value := readValue()
    if value == "stop" {
        break
    }
    process(value)
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a platform timeout does not make `for {}` finite; it interrupts the entire program externally.

An infinite loop can be intentional when a server, worker, or event processor waits repeatedly, but the surrounding design still needs a shutdown or cancellation path.

An empty loop may also consume CPU continuously. If the body waits on input or another blocking operation, execution behaves differently even though the loop is still logically infinite.

Use `for condition` when the continuation rule belongs naturally in the header; use `for {}` when exit decisions occur inside the body.

## 5. Quick Cheat Sheet

- Infinite form → `for { ... }`
- No init
- No condition
- No post
- No semicolons
- Missing condition acts as `true`
- Body repeats indefinitely
- `break` exits the loop
- `return` exits the function
- External timeout is not loop termination logic
- Real code should have a deliberate shutdown path

### Test Yourself

1. Why is `for {}` infinite even if an online runner stops it after a few seconds?
2. When is `for {}` clearer than `for condition`?
3. What distinguishes a busy infinite loop from one blocked waiting for input?
