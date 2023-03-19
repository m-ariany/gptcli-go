FROM golang:1.19.1

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app
RUN go build -a -o /app/gptcli cmd/cli/main.go

CMD ["/app/gptcli"]
