# Bug Reproduction

## Bug

Authentication reuses an earlier request context and breaks cancellation propagation across registration, login, repository lookup, and audit work.

## Trigger

Run the four targeted authentication context tests recorded in `verify_cmds` against the bug branch.

## Observed Error

The tests report that caller context values disappear, canceled audit work continues, repository cancellation is lost, and a later login inherits an expired deadline.
