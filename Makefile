#
# Basic making steps
#
#

clean:
	rm -rf ./bin
	go clean -x

build-grpc:
	go build -o ./bin/snapurl-grpc ./cmd/snapurl-grpc

build-cli:
	go build -o ./bin/snapurl ./cmd/snapurl

build-docker:
	docker build -t snapurl-cli --file cmd/snapurl/Dockerfile .
	docker build -t snapurl-grpc --file cmd/snapurl-grpc/Dockerfile .

#
# OS Specific
#
#

build-cli-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/snapurl-cli-linux ./cmd/snapurl

build-cli-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/snapurl-cli-windows ./cmd/snapurl

build-cli-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/snapurl-cli-darwin ./cmd/snapurl

all: clean build-grpc build-cli-linux build-cli-windows build-cli-darwin

#
# Protobuf definitions
#
#

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
