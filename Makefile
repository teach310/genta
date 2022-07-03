generate-example:
	protoc --proto_path=pb --go_out=out pb/*.proto

generate-example2:
	protoc --proto_path=pb --plugin=./protoc-gen-genta --genta_out=out pb/*.proto

build:
	go build github.com/teach310/genta/cmd/protoc-gen-genta
