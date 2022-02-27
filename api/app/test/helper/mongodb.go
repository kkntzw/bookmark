package helper

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func ToBookmarkDocument(t *testing.T, id, name, uri string, tags ...string) bson.D {
	t.Helper()
	tagArray := bson.A{}
	for _, tag := range tags {
		tagArray = append(tagArray, tag)
	}
	doc := bson.D{
		{Key: "_id", Value: id},
		{Key: "name", Value: name},
		{Key: "uri", Value: uri},
		{Key: "tags", Value: tagArray},
	}
	return doc
}
