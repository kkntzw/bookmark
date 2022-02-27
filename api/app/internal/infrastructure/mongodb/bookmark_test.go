package mongodb

import (
	"errors"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/kkntzw/bookmark/test/helper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewBookmarkRepository(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("implementing repository.Bookmark", func(mt *mtest.T) {
		mt.Parallel()
		// given
		collection := mt.Coll
		// when
		object := NewBookmarkRepository(collection)
		// then
		assert.NotNil(mt, object)
		interfaceObject := (*repository.Bookmark)(nil)
		assert.Implements(mt, interfaceObject, object)
	})
	mt.Run("fields", func(mt *mtest.T) {
		mt.Parallel()
		// given
		collection := mt.Coll
		abstractRepository := NewBookmarkRepository(collection)
		// when
		concreteRepository, ok := abstractRepository.(*bookmarkRepository)
		actualCollection := concreteRepository.collection
		// then
		assert.True(mt, ok)
		expectedCollection := collection
		assert.Exactly(mt, expectedCollection, actualCollection)
	})
}

func TestBookmark_NextID(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	// given
	collection := mt.Coll
	repository := NewBookmarkRepository(collection)
	// when
	id := repository.NextID()
	// then
	assert.NotNil(t, id)
	expectedType := &entity.ID{}
	assert.IsType(t, expectedType, id)
}

func TestBookmark_Save(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare     func(*mtest.T)
		bookmark    *entity.Bookmark
		expectedErr error
	}{
		"non-nil bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"nil bookmark": {
			func(mt *mtest.T) {},
			nil,
			errors.New("argument \"bookmark\" is nil"),
		},
		"failed at collection.UpdateByID": {
			func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			errors.New("failed at collection.UpdateByID: command failed"),
		},
	}
	for name, tc := range cases {
		tc := tc
		mt.Run(name, func(mt *mtest.T) {
			mt.Parallel()
			tc.prepare(mt)
			// given
			collection := mt.Coll
			repository := NewBookmarkRepository(collection)
			// when
			actualErr := repository.Save(tc.bookmark)
			// then
			if tc.expectedErr == nil {
				assert.NoError(mt, actualErr)
			} else {
				assert.Exactly(mt, tc.expectedErr.Error(), actualErr.Error())
			}
		})
	}
}

func TestBookmark_FindAll(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare           func(*mtest.T)
		expectedBookmarks []entity.Bookmark
		expectedErr       error
	}{
		"stored bookmarks": {
			func(mt *mtest.T) {
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, helper.ToBookmarkDocument(t, "1", "Example A", "https://foo.example.com")),
				)
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, helper.ToBookmarkDocument(t, "2", "Example B", "https://bar.example.com")),
				)
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, helper.ToBookmarkDocument(t, "3", "Example C", "https://baz.example.com")),
				)
				mt.AddMockResponses(
					mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch),
				)
			},
			[]entity.Bookmark{
				*helper.ToBookmark(t, "1", "Example A", "https://foo.example.com"),
				*helper.ToBookmark(t, "2", "Example B", "https://bar.example.com"),
				*helper.ToBookmark(t, "3", "Example C", "https://baz.example.com"),
			},
			nil,
		},
		"unstored bookmarks": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
			},
			[]entity.Bookmark{},
			nil,
		},
		"failed at collection.Find": {
			func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			nil,
			errors.New("failed at collection.Find: command failed"),
		},
		"failed at cursor.All": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			},
			nil,
			errors.New("failed at cursor.All: no responses remaining"),
		},
	}
	for name, tc := range cases {
		tc := tc
		mt.Run(name, func(mt *mtest.T) {
			mt.Parallel()
			tc.prepare(mt)
			// given
			collection := mt.Coll
			repository := NewBookmarkRepository(collection)
			// when
			actualBookmarks, actualErr := repository.FindAll()
			// then
			assert.ElementsMatch(mt, tc.expectedBookmarks, actualBookmarks)
			if tc.expectedErr == nil {
				assert.NoError(mt, actualErr)
			} else {
				assert.Exactly(mt, tc.expectedErr.Error(), actualErr.Error())
			}
		})
	}
}

func TestBookmark_FindByID(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare          func(*mtest.T)
		id               *entity.ID
		expectedBookmark *entity.Bookmark
		expectedErr      error
	}{
		"id of stored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, helper.ToBookmarkDocument(t, "1", "Example", "https://example.com")),
				)
			},
			helper.ToID(t, "1"),
			helper.ToBookmark(t, "1", "Example", "https://example.com"),
			nil,
		},
		"id of unstored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			},
			helper.ToID(t, "1"),
			nil,
			nil,
		},
		"nil id": {
			func(mt *mtest.T) {},
			nil,
			nil,
			errors.New("argument \"id\" is nil"),
		},
		"failed at collection.FindOne": {
			func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			helper.ToID(t, "1"),
			nil,
			errors.New("failed at collection.FindOne: command failed"),
		},
	}
	for name, tc := range cases {
		tc := tc
		mt.Run(name, func(mt *mtest.T) {
			mt.Parallel()
			tc.prepare(mt)
			// given
			collection := mt.Coll
			repository := NewBookmarkRepository(collection)
			// when
			actualBookmark, actualErr := repository.FindByID(tc.id)
			// then
			assert.Exactly(mt, tc.expectedBookmark, actualBookmark)
			if tc.expectedErr == nil {
				assert.NoError(mt, actualErr)
			} else {
				assert.Exactly(mt, tc.expectedErr.Error(), actualErr.Error())
			}
		})
	}
}

func TestBookmark_Delete(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare     func(*mtest.T)
		bookmark    *entity.Bookmark
		expectedErr error
	}{
		"stored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "acknowledged", Value: true}, bson.E{Key: "n", Value: 1}))
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"unstored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "acknowledged", Value: true}, bson.E{Key: "n", Value: 0}))
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"nil bookmark": {
			func(mt *mtest.T) {},
			nil,
			errors.New("argument \"bookmark\" is nil"),
		},
		"failed at collection.DeleteOne": {
			func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			errors.New("failed at collection.DeleteOne: command failed"),
		},
	}
	for name, tc := range cases {
		tc := tc
		mt.Run(name, func(mt *mtest.T) {
			mt.Parallel()
			tc.prepare(mt)
			// given
			collection := mt.Coll
			repository := NewBookmarkRepository(collection)
			// when
			actualErr := repository.Delete(tc.bookmark)
			// then
			if tc.expectedErr == nil {
				assert.NoError(mt, actualErr)
			} else {
				assert.Exactly(mt, tc.expectedErr.Error(), actualErr.Error())
			}
		})
	}
}
