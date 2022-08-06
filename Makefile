.PHONY: runbot runddb build gen dep

runbot: gen
	go run cmd/bot/main.go

rundb: gen
	go run cmd/db/main.go

build: gen
	go build -o bin/bot cmd/bot/main.go
	go build -o bin/db cmd/db/main.go

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