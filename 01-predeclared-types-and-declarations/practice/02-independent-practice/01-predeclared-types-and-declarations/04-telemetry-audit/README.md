# Problem 4 — Telemetry Audit

> **Difficulty:** Integrative
> **Estimated time:** 45 minutes
> **Topics:** Constants, numeric types, conversions, bitwise operations, booleans, complex values, strings, runes, zero values

## Goal

Design a complete telemetry report that uses the studied declaration and predeclared-type rules consistently.

## Scenario

Session `Vega` received 13 of 20 frames and used 2 of 4 permitted retries. Its three flag bits represent connected, encrypted, and cached states; the current mask enables connected and cached only. A signal sample has real component `3` and imaginary component `-4.5`.

## What You Need to Do

- [x] Report the session name, a rune status marker, completion percentage, retries remaining, and whether completion has reached 75 percent.
- [x] Represent the three flags with named binary masks and report the boolean state of each flag from the current combined value.
- [x] Report the complex signal sample and its real and imaginary components.
- [x] Include an optional text field declared with its zero value and make that empty value visible.
- [x] Provide a short written justification for the types, declaration forms, and constant forms chosen for the report.

## Rules and Constraints

The program must include a grouped constant declaration, at least one typed constant, at least one untyped constant, a fixed-width integer variable, a `byte`, a `rune`, a zero-value declaration, and both `var` and `:=`. Calculated fields must be derived from the scenario values. Use explicit conversions where concrete numeric types differ or where integer arithmetic would lose required precision.

## Expected Result

The report shows 65 percent completion, 2 retries remaining, and a false threshold result. Connected and cached are true while encrypted is false. The signal components and visibly empty optional field match the scenario.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior is implemented.
- [x] I solved it without help.
- [x] I can explain my decisions.
- [x] Ready for review.
