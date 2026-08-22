# Bug Reproduction

## Bug

Cancellation and deadlines are cached, ignored, detached, and converted into an unidentifiable error across authentication, rate limiting, JWT parsing, and middleware errors.

## Trigger

Run the four targeted cancellation tests recorded in `verify_cmds` against the bug branch.

## Observed Error

The tests report reuse of a canceled context, a rate-limit wait that ignores cancellation, token parsing after cancellation, and loss of `context.Canceled` identity.
