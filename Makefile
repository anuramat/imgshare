.PHONY: run build deps compile
gen: deps
	protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=internal --go-grpc_out=internal api/api.proto

run: gen
	go run cmd/bot/main.go

serve: gen
	go run cmd/db/main.go

build: compile
	go build -o bin/bot cmd/bot/main.go
	go build -o bin/db cmd/db/main.go

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
