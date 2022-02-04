package entity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 関数 NewBookmark() のサンプル引数。
func args() (*ID, *Name, *URI, []Tag) {
	id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := NewName("example")
	uri, _ := NewURI("https://example.com")
	tag1, _ := NewTag("1")
	tag2, _ := NewTag("2")
	tag3, _ := NewTag("3")
	tags := []Tag{*tag1, *tag2, *tag3}
	return id, name, uri, tags
}

func TestNewBookmark_正当な値を受け取るとBookmark型のインスタンスを返却する(t *testing.T) {
	tag1, _ := NewTag("1")
	tag2, _ := NewTag("2")
	tag3, _ := NewTag("3")
	params := []struct {
		tags []Tag
	}{
		{tags: []Tag{}},
		{tags: []Tag{*tag1}},
		{tags: []Tag{*tag1, *tag2}},
		{tags: []Tag{*tag1, *tag2, *tag3}},
	}
	for _, p := range params {
		// given
		id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
		name, _ := NewName("example")
		uri, _ := NewURI("https://example.com")
		tags := p.tags
		// when
		actual, err := NewBookmark(id, name, uri, tags)
		// then
		expected := &Bookmark{*id, *name, *uri, tags}
		assert.Exactly(t, expected, actual)
		assert.NoError(t, err)
	}
}

func TestNewBookmark_不正な値を受け取るとエラーを返却する(t *testing.T) {
	id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := NewName("example")
	uri, _ := NewURI("https://example.com")
	tags := []Tag{}
	params := []struct {
		id        *ID
		name      *Name
		uri       *URI
		tags      []Tag
		errString string
	}{
		{id: nil, name: name, uri: uri, tags: tags, errString: "argument \"id\" is nil"},
		{id: id, name: nil, uri: uri, tags: tags, errString: "argument \"name\" is nil"},
		{id: id, name: name, uri: nil, tags: tags, errString: "argument \"uri\" is nil"},
		{id: id, name: name, uri: uri, tags: nil, errString: "argument \"tags\" is nil"},
	}
	for _, p := range params {
		// given
		id := p.id
		name := p.name
		uri := p.uri
		tags := p.tags
		// when
		object, err := NewBookmark(id, name, uri, tags)
		// then
		assert.Nil(t, object)
		assert.EqualError(t, err, p.errString)
	}
}

func TestNewBookmark_引数tagsとフィールドtagsは同一でないが同値となる(t *testing.T) {
	// given
	id, name, uri, tags := args()
	bookmark, _ := NewBookmark(id, name, uri, tags)
	// when
	same := reflect.ValueOf(tags).Pointer() == reflect.ValueOf(bookmark.tags).Pointer()
	equiv := reflect.DeepEqual(tags, bookmark.tags)
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestID_フィールドidを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	// when
	actual := bookmark.ID()
	// then
	id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	expected := *id
	assert.Exactly(t, expected, actual)
}

func TestName_フィールドnameを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	// when
	actual := bookmark.Name()
	// then
	name, _ := NewName("example")
	expected := *name
	assert.Exactly(t, expected, actual)
}

func TestURI_フィールドuriを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	// when
	actual := bookmark.URI()
	// then
	uri, _ := NewURI("https://example.com")
	expected := *uri
	assert.Exactly(t, expected, actual)
}

func TestTags_フィールドtagsを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	// when
	actual := bookmark.Tags()
	// then
	tag1, _ := NewTag("1")
	tag2, _ := NewTag("2")
	tag3, _ := NewTag("3")
	expected := []Tag{*tag1, *tag2, *tag3}
	assert.Exactly(t, expected, actual)
}

func TestTags_戻り値とフィールドtagsは同一でないが同値となる(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	tags := bookmark.Tags()
	// when
	same := reflect.ValueOf(tags).Pointer() == reflect.ValueOf(bookmark.tags).Pointer()
	equiv := reflect.DeepEqual(tags, bookmark.tags)
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestRename_正当な値を受け取るとフィールドnameを変更してnilを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	name, _ := NewName("EXAMPLE")
	// when
	err := bookmark.Rename(name)
	// then
	expected := *name
	actual := bookmark.Name()
	assert.Exactly(t, expected, actual)
	assert.NoError(t, err)
}

func TestRename_不正な値を受け取るとフィールドnameを変更せずエラーを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	name := (*Name)(nil)
	// when
	err := bookmark.Rename(name)
	// then
	errString := "argument \"name\" is nil"
	assert.EqualError(t, err, errString)
}

func TestRewriteURI_正当な値を受け取るとフィールドuriを変更してnilを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	uri, _ := NewURI("http://example.com")
	// when
	err := bookmark.RewriteURI(uri)
	// then
	expected := *uri
	actual := bookmark.URI()
	assert.Exactly(t, expected, actual)
	assert.NoError(t, err)
}

func TestRename_不正な値を受け取るとフィールドuriを変更せずエラーを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	uri := (*URI)(nil)
	// when
	err := bookmark.RewriteURI(uri)
	// then
	errString := "argument \"uri\" is nil"
	assert.EqualError(t, err, errString)
}

func TestDeepCopy_同じ値で異なるポインタを持つBookmark型のインスタンスを返却する(t *testing.T) {
	// given
	bookmark, _ := NewBookmark(args())
	// when
	copy := bookmark.DeepCopy()
	// then
	assert.Exactly(t, bookmark, copy)
	assert.NotSame(t, bookmark, copy)
}

func TestDeepCopy_オリジナルとコピーのフィールドtagsは同一でないが同値となる(t *testing.T) {
	// given
	original, _ := NewBookmark(args())
	copy := original.DeepCopy()
	// when
	same := reflect.ValueOf(original.tags).Pointer() == reflect.ValueOf(copy.tags).Pointer()
	equiv := reflect.DeepEqual(original.tags, copy.tags)
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}
