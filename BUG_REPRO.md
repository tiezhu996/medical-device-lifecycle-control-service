# Bug Reproduction

## Bug

The HTTP middleware lifecycle injects request IDs too late, registers panic recovery after downstream execution, fails to stop rejected requests, and places CORS before request correlation.

## Trigger

Run the four targeted request-flow tests recorded in `verify_cmds` against the original bug snapshot.

## Observed Error

The handler sees no request ID, a downstream panic escapes recovery, the business handler runs after a failure response, and an OPTIONS response omits `X-Request-Id`.
