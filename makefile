# https://stackoverflow.com/questions/2145590/what-is-the-purpose-of-phony-in-a-makefile
.PHONY: lint

lint:
	golangci-lint run -c .golangci.yml

proto_gen:
	protoc --go_out=. --go-grpc_out=. pb/*.proto

http_load_test:
	k6 run .\loadtest\http.js