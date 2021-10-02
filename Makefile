generate:
	@protoc --proto_path=. --go_out=plugins=grpc:api/app/application/ proto/bookmark.proto
