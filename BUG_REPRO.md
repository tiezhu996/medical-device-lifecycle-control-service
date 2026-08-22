# Bug Reproduction

## Bug

Concurrent device refresh and aggregation expose shared mutable snapshots, producing stale results and data races.

## Trigger

Run the four targeted device snapshot tests recorded in `verify_cmds` with the race detector against the original bug snapshot.

## Observed Error

The race detector reports a concurrent write through `runtime.slicecopy`, while the functional checks report mutated copies and unstable ordering.
