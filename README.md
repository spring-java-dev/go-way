# go-way

Idiomatic Go solutions to simple programming tasks, exploring concurrency
patterns such as goroutines, channels, fan-in/fan-out, worker pools, and pipelines.

## Overview

This repository contains solutions to small programming problems
(e.g. Codewars katas) implemented in Go.

Many of these tasks can be solved in a simpler, purely sequential way.
Here, concurrency is used **intentionally** to explore how Go’s concurrency
primitives behave in practice and how common patterns are composed.

The focus is not on code golf or minimal solutions, but on:
- expressing concurrency clearly
- understanding trade-offs
- practicing Go’s core abstractions

## Motivation

Go encourages a particular way of thinking about concurrency:
*share memory by communicating*.

This repository uses small, well-known problems as a safe sandbox to:
- practice goroutines and channels
- implement fan-in / fan-out patterns
- build worker pools
- compose pipelines
- reason about coordination and cancellation

Using simple tasks helps isolate concurrency concerns without domain noise.

## What this repository is (and is not)

**This repository is:**
- a learning and exploration project
- a collection of idiomatic Go examples
- a showcase of concurrency patterns on small problems

**This repository is NOT:**
- a claim that concurrency is always the best solution
- an attempt to over-optimize trivial problems
- a benchmark-focused project

## Concurrency Patterns Used

- goroutines
- channels (buffered and unbuffered)
- fan-in / fan-out
- worker pools
- pipelines
- basic synchronization and coordination

## Why Codewars katas?

Katas provide:
- small, well-defined problem statements
- familiar constraints
- easy comparison between approaches

They are convenient test cases for experimenting with design decisions
rather than problem complexity.

## Disclaimer

Concurrency is a tool, not a goal.

Examples in this repository prioritize clarity of concurrency patterns
over brevity and may not represent the simplest possible solution.

## License

MIT
