# event-gateway
Receives events via grpc

# Setup
## name should be unique, convention is to use name of github repo
go mod init github.com/jeremyseow/event-gateway

# Dependencies
go get google.golang.org/grpc
go get google.golang.org/protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# If you want to download all dependencies
go mod vendor

# Protobuf
protoc --go_out=. --go-grpc_out=. pb/*.proto

# Run
In the root of project
go run main.go