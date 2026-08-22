# Bug Reproduction

## Bug

Transfer approval reuses slice backing storage, so evidence normalization and persistence mutate the submitted historical snapshot.

## Trigger

Run the four targeted transfer evidence tests recorded in `verify_cmds` against the original bug snapshot.

## Observed Error

The tests report mutated input evidence, shared backing arrays, repository aliasing, and changed submitted evidence after approval.
