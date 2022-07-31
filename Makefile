.PHONY: run build gen serve .dep
gen: .dep
	buf generate

run: gen
	go run cmd/bot/main.go

serve: gen
	go run cmd/db/main.go

build: gen
	go build -o bin/bot cmd/bot/main.go
	go build -o bin/db cmd/db/main.go

LOCAL_GOPATH:=$(CURDIR)/bin
.dep: 
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

