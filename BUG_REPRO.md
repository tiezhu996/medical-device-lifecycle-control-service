# Bug Reproduction

## Bug

Calibration result validation and error propagation lose record identity, allowing invalid repeated state changes and mapping missing records incorrectly.

## Trigger

Run the four targeted calibration tests recorded in `verify_cmds` against the original bug snapshot.

## Observed Error

The tests fail on incomplete unqualified results, repeated failure state, missing-record error identity, and preservation of the prior calibration status.
