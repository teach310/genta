generate-example:
	protoc --proto_path=pb --go_out=out pb/*.proto
