package mongodb

import (
	"errors"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func ToID(t *testing.T, v string) *entity.ID {
	t.Helper()
	id, err := entity.NewID(v)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func ToBookmark(t *testing.T, iv, nv, uv string, tvs ...string) *entity.Bookmark {
	t.Helper()
	id, err := entity.NewID(iv)
	if err != nil {
		t.Fatal(err)
	}
	name, err := entity.NewName(nv)
	if err != nil {
		t.Fatal(err)
	}
	uri, err := entity.NewURI(uv)
	if err != nil {
		t.Fatal(err)
	}
	tags := make([]entity.Tag, len(tvs))
	for i, tv := range tvs {
		tag, err := entity.NewTag(tv)
		if err != nil {
			t.Fatal(err)
		}
		tags[i] = *tag
	}
	bookmark, err := entity.NewBookmark(id, name, uri, tags)
	if err != nil {
		t.Fatal(err)
	}
	return bookmark
}

func ToBatch(t *testing.T, id, name, uri string, tags ...string) bson.D {
	t.Helper()
	a := bson.A{}
	for _, tag := range tags {
		a = append(a, tag)
	}
	bookmark := bson.D{
		{Key: "_id", Value: id},
		{Key: "name", Value: name},
		{Key: "uri", Value: uri},
		{Key: "tags", Value: a},
	}
	return bookmark
}

func TestNewBookmarkRepository(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	{
		mt.Run("implementing bookmark repository", func(mt *mtest.T) {
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
	}
	{
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
			ToBookmark(t, "1", "Example A", "https://foo.example.com"),
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
			ToBookmark(t, "1", "Example A", "https://foo.example.com"),
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
	batch1 := ToBatch(t, "1", "Example A", "https://foo.example.com")
	batch2 := ToBatch(t, "2", "Example B", "https://bar.example.com")
	batch3 := ToBatch(t, "3", "Example C", "https://baz.example.com")
	bookmark1 := ToBookmark(t, "1", "Example A", "https://foo.example.com")
	bookmark2 := ToBookmark(t, "2", "Example B", "https://bar.example.com")
	bookmark3 := ToBookmark(t, "3", "Example C", "https://baz.example.com")
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare           func(*mtest.T)
		expectedBookmarks []entity.Bookmark
		expectedErr       error
	}{
		"stored bookmarks": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, batch1))
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, batch2))
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, batch3))
				mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch))
			},
			[]entity.Bookmark{*bookmark1, *bookmark2, *bookmark3},
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
	batch := ToBatch(t, "1", "Example A", "https://foo.example.com")
	id := ToID(t, "1")
	bookmark := ToBookmark(t, "1", "Example A", "https://foo.example.com")
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	cases := map[string]struct {
		prepare          func(*mtest.T)
		id               *entity.ID
		expectedBookmark *entity.Bookmark
		expectedErr      error
	}{
		"stored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, batch))
			},
			id,
			bookmark,
			nil,
		},
		"unstored bookmark": {
			func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			},
			id,
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
			id,
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
