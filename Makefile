run-client:

	docker run -it client_app

rebuild-docker:
	docker-compose down --remove-orphans --volumes
	docker-compose up --build

update-docker:
	docker build -t rtgc-project_server .
	docker run -dp

down-docker:
	docker-compose down

up-docker:
	docker-compose up -d

go-lint:
	gofmt -l -s -w .
	golangci-lint run --enable-all

go-formatter:
	gofumpt -l -s -w .

go-gci:
	gci -local "github.com/nndergunov" -w .

go-server:
	go run server/cmd/main.go

go-client:
	go run client/main.go

go-unit-test:
	go test pkg/app/room/room_internal_test.go -v

go-integration-test:
	go test pkg/db/integration_test.go -v

go_e2e-test:
	go test tests/e2eserver_test.go -v

go-clearDB:
	go run server/pkg/db/clear/main.go

go-build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/server server/cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/client client/main.go

go-build-win:
	GOOS=windows GOARCH=amd64 go build -o bin/win/server.exe server/cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/win/client.exe client/main.go

go-build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server server/cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/client client/main.go