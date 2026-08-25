# clock

[![ci](https://github.com/platformfix/clock/actions/workflows/ci.yml/badge.svg)](https://github.com/platformfix/clock/actions/workflows/ci.yml)
[![e2e](https://github.com/platformfix/clock/actions/workflows/e2e.yml/badge.svg)](https://github.com/platformfix/clock/actions/workflows/e2e.yml)
[![lint](https://github.com/platformfix/clock/actions/workflows/lint.yml/badge.svg)](https://github.com/platformfix/clock/actions/workflows/lint.yml)
[![commit-lint](https://github.com/platformfix/clock/actions/workflows/commit-lint.yaml/badge.svg)](https://github.com/platformfix/clock/actions/workflows/commit-lint.yaml)
[![pr-lint](https://github.com/platformfix/clock/actions/workflows/pr-lint.yml/badge.svg)](https://github.com/platformfix/clock/actions/workflows/pr-lint.yml)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/platformfix/clock/badge)](https://scorecard.dev/viewer/?uri=github.com/platformfix/clock)
[![Latest Release](https://img.shields.io/github/v/release/platformfix/clock)](https://github.com/platformfix/clock/releases)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Prints the current time, once a second, until you stop it.

## Quickstart

```bash
docker run ghcr.io/platformfix/clock:latest
```

Or build and run it directly:

```bash
go run ./cmd/clock
```

## How it works

`internal/clock.Run` writes an RFC3339 timestamp (UTC) to the given writer
immediately, then once a second, until the given context is cancelled.
`cmd/clock/main.go` supplies stdout as that writer and cancels the context on
SIGINT or SIGTERM. That's the whole thing: no flags, no configuration.

## Local development

```bash
go test ./...
golangci-lint run ./...
goreleaser build --single-target --snapshot --clean -o clock
docker build -t clock:dev .
```

## Releases

Tagged releases (`vX.Y.Z`) are built and published by
[goreleaser](.goreleaser.yaml): a multi-arch image pushed to
`ghcr.io/platformfix/clock`, cosign-signed (keyless, via GitHub's OIDC
identity), with an SBOM and SLSA build provenance attached.

## License

[MIT](LICENSE)
