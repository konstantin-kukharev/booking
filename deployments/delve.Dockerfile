## Builder
FROM golang:1.23-alpine3.19 AS dev-builder

ARG APPLICATION_NAME

# Create a workspace for the app
WORKDIR /app
ADD ../ ./

# Build
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o ./application ./cmd/$APPLICATION_NAME
RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest

## Runner
FROM scratch AS dev-runner

WORKDIR /
COPY --from=dev-builder /go/bin/dlv /dlv
COPY --from=dev-builder /app/application /application

ENTRYPOINT [ "/dlv" ]