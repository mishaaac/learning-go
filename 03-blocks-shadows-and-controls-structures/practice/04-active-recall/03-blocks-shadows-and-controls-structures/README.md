# Active Recall

## Question 1 — Visibility Across Blocks

Describe the universe, package, file, function, and nested control blocks in Go. For each one, state what kind of identifiers it commonly contains and where those identifiers can be used.

- [x] Answered without notes.

## Question 2 — Short Declarations and Shadowing

When a short declaration contains a name that is already visible, what determines whether that name is reused or a new shadowing identifier is declared? Explain what happens to the enclosing identifier while the inner one is in scope and after the inner block ends.

- [x] Answered without notes.

## Question 3 — Conditional Scope and Selection

What is the scope of a value declared in the initial statement of an `if` or `switch`, and how does that differ from a value declared inside one branch or case? Also state how Go decides which branch or case executes.

- [x] Answered without notes.

## Question 4 — Range Contracts

For arrays or slices, maps, and strings, identify the values produced by `for-range` and the important guarantee or limitation attached to each kind of operand.

- [x] Answered without notes.

## Question 5 — Why Break May Not Leave the Loop

Explain why an unlabeled `break` inside a `switch` nested in a `for` does not terminate the loop, and why a label changes the result.

- [x] Answered without notes.

## Question 6 — Why Goto Is Restricted

Explain why Go rejects a `goto` that enters an inner block or skips a declaration whose variable would be in scope at the destination.

- [x] Answered without notes.

## Question 7 — Comparing Loop Forms

Compare a three-part `for`, a condition-only `for`, an infinite `for`, and `for-range`. For each form, describe the kind of iteration it communicates best and one situation in which choosing another form would be clearer.

- [x] Answered without notes.

## Question 8 — Nested Audit Design

A program examines a slice of strings in order and scans each string as runes. A `!` rejects only the current string, an `X` stops the entire audit, accepted strings are reported only after a complete scan, and final counters must remain available after all loops finish. Explain how block scope, range behavior, branching, and labeled control interact in a correct design, including what must remain unprocessed after each marker.

- [x] Answered without notes.

## Completion

- [x] All 8 questions answered.
- [x] I answered without opening my notes.
- [x] I marked questions I was unsure about.
- [x] Ready for review.
