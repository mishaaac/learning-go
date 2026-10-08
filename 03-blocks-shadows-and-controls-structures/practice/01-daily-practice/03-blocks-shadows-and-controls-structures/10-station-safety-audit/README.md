# Exercise 10 — Station Safety Audit

> **Difficulty:** Challenge
> **Estimated time:** 45 minutes
> **Topics:** Blocks and scope, nested `for-range`, labels, `break`, `continue`, `if`, expressionless `switch`

## Goal

Build a deterministic station audit with station-level rejection, audit-wide emergency shutdown, ordered reading tiers, and correctly scoped summary state.

## Scenario

Process these stations and readings in order: `North: [12, 0, 38]`, `East: [-1, 45, 52]`, `South: [67, 91, 20]`, and `West: [22, 44]`. Negative readings invalidate the current station, while readings of `90` or more stop the entire audit.

## What You Need to Do

- [x] Scan each station's readings in order; ignore zero readings and reject the rest of a station immediately after a negative reading.
- [x] Classify each reached positive reading below `90` as `low` when below `30`, `moderate` when below `60`, or `high` otherwise.
- [x] On an emergency reading, print the station and reading, stop all remaining work, and exclude that reading from the tier totals.
- [x] Report the number of fully completed stations, rejected stations, and readings in each tier after processing stops.

## Rules and Constraints

- Use nested `for-range` loops with labeled `continue` for station rejection and labeled `break` for audit shutdown.
- Use one ordered, mutually exclusive decision for the three reading tiers.
- Keep audit totals in a scope that remains available after both loops; keep per-station state limited to its outer iteration.
- Do not process later readings in a rejected station, later readings after an emergency, or any later station after shutdown.

## Expected Result

`North` completes, `East` is rejected before its later readings, and `South` triggers shutdown after its first non-emergency reading. `West` is untouched, and the final tier and station totals include only work completed before shutdown.

## Definition of Done

- [x] The program compiles and runs.
- [x] Every required behavior above is implemented.
- [x] I manually checked the important cases.
- [x] I can explain my main technical decisions.
- [x] Ready for review.
