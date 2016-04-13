
all: hello.pb.go

hello.pb.go: hello.proto
	protoc -I . hello.proto --go_out=plugins=grpc:.

.PHONY: all
