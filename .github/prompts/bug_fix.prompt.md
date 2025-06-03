# Bug Fix Prompt

## Objective
{{USER_GOAL}}

## Instructions
- Reproduce the issue with existing tests or by running the game.
- Write or update tests that fail because of the bug.
- Apply the minimal fix required to make the tests pass.
- Run `gofmt`, `go vet`, and `go test ./...` from the `v1` directory.
- Document the cause and fix with comments and in the commit message.

