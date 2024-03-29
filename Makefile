
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
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

