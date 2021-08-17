go-formatter:
	go fmt main.go
	go fmt internal/chat/hub.go
	go fmt internal/chat/init.go
	go fmt internal/chat/room.go
	go fmt internal/chat/server.go
	go fmt internal/chat/session.go

go-build:
	GOOS=darwin GOARCH=amd64 go build -o RTGC-Project