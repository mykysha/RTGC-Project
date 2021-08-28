go-formatter:
	go fmt main.go

	go fmt github.com/nndergunov/RTGC-Project/api
	go fmt github.com/nndergunov/RTGC-Project/api/v1
	go fmt github.com/nndergunov/RTGC-Project/api/v1/client

	go fmt github.com/nndergunov/RTGC-Project/cmd/server

	go fmt github.com/nndergunov/RTGC-Project/internal/chat
	go fmt github.com/nndergunov/RTGC-Project/internal/httpserver

go-client:
	go run api/v1/client/main.go

go-build-mac:
	GOOS=darwin GOARCH=amd64 go build github.com/nndergunov/RTGC-Project

go-build-win:
	GOOS=windows GOARCH=amd64 go build github.com/nndergunov/RTGC-Project

go-build-linux:
	GOOS=linux GOARCH=amd64 go build github.com/nndergunov/RTGC-Project