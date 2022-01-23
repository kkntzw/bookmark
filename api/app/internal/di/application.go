package di

import (
	"github.com/kkntzw/bookmark/internal/application/usecase"
)

// ブックマークに関するユースケースを注入する。
func InjectBookmarkUsecase() usecase.Bookmark {
	return usecase.NewBookmarkUsecase(
		InjectMongoDBBookmarkRepository(),
		InjectBookmarkService(),
	)
}

// ブックマークに関するテスト用ユースケースを注入する。
func InjectTestBookmarkUsecase() usecase.Bookmark {
	return usecase.NewBookmarkUsecase(
		InjectInMemoryBookmarkRepository(),
		InjectTestBookmarkService(),
	)
}
