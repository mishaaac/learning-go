# Active Recall

## Question 1 — Arrays and Slice Views

Describe how array and slice types represent sequences, including how length affects their types, how zero values behave, which values can be compared, and what happens to storage when an array is sliced.

- [x] Answered without notes.

## Question 2 — Slice State and Growth

Describe the roles of slice length, capacity, and backing storage. How do nil slices, empty non-nil slices, slice literals, the different forms of `make`, `append`, and `copy` affect those three properties?

- [x] Answered without notes.

## Question 3 — Unicode Text Representation

How does Go represent a string, and what do string length, indexing, slicing, conversion to `[]byte`, conversion to `[]rune`, and ranging over a string each observe when the text contains multi-byte UTF-8 code points?

- [x] Answered without notes.

## Question 4 — Map States and Lookups

Describe the requirements for map keys, the behavior of nil and initialized maps, the result of reading a missing key, the purpose of comma-ok, the meaning of a map size hint, and the guarantees Go makes about iteration order.

- [x] Answered without notes.

## Question 5 — Why Sub-Slice Appends Can Change Other Values

Explain why appending to a sub-slice can overwrite an element visible through the original slice, while a later append may stop affecting it. Why can a full slice expression change the append behavior without preventing mutations to elements already in the shared range?

- [x] Answered without notes.

## Question 6 — Why Struct Operations Depend on Their Fields and Names

Explain why adding a slice field can make an otherwise comparable struct invalid for `==`. Also explain why two separately named struct types with matching fields can require explicit conversion even though a compatible anonymous struct may be directly assignable.

- [x] Answered without notes.

## Question 7 — Comparing Reset Operations

Compare `clear(s)`, `s = s[:0]`, and `s = nil` for a slice, then compare `clear(m)` with `m = nil` for a map. For each operation, discuss length, capacity where applicable, nil status, retained values or entries, and whether the result can be written to immediately.

- [x] Answered without notes.

## Question 8 — Design a Composite Data Audit

A data audit receives ordered records with an identifier, a category from a fixed allowed list, and a Unicode label. It must reject duplicate identifiers and unknown categories, preserve accepted order, distinguish missing counts from stored zero counts, report byte and rune totals, and keep an independent snapshot before resetting its working data. Describe the arrays, slices, maps, set representation, structs, conversions, comparisons, and copy or clear operations you would choose, including which values would share storage and which would not.

- [x] Answered without notes.

## Completion

- [x] All 8 questions answered.
- [x] I answered without opening my notes.
- [x] I marked questions I was unsure about.
- [x] Ready for review.
