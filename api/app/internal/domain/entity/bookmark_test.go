package entity

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToID(t *testing.T, v string) *ID {
	t.Helper()
	id, err := NewID(v)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func ToName(t *testing.T, v string) *Name {
	t.Helper()
	name, err := NewName(v)
	if err != nil {
		t.Fatal(err)
	}
	return name
}

func ToURI(t *testing.T, v string) *URI {
	t.Helper()
	uri, err := NewURI(v)
	if err != nil {
		t.Fatal(err)
	}
	return uri
}

func ToTags(t *testing.T, vs ...string) []Tag {
	t.Helper()
	tags := make([]Tag, len(vs))
	for i, v := range vs {
		tag, err := NewTag(v)
		if err != nil {
			t.Fatal(err)
		}
		tags[i] = *tag
	}
	return tags
}

func TestNewBookmark(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	zeroTags := ToTags(t)
	oneTag := ToTags(t, "1-A")
	twoTags := ToTags(t, "1-A", "1-B")
	threeTags := ToTags(t, "1-A", "1-B", "1-C")
	cases := map[string]struct {
		id               *ID
		name             *Name
		uri              *URI
		tags             []Tag
		expectedBookmark *Bookmark
		expectedErr      error
	}{
		"non-nil arguments (0 tags)": {id, name, uri, zeroTags, &Bookmark{*id, *name, *uri, zeroTags}, nil},
		"non-nil arguments (1 tag)":  {id, name, uri, oneTag, &Bookmark{*id, *name, *uri, oneTag}, nil},
		"non-nil arguments (2 tags)": {id, name, uri, twoTags, &Bookmark{*id, *name, *uri, twoTags}, nil},
		"non-nil arguments (3 tags)": {id, name, uri, threeTags, &Bookmark{*id, *name, *uri, threeTags}, nil},
		"nil id":                     {nil, name, uri, threeTags, nil, errors.New("argument \"id\" is nil")},
		"nil name":                   {id, nil, uri, threeTags, nil, errors.New("argument \"name\" is nil")},
		"nil uri":                    {id, name, nil, threeTags, nil, errors.New("argument \"uri\" is nil")},
		"nil tags":                   {id, name, uri, nil, nil, errors.New("argument \"tags\" is nil")},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualBookmark, actualErr := NewBookmark(tc.id, tc.name, tc.uri, tc.tags)
			// then
			assert.Exactly(t, tc.expectedBookmark, actualBookmark)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
	{
		t.Run("tags pointer", func(t *testing.T) {
			t.Parallel()
			// given
			bookmark, _ := NewBookmark(id, name, uri, threeTags)
			x := bookmark.tags
			y := threeTags
			// when
			same := reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
			equiv := reflect.DeepEqual(x, y)
			// then
			assert.False(t, same)
			assert.True(t, equiv)
		})
	}
}

func TestBookmark_ID(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	// given
	bookmark, _ := NewBookmark(id, name, uri, tags)
	// when
	actualId := bookmark.ID()
	// then
	expectedId := *id
	assert.Exactly(t, expectedId, actualId)
}

func TestBookmark_Name(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	// given
	bookmark, _ := NewBookmark(id, name, uri, tags)
	// when
	actualName := bookmark.Name()
	// then
	expectedName := *name
	assert.Exactly(t, expectedName, actualName)
}

func TestBookmark_URI(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	// given
	bookmark, _ := NewBookmark(id, name, uri, tags)
	// when
	actualUri := bookmark.URI()
	// then
	expectedUri := *uri
	assert.Exactly(t, expectedUri, actualUri)
}

func TestBookmark_Tags(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	{
		t.Run("value", func(t *testing.T) {
			t.Parallel()
			// given
			bookmark, _ := NewBookmark(id, name, uri, tags)
			// when
			actualTags := bookmark.Tags()
			// then
			expectedTags := tags
			assert.ElementsMatch(t, expectedTags, actualTags)
		})
	}
	{
		t.Run("pointer", func(t *testing.T) {
			t.Parallel()
			// given
			bookmark, _ := NewBookmark(id, name, uri, tags)
			x := bookmark.Tags()
			y := bookmark.tags
			// when
			same := reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
			equiv := reflect.DeepEqual(x, y)
			// then
			assert.False(t, same)
			assert.True(t, equiv)
		})
	}
}

func TestBookmark_Rename(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	oldName := ToName(t, "Example")
	newName := ToName(t, "EXAMPLE")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	cases := map[string]struct {
		name         *Name
		expectedName Name
		expectedErr  error
	}{
		"non-nil name": {newName, *newName, nil},
		"nil name":     {nil, *oldName, errors.New("argument \"name\" is nil")},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			bookmark, _ := NewBookmark(id, oldName, uri, tags)
			// when
			actualErr := bookmark.Rename(tc.name)
			actualName := bookmark.name
			// then
			assert.Exactly(t, tc.expectedName, actualName)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_RewriteURI(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	oldUri := ToURI(t, "http://example.com")
	newUri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	cases := map[string]struct {
		uri         *URI
		expectedUri URI
		expectedErr error
	}{
		"non-nil uri": {newUri, *newUri, nil},
		"nil uri":     {nil, *oldUri, errors.New("argument \"uri\" is nil")},
	}
	for casename, tc := range cases {
		tc := tc
		t.Run(casename, func(t *testing.T) {
			t.Parallel()
			// given
			bookmark, _ := NewBookmark(id, name, oldUri, tags)
			// when
			actualErr := bookmark.RewriteURI(tc.uri)
			actualUri := bookmark.uri
			// then
			assert.Exactly(t, tc.expectedUri, actualUri)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_DeepCopy(t *testing.T) {
	t.Parallel()
	id := ToID(t, "00a0c91e6bf6")
	name := ToName(t, "Example")
	uri := ToURI(t, "https://example.com")
	tags := ToTags(t, "1-A", "1-B", "1-C")
	{
		t.Run("bookmark", func(t *testing.T) {
			t.Parallel()
			// given
			original, _ := NewBookmark(id, name, uri, tags)
			// when
			copy := original.DeepCopy()
			// then
			assert.Exactly(t, original, copy)
			assert.NotSame(t, original, copy)
		})
	}
	{
		t.Run("tags pointer", func(t *testing.T) {
			t.Parallel()
			// given
			original, _ := NewBookmark(id, name, uri, tags)
			copy := original.DeepCopy()
			x := copy.tags
			y := original.tags
			// when
			same := reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
			equiv := reflect.DeepEqual(x, y)
			// then
			assert.False(t, same)
			assert.True(t, equiv)
		})
	}
}
