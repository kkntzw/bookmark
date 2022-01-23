package di

import (
	"github.com/kkntzw/bookmark/internal/domain/service"
)

// ブックマークに関するドメインサービスを注入する。
func InjectBookmarkService() service.Bookmark {
	return service.NewBookmarkService(
		InjectMongoDBBookmarkRepository(),
	)
}

// ブックマークに関するテスト用ドメインサービスを注入する。
func InjectTestBookmarkService() service.Bookmark {
	return service.NewBookmarkService(
		InjectInMemoryBookmarkRepository(),
	)
}
