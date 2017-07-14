# Our base image is Go 1.8 on Alpine Linux.
FROM golang:1.8-alpine

# Establish a working directory and copy our application files into it.
WORKDIR /opt/hello-go
COPY . .

# Build your application.
RUN go test -v ./...
