package di

import (
	"os"

	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/kkntzw/bookmark/internal/infrastructure/inmemory"
	"github.com/kkntzw/bookmark/internal/infrastructure/mongodb"
)

var (
	inMemoryBookmarkRepository repository.Bookmark // ブックマークを扱うインメモリ型リポジトリ
	mongoDbBookmarkRepository  repository.Bookmark // ブックマークを扱うMongoDBリポジトリ
)

// ブックマークの永続化を担うインメモリ型リポジトリを注入する。
func InjectInMemoryBookmarkRepository() repository.Bookmark {
	return inMemoryBookmarkRepository
}

// ブックマークの永続化を担うMongoDBリポジトリを注入する。
func InjectMongoDBBookmarkRepository() repository.Bookmark {
	return mongoDbBookmarkRepository
}

// シングルトンでインスタンスを扱うために初期化する。
func init() {
	inMemoryBookmarkRepository = inmemory.NewBookmarkRepository()

	db := mongodb.NewMongoDatabase(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE"))
	collection := db.Collection(os.Getenv("MONGO_COLLECTION"))
	mongoDbBookmarkRepository = mongodb.NewBookmarkRepository(collection)
}
