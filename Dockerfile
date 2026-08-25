# syntax=docker/dockerfile:1
FROM gcr.io/distroless/static-debian12:nonroot
COPY clock /clock
USER nonroot:nonroot
ENTRYPOINT ["/clock"]
