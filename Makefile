all: deps protoc

deps:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go

protoc:
	protoc --go_out=plugins=grpc:. proto/service.proto
