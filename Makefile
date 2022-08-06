.PHONY: build gen dep

build: gen
	go build -o bin/imgshare_bot cmd/imgshare_bot/main.go
	go build -o bin/imgshare cmd/imgshare/main.go

gen: dep
	buf mod update
	buf generate

LOCAL_GOPATH:=$(CURDIR)/bin
dep: 
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_GOPATH) go install github.com/pressly/goose/v3/cmd/goose