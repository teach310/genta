generate-example:
	protoc --proto_path=pb --go_out=out pb/*.proto

generate-example2:
	protoc --proto_path=pb --plugin=./protoc-gen-genta --genta_out=out pb/api/*.proto pb/*.proto

generate-types:
	protoc --proto_path=pb --go_out=. --go_opt=module=github.com/teach310/genta pb/genta/api/*.proto

build:
	go build github.com/teach310/genta/cmd/protoc-gen-genta

run: build generate-example2
