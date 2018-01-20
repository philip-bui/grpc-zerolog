PORT?=8000
PACKAGE:=github.com/philip-bui/grpc-zerolog
COVERAGE:=coverage.txt
proto:
	protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

godoc:
	echo "localhost:${PORT}/pkg/${PACKAGE}"
	godoc -http=:${PORT}

.PHONY: test

test:
	go test -race -coverprofile=${COVERAGE} -covermode=atomic

coverage:
	go tool cover -html=${COVERAGE}
