generate-example:
	protoc --proto_path=pb --go_out=out pb/*.proto

build:
	go build github.com/teach310/genta/cmd/protoc-gen-genta
