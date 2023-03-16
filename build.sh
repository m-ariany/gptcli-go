# Needs to be replaced with a build pipeline

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o chatgpt.exe main.go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o chatgpt_darwin_amd64 main.go

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o chatgpt_darwin_arm64 main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o chatgpt main.go