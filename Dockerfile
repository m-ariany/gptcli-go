# Build layer
FROM golang:1.19.1 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -race -a -o gptcli cmd/cli/main.go

# Run layer
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /app/gptcli .
USER 65532:65532
ENTRYPOINT ["/app/gptcli"]
