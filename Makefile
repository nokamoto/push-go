all: deps protoc test install

deps:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go

protoc:
	protoc --go_out=plugins=grpc:. proto/service.proto

test:
	go test ./push-go-app
	go test ./push-go-endpoint
	go test ./push-go-notification

install:
	go install ./push-go-app
	go install ./push-go-endpoint
	go install ./push-go-notification
