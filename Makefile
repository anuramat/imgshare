.PHONY: build mock generate dependencies clean

build: generate
	go build -o bin/imgsharebot cmd/imgsharebot/main.go
	go build -o bin/imgshare cmd/imgshare/main.go

mock: generate
	./bin/mockgen -source=internal/api/api_grpc.pb.go -destination=internal/mocks/grpc_mocks.go -package=mocks
	./bin/mockgen -source=internal/imgshare/server.go -destination=internal/mocks/db_mocks.go -package=mocks

generate: dependencies
	buf mod update
	buf generate

LOCAL_GOPATH:=$(CURDIR)/bin
dependencies: 
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_GOPATH) go install github.com/pressly/goose/v3/cmd/goose && \
	GOBIN=$(LOCAL_GOPATH) go install github.com/golang/mock/mockgen

clean:
	rm -rf ./bin
	rm -rf ./internal/api
	rm -rf ./swagger/api
