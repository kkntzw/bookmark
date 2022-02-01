package sample_pb

import (
	"github.com/kkntzw/bookmark/internal/presentation/pb"
)

// pb.Bookmark型のサンプルインスタンス。
func Bookmark() *pb.Bookmark {
	return &pb.Bookmark{
		BookmarkId:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		BookmarkName: "example",
		Uri:          "https://example.com",
		Tags: []*pb.Tag{
			{TagName: "1"},
			{TagName: "2"},
			{TagName: "3"},
		},
	}
}

// pb.Bookmark型のサンプルインスタンスA。
func BookmarkA() *pb.Bookmark {
	return &pb.Bookmark{
		BookmarkId:   "f8ddce3a-0e87-4f3b-9f5d-148ba3125e42",
		BookmarkName: "example A",
		Uri:          "https://example.com/foo",
		Tags:         []*pb.Tag{},
	}
}

// pb.Bookmark型のサンプルインスタンスB。
func BookmarkB() *pb.Bookmark {
	return &pb.Bookmark{
		BookmarkId:   "7a5c72ca-6e7d-4592-abb7-363ecac0d847",
		BookmarkName: "example B",
		Uri:          "https://example.com/bar",
		Tags: []*pb.Tag{
			{TagName: "B-1"},
		},
	}
}

// pb.Bookmark型のサンプルインスタンスC。
func BookmarkC() *pb.Bookmark {
	return &pb.Bookmark{
		BookmarkId:   "14bce21c-c9f1-43d2-a399-3e42954400f2",
		BookmarkName: "example C",
		Uri:          "https://example.com/baz",
		Tags: []*pb.Tag{
			{TagName: "C-1"},
			{TagName: "C-2"},
		},
	}
}
