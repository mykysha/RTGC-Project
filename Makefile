go-formatter:
	go fmt main.go
	go fmt github.com/nndergunov/RTGC-Project/internal/chat
	go fmt github.com/nndergunov/RTGC-Project/internal/server

go-build-mac:
	GOOS=darwin GOARCH=amd64 go build github.com/nndergunov/RTGC-Project

go-build-win:
	GOOS=windows GOARCH=amd64 go build github.com/nndergunov/RTGC-Project

go-build-linux:
	GOOS=linux GOARCH=amd64 go build github.com/nndergunov/RTGC-Project