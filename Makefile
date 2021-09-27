go-lint:
	gofmt -l -s -w .
	golangci-lint run --enable-all

go-formatter:
	gofumpt -l -s -w .

go-server:
	go run cmd/server/main.go

go-client:
	go run cmd/client/main.go

go-build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/server cmd/server/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/client cmd/client/main.go

go-build-win:
	GOOS=windows GOARCH=amd64 go build -o bin/win/server.exe cmd/server/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/win/client.exe cmd/client/main.go

go-build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server cmd/server/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/client cmd/client/main.go