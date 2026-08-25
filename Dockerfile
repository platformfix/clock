# syntax=docker/dockerfile:1
#
# Base image tag is deliberately floating, not digest-pinned: it's the one
# dependency in this repo that isn't SHA-pinned (every GitHub Action is),
# and that's a conscious choice, not an oversight. Digest-pinning trades
# reproducibility for a manual bump burden with no automation in place to
# carry it. The floating tag keeps picking up distroless's own security
# patches automatically.
FROM gcr.io/distroless/static-debian12:nonroot
COPY clock /clock
USER nonroot:nonroot
ENTRYPOINT ["/clock"]
