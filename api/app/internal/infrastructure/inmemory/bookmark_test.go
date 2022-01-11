package inmemory

import (
	"reflect"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

// entity.Bookmark型のサンプルインスタンス。
func sampleBookmark() *entity.Bookmark {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := entity.NewName("example")
	uri, _ := entity.NewURI("https://example.com")
	tag1, _ := entity.NewTag("1")
	tag2, _ := entity.NewTag("2")
	tag3, _ := entity.NewTag("3")
	tags := []entity.Tag{*tag1, *tag2, *tag3}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.ID型のサンプルインスタンス。
func sampleBookmarkID() *entity.ID {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	return id
}

func TestNewBookmarkRepository_repository_Bookmark型のインスタンスを返却する(t *testing.T) {
	// when
	object := NewBookmarkRepository()
	// then
	interfaceObject := (*repository.Bookmark)(nil)
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
	expected := map[entity.ID]entity.Bookmark{}
	assert.Exactly(t, expected, concrete.store)
}

func TestNextID_entity_ID型のインスタンスを返却する(t *testing.T) {
	// given
	repository := NewBookmarkRepository()
	// when
	object := repository.NextID()
	// then
	expectedType := &entity.ID{}
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
	bookmark := (*entity.Bookmark)(nil)
	// when
	err := repository.Save(bookmark)
	// then
	errString := "argument \"bookmark\" is nil"
	assert.EqualError(t, err, errString)
}

func TestFindByID_該当するブックマークが存在する場合はentity_Bookmark型のインスタンスを返却する(t *testing.T) {
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
	id := (*entity.ID)(nil)
	// when
	object, err := repository.FindByID(id)
	// then
	assert.Nil(t, object)
	errString := "argument \"id\" is nil"
	assert.EqualError(t, err, errString)
}
