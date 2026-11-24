# Compare Version Numbers

> Striver SDE Sheet 103 · Day 16 String II — split on dots, compare numerically per segment, treat missing segments as zero. Go, stdlib only, tested.

## Problem

Order two dotted version strings: `1.2 < 1.10` (numeric segments, so string comparison is wrong), `1.0 == 1.0.0` (missing trailing segments are zeros), `1.01 == 1.1` (leading zeros don't matter).

## Approach

All three traps dissolve under one framing: a version is a sequence of integers, compared left to right, padded with zeros to equal length. Split on dots, Atoi each segment (which erases leading zeros for free), and run one loop to the *longer* length with missing segments defaulting to 0 — the padding never gets materialized. The antisymmetry test (`cmp(a,b) == -cmp(b,a)`) is a cheap property that catches sign mistakes in any comparator.

**Complexity:** O(n+m) time and space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
