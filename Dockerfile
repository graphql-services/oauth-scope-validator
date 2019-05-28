FROM golang as builder
WORKDIR /go/src/github.com/graphql-services/oauth-scope-validator
COPY . .
RUN go get ./... 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /tmp/app main.go
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    curl -L https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 -o /tmp/grpc_health_probe

FROM alpine:3.5

WORKDIR /app

COPY --from=builder /tmp/app /usr/local/bin/app
COPY --from=builder /tmp/grpc_health_probe /bin/grpc_health_probe

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /bin/grpc_health_probe && chmod +x /usr/local/bin/app

ENTRYPOINT []
CMD [ "app", "server" ]