# Blocks in Go

Blocks define where declared identifiers are visible and usable. Understanding nested scopes is essential because Go control structures introduce blocks and inner declarations can hide outer names.

## 1. What I Need to Understand

- The universe block contains Go's predeclared identifiers.
- Package-level declarations belong to the package block; imported package names belong to a file block.
- A function body is a block, and its parameters are visible throughout that body.
- Explicit braces and control structures introduce nested blocks.
- An inner block can use names from enclosing blocks unless it declares the same name.

## 2. Key Concepts

| Block | Typical contents | Reach |
|---|---|---|
| Universe | `int`, `true`, `nil`, `make` | All Go source |
| Package | Package-level variables, constants, types, functions | Package |
| File | Names introduced by `import` | One source file |
| Function body | Parameters and local declarations | Function body |
| Nested/control block | Declarations inside braces or clauses | That inner scope |

## 3. How It Works in Go

```go
var packageCount int

func printValue(input int) {
    value := input

    if value > 0 {
        message := "positive"
        fmt.Println(value, message)
    }

    // message is out of scope here.
}
```

```text
universe
└── package
    └── file imports
        └── function
            └── control block
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** visibility flows from an enclosing block into an inner block, not from an inner block back out.

The braces of a function body form an explicit block. Go also defines implicit blocks for constructs such as each `if`, `for`, and `switch`, and for clauses within `switch` statements.

> ⚠️ **Important correction:** imports do not belong to the package block. An imported name is scoped to the file containing that import, even though package-level declarations are visible across files in the same package.

Declaring the same name in an inner block creates a different identifier and shadows the outer one for that scope.

## 5. Quick Cheat Sheet

- Blocks control identifier scope
- Universe block → predeclared identifiers
- Package block → package-level declarations
- File block → imported names
- Function parameters are visible in the function body
- `{}` creates an explicit block
- Control structures also define implicit blocks
- Inner code can see enclosing declarations
- Outer code cannot see inner declarations
- Same inner name → shadowing

### Test Yourself

1. Why can two files in one package use the same package-level declaration but not automatically the same import name?
2. In which direction does visibility work between nested blocks?
3. What happens when an inner block declares a name already used by an outer block?
