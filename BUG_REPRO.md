# Bug Reproduction

## Bug

Statistics responses share slice backing storage across DTO cloning, service refresh, handler preparation, and routed payloads.

## Trigger

Run the four targeted statistics snapshot tests recorded in `verify_cmds` against the bug branch.

## Observed Error

The tests report changed source groups, a drifting previous overview, handler append reuse, and a routed payload that changes after source reuse.
