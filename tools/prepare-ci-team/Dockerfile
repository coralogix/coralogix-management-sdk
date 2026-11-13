# Build with tools/prepare-ci-team as the Docker context.
FROM golang:1.26.10-bookworm@sha256:5de7c99eb78ca5e515a4c340235d0955ec5964a5071c492c15503d7c267da473 AS build
WORKDIR /src
RUN mkdir -p -m 1777 /out/tmp
COPY *.go ./
RUN CGO_ENABLED=0 GO111MODULE=off go build -trimpath -o /out/prepare-ci-team .

FROM otel/opentelemetry-collector-contrib:0.165.0@sha256:e40d61f10c9a617cca1ffd1be4d7e2b3a4b0f0001798ff07e58cabe82e618cea AS collector

FROM scratch
COPY --from=collector /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=collector /otelcol-contrib /usr/local/bin/otelcol-contrib
COPY --from=build /out/prepare-ci-team /usr/local/bin/prepare-ci-team
COPY --from=build --chown=10001:10001 /out/tmp /tmp
USER 10001:10001
ENV PATH=/usr/local/bin
ENTRYPOINT ["/usr/local/bin/prepare-ci-team"]
