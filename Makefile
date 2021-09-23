generate:
	@protoc --proto_path=. --go_out=plugins=grpc:api/app/domain/ proto/bookmark.proto
