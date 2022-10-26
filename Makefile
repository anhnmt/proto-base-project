
buf.gen:
	buf generate

buf.update:
	cd proto/ && buf mod update

go.get:
	go get -u ./...

go.gen: buf.gen

go.tidy:
	go mod tidy -compat=1.19

go.test:
	go test ./...

go.install:
	go install github.com/bufbuild/buf/cmd/buf@v1.9.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.12.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

