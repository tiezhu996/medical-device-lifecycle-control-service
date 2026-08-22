# Bug Reproduction

## Bug

The audit batch lifecycle closes completion too early and can block forever while delivering an error, while cloned and middleware payloads lose fields.

## Trigger

Run the four targeted audit tests recorded in `verify_cmds` with the race detector against the bug branch.

## Observed Error

The failures report an incomplete cloned audit record, a blocked error sender, completion while writers remain blocked, and an incomplete middleware payload.
