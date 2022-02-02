package sample_mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
)

// ブックマークを表すbson.D型のサンプルインスタンス。
func Bookmark() bson.D {
	bookmark := bson.D{
		{Key: "_id", Value: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6"},
		{Key: "name", Value: "example"},
		{Key: "uri", Value: "https://example.com"},
		{Key: "tags", Value: bson.A{"1", "2", "3"}},
	}
	return bookmark
}

// ブックマークを表すbson.D型のサンプルインスタンスA。
func BookmarkA() bson.D {
	bookmark := bson.D{
		{Key: "_id", Value: "f8ddce3a-0e87-4f3b-9f5d-148ba3125e42"},
		{Key: "name", Value: "example A"},
		{Key: "uri", Value: "https://example.com/foo"},
		{Key: "tags", Value: bson.A{}},
	}
	return bookmark
}

// ブックマークを表すbson.D型のサンプルインスタンスB。
func BookmarkB() bson.D {
	bookmark := bson.D{
		{Key: "_id", Value: "7a5c72ca-6e7d-4592-abb7-363ecac0d847"},
		{Key: "name", Value: "example B"},
		{Key: "uri", Value: "https://example.com/bar"},
		{Key: "tags", Value: bson.A{"B-1"}},
	}
	return bookmark
}

// ブックマークを表すbson.D型のサンプルインスタンスC。
func BookmarkC() bson.D {
	bookmark := bson.D{
		{Key: "_id", Value: "14bce21c-c9f1-43d2-a399-3e42954400f2"},
		{Key: "name", Value: "example C"},
		{Key: "uri", Value: "https://example.com/baz"},
		{Key: "tags", Value: bson.A{"C-1", "C-2"}},
	}
	return bookmark
}
