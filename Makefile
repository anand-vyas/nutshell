.PHONY: proto

proto:
	protoc --go_out=paths=source_relative:. ./nutshellpb/*.proto

