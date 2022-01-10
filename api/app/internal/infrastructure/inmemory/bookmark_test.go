package inmemory

import (
	"reflect"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/bookmark"
	"github.com/stretchr/testify/assert"
)

// bookmark.Bookmark型のサンプルインスタンス。
func sampleBookmark() *bookmark.Bookmark {
	id, _ := bookmark.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := bookmark.NewName("example")
	uri, _ := bookmark.NewURI("https://example.com")
	tag1, _ := bookmark.NewTag("1")
	tag2, _ := bookmark.NewTag("2")
	tag3, _ := bookmark.NewTag("3")
	tags := []bookmark.Tag{*tag1, *tag2, *tag3}
	bookmark, _ := bookmark.New(id, name, uri, tags)
	return bookmark
}

// bookmark.ID型のサンプルインスタンス。
func sampleBookmarkID() *bookmark.ID {
	id, _ := bookmark.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	return id
}

func TestNewBookmarkRepository_bookmark_Repository型のインスタンスを返却する(t *testing.T) {
	// when
	object := NewBookmarkRepository()
	// then
	interfaceObject := (*bookmark.Repository)(nil)
	assert.Implements(t, interfaceObject, object)
	assert.NotNil(t, object)
}

func TestNewBookmarkRepository_戻り値は初期化済みのフィールドstoreを持つ(t *testing.T) {
	// given
	abstract := NewBookmarkRepository()
	// when
	concrete, ok := abstract.(*bookmarkRepository)
	// then
	assert.True(t, ok)
	expected := map[bookmark.ID]bookmark.Bookmark{}
	assert.Exactly(t, expected, concrete.store)
}

func TestNextID_bookmark_ID型のインスタンスを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	// when
	object := repository.NextID()
	// then
	expectedType := &bookmark.ID{}
	assert.IsType(t, expectedType, object)
	assert.NotNil(t, object)
}

func TestSave_正当な値を受け取るとnilを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	bookmark := sampleBookmark()
	// when
	err := repository.Save(bookmark)
	// then
	assert.NoError(t, err)
}

func TestSave_引数bookmarkとフィールドstoreに保存した値は同一でないが同値となる(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	bookmark := sampleBookmark()
	repository.Save(bookmark)
	// when
	concrete, _ := repository.(*bookmarkRepository)
	stored := concrete.store[bookmark.ID()]
	same := bookmark == &stored
	equiv := reflect.DeepEqual(*bookmark, stored)
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestSave_不正な値を受け取るとエラーを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	bookmark := (*bookmark.Bookmark)(nil)
	// when
	err := repository.Save(bookmark)
	// then
	errString := "argument \"bookmark\" is nil"
	assert.EqualError(t, err, errString)
}

func TestFindByID_該当するブックマークが存在する場合はbookmark_Bookmark型のインスタンスを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	repository.Save(sampleBookmark())
	id := sampleBookmarkID()
	// when
	actual, err := repository.FindByID(id)
	// then
	expected := sampleBookmark()
	assert.Exactly(t, expected, actual)
	assert.NoError(t, err)
}

func TestFindByID_戻り値bookmarkとフィールドstoreに保存した値は同一でないが同値となる(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	repository.Save(sampleBookmark())
	id := sampleBookmarkID()
	bookmark, _ := repository.FindByID(id)
	// when
	concrete, _ := repository.(*bookmarkRepository)
	stored := concrete.store[bookmark.ID()]
	same := bookmark == &stored
	equiv := reflect.DeepEqual(*bookmark, stored)
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestFindByID_該当するブックマークが存在しない場合はnilを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	id := sampleBookmarkID()
	// when
	object, err := repository.FindByID(id)
	// then
	assert.Nil(t, object)
	assert.NoError(t, err)
}

func TestFindByID_不正な値を受け取るとエラーを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	id := (*bookmark.ID)(nil)
	// when
	object, err := repository.FindByID(id)
	// then
	assert.Nil(t, object)
	errString := "argument \"id\" is nil"
	assert.EqualError(t, err, errString)
}
