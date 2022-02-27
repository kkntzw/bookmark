package helper

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
)

func ToID(t *testing.T, v string) *entity.ID {
	t.Helper()
	id, err := entity.NewID(v)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func ToName(t *testing.T, v string) *entity.Name {
	t.Helper()
	name, err := entity.NewName(v)
	if err != nil {
		t.Fatal(err)
	}
	return name
}

func ToURI(t *testing.T, v string) *entity.URI {
	t.Helper()
	uri, err := entity.NewURI(v)
	if err != nil {
		t.Fatal(err)
	}
	return uri
}

func ToTags(t *testing.T, vs ...string) []entity.Tag {
	t.Helper()
	tags := make([]entity.Tag, len(vs))
	for i, v := range vs {
		tag, err := entity.NewTag(v)
		if err != nil {
			t.Fatal(err)
		}
		tags[i] = *tag
	}
	return tags
}

func ToErrID(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewID(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrName(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewName(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrURI(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewURI(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrTag(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewTag(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToBookmark(t *testing.T, iv, nv, uv string, tvs ...string) *entity.Bookmark {
	t.Helper()
	id := ToID(t, iv)
	name := ToName(t, nv)
	uri := ToURI(t, uv)
	tags := ToTags(t, tvs...)
	bookmark, err := entity.NewBookmark(id, name, uri, tags)
	if err != nil {
		t.Fatal(err)
	}
	return bookmark
}
