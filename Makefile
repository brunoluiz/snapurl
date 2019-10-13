build-grpc:
	go build -o ./bin/grpc ./cmd/grpc

build-cli-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/snapurl ./cmd/cli

build-cli-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/snapurl ./cmd/cli

build-cli-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/snapurl ./cmd/cli

all: build-grpc build-cli-linux build-cli-windows build-cli-darwin

protos:
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--swagger_out=logtostderr=true:. \
		api/api.proto
