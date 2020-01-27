
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off
GOOS?=darwin
GOARCH=amd64

grpc:
	protoc -I ./api/ -I ${GOPATH}/src --go_out=plugins=grpc:./api ./api/config.proto

run-cluster:
	go run -race ./main.go

run-client:
	go run ./client/main.go