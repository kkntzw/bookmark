package mongodb

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	sample_entity "github.com/kkntzw/bookmark/test/data/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewBookmarkRepository_repository_Bookmark型のインスタンスを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		// when
		object := NewBookmarkRepository(collection)
		// then
		interfaceObject := (*repository.Bookmark)(nil)
		assert.Implements(t, interfaceObject, object)
		assert.NotNil(t, object)
	})
}

func TestNewBookmarkRepository_戻り値は初期化済みのフィールドcollectionを持つ(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		abstract := NewBookmarkRepository(collection)
		// when
		concrete, ok := abstract.(*bookmarkRepository)
		// then
		assert.True(t, ok)
		expectedType := &mongo.Collection{}
		assert.IsType(t, expectedType, concrete.collection)
	})
}

func TestNextID_entity_ID型のインスタンスを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		repository := NewBookmarkRepository(collection)
		// when
		object := repository.NextID()
		// then
		expectedType := &entity.ID{}
		assert.IsType(t, expectedType, object)
		assert.NotNil(t, object)
	})
}

func TestSave_正当な値を受け取るとnilを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		repository := NewBookmarkRepository(collection)
		bookmark := sample_entity.Bookmark()
		// when
		err := repository.Save(bookmark)
		// then
		assert.NoError(t, err)
	})
}

func TestSave_不正な値を受け取るとエラーを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		repository := NewBookmarkRepository(collection)
		bookmark := (*entity.Bookmark)(nil)
		// when
		err := repository.Save(bookmark)
		// then
		errString := "argument \"bookmark\" is nil"
		assert.EqualError(t, err, errString)
	})
}

func TestSave_ブックマークドキュメントの保存に失敗した場合はエラーを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
		repository := NewBookmarkRepository(collection)
		bookmark := sample_entity.Bookmark()
		// when
		err := repository.Save(bookmark)
		// then
		errString := "failed at collection.UpdateByID: command failed"
		assert.EqualError(t, err, errString)
	})
}

func TestFindByID_該当するブックマークが存在する場合はentity_Bookmark型のインスタンスを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		batch := bson.D{
			{Key: "_id", Value: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6"},
			{Key: "name", Value: "example"},
			{Key: "uri", Value: "https://example.com"},
			{Key: "tags", Value: bson.A{"1", "2", "3"}},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, batch))
		repository := NewBookmarkRepository(collection)
		id := sample_entity.BookmarkID()
		// when
		actual, err := repository.FindByID(id)
		// then
		expected := sample_entity.Bookmark()
		assert.Exactly(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestFindByID_該当するブックマークが存在しない場合はnilを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		batch := bson.D{}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, batch))
		repository := NewBookmarkRepository(collection)
		id := sample_entity.BookmarkID()
		// when
		object, err := repository.FindByID(id)
		// then
		assert.Nil(t, object)
		assert.NoError(t, err)
	})
}

func TestFindByID_不正な値を受け取るとエラーを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		repository := NewBookmarkRepository(collection)
		id := (*entity.ID)(nil)
		// when
		object, err := repository.FindByID(id)
		// then
		assert.Nil(t, object)
		errString := "argument \"id\" is nil"
		assert.EqualError(t, err, errString)
	})
}

func TestFindByID_ブックマークドキュメントの検索に失敗した場合はエラーを返却する(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		// given
		collection := mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
		repository := NewBookmarkRepository(collection)
		id := sample_entity.BookmarkID()
		// when
		object, err := repository.FindByID(id)
		// then
		assert.Nil(t, object)
		errString := "failed at collection.FindOne: command failed"
		assert.EqualError(t, err, errString)
	})
}
