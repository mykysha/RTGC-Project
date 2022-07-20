# Real-Time Group Chat

**Built and executed on Go v1.17**

## Needed preparations

**Download and install all resources**

install `Go`

## Local run:

For running server in your IDE, in terminal...

```
make go-server
```

For running client in your IDE, in terminal...

```
make go-client
```

## Server:

Port:

```
:8080
```

#### Endpoints:

* status check:

```
/v1/status
```

* websocket connection upgrader:

```
/v1/ws
```

## Client:

#### Available commands:

* join:

```
join:_RoomName_:_UserName_
```

* send:
```
join:_RoomName_:_Text_
```

* leave:
```
leave:_RoomName_:_ReasonToLeave_
```


## Useful makefile commands:

To run client run

```
make go-client
```

To build for windows-amd64 run (creates folder in bin/win)

```
make go-build-win
```

To build for mac-amd64 run (creates folder in bin/mac)

```
make go-build-mac
```

To build for linux-amd64 run (creates folder in bin/linux)

```
make go-build-linux
```

To format/beautify all code run

```
make go-formatter
```
