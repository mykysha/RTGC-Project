FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ../go.mod go.sum ./

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/client .
COPY --from=builder /app/.env .

ENTRYPOINT ["./client"]