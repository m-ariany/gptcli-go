# Needs to be replaced with a build pipeline

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -race -a -o gptcli-go cmd/cli/main.go

CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -race -a -o gptcli-go cmd/cli/main.go

CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -race -a -o gptcli-go cmd/cli/main.go

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -race -a -o gptcli-go cmd/cli/main.go