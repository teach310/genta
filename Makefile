generate-example:
	protoc --proto_path=pb --go_out=out pb/*.proto

generate-example2:
	@mkdir -p out
	protoc --proto_path=pb --plugin=./protoc-gen-genta --genta_opt=templates_path=templates --genta_out=out pb/*.proto

build:
	go build github.com/teach310/genta/cmd/protoc-gen-genta

run: build generate-example2

debug_proto: build
	@mkdir -p out
	protoc --proto_path=pb --plugin=./protoc-gen-genta \
	--genta_opt=templates_path=templates \
	--genta_opt=model=txt \
	--genta_out=out pb/*.proto
