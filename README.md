Test case for https://github.com/docker/for-mac/issues/1863

To reproduce, run:

```sh
$ docker build .
Sending build context to Docker daemon  43.01kB
Step 1/4 : FROM golang:1.8-alpine
 ---> 310e63753884
Step 2/4 : WORKDIR /opt/hello-go
 ---> Using cache
 ---> 7a21a7e60130
Step 3/4 : COPY . .
 ---> 32d25ddacdc6
Removing intermediate container 00c11fd40f97
Step 4/4 : RUN go test -v ./...
 ---> Running in 02c6ccd397f9
=== RUN   TestConnectingToAServer
--- FAIL: TestConnectingToAServer (0.00s)
	server_test.go:27: Got an error: Get http://[::]:36703: dial tcp [::]:36703: connect: cannot assign requested address
FAIL
exit status 1
FAIL	_/opt/hello-go	0.005s
The command '/bin/sh -c go test -v ./...' returned a non-zero code: 1
```

But running outside of Docker seems to work fine:

```sh
$ go test ./...
ok  	_/Users/jabley/Projects/jabley/docker-networking-repro	0.010s
```
