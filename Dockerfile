FROM golang:1.20.6-alpine3.18 AS base

WORKDIR /var/code
COPY ./ ./

RUN \
  CGO_ENABLED=0 GOOS=linux go build \
  -installsuffix "static" \
  -o /usr/local/bin/supergood-proxy \
  ./cmd/main.go

FROM alpine:3.18.2 AS app
COPY _config/ /var/_config/
COPY --from=base /usr/local/bin/supergood-proxy /usr/local/bin/supergood-proxy
# COPY --from=datadog/serverless-init:1 /datadog-init /app/datadog-init
# ENTRYPOINT ["/app/datadog-init"]
ENTRYPOINT ["/usr/local/bin/supergood-proxy"]
