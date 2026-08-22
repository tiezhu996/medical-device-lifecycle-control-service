# Bug Reproduction

## Bug

Maintenance plan generation loses the original worker error and can commit a partial batch with incorrect schedule boundaries.

## Trigger

Run the four targeted maintenance batch tests recorded in `verify_cmds` against the original bug snapshot.

## Observed Error

The failing output reports duplicate or unsupported plan types, incorrect calendar boundaries, a replaced work error, and a device batch that was not rolled back.
