.PHONY: run build3 compile gen .dep
gen: .dep
	protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=internal --go-grpc_out=internal api/api.proto

	protoc -I . --grpc-gateway_out ./internal \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    api/api.proto

	protoc -I . --openapiv2_out ./openapi \
    --openapiv2_opt logtostderr=true \
 	api/api.proto

run: gen
	go run cmd/bot/main.go

serve: gen
	go run cmd/db/main.go

build: compile
	go build -o bin/bot cmd/bot/main.go
	go build -o bin/db cmd/db/main.go

LOCAL_GOPATH:=$(CURDIR)/BIN
.dep: 
	GOBIN=$(LOCAL_GOPATH) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_GOPATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

