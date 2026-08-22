# Bug Reproduction

## Bug

Purchase acceptance can skip delivery, accept incomplete evidence, race a stale status update, or leave a created device behind when the purchase update fails.

## Trigger

Run the four targeted acceptance tests recorded in `verify_cmds` against the original bug snapshot.

## Observed Error

The tests report that an undelivered purchase can be accepted, required evidence is missing, a stale status is updated, or the device transaction is not rolled back.
