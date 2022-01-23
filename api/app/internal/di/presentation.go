package di

import (
	"github.com/kkntzw/bookmark/internal/presentation/pb"
	"github.com/kkntzw/bookmark/internal/presentation/server"
)

// ブックマークに関するgRPCサーバを注入する。
func InjectBookmarkServer() pb.BookmarkerServer {
	return server.NewBookmarkServer(
		InjectBookmarkUsecase(),
	)
}

// ブックマークに関するテスト用gRPCサーバを注入する。
func InjectTestBookmarkServer() pb.BookmarkerServer {
	return server.NewBookmarkServer(
		InjectTestBookmarkUsecase(),
	)
}
