# Real-Time Group Chat

**This project is still under construction**
-
**Built and executed on Go v1.17**


**Needed preparations**
-
**Download and install all resources**

install `Go`


**Local run:**
-
>For running server in your IDE, in terminal...
>
* `go run main.go`

>For running client in your IDE, in terminal...
>
* `make go-client`


**Server:**
-
>Port:
>
* `:8080`

>Endpoints:
>
* `/v1/status`

* `/v1/ws`


**Useful makefile commands:**
-
>To run client run
>
* `make go-client`

>To build for windows-amd64 run
>
* `make go-build-win`

>To build for mac-amd64 run
>
* `make go-build-mac`

>To build for linux-amd64 run
>
* `make go-build-linux`

>To format/beautify all code run
>
* `make go-formatter`