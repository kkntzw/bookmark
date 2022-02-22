package helper

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/presentation/pb"
)

func ToBookmarkMessage(t *testing.T, id, name, uri string, tagNames ...string) *pb.Bookmark {
	t.Helper()
	tags := make([]*pb.Tag, len(tagNames))
	for i, tagName := range tagNames {
		tags[i] = &pb.Tag{TagName: tagName}
	}
	msg := &pb.Bookmark{
		BookmarkId:   id,
		BookmarkName: name,
		Uri:          uri,
		Tags:         tags,
	}
	return msg
}

func ToCreateBookmarkRequest(t *testing.T, name, uri string, tagNames ...string) *pb.CreateBookmarkRequest {
	t.Helper()
	tags := make([]*pb.Tag, len(tagNames))
	for i, tagName := range tagNames {
		tags[i] = &pb.Tag{TagName: tagName}
	}
	req := &pb.CreateBookmarkRequest{
		BookmarkName: name,
		Uri:          uri,
		Tags:         tags,
	}
	return req
}

func ToUpdateBookmarkRequest(t *testing.T, id, name, uri string) *pb.UpdateBookmarkRequest {
	t.Helper()
	req := &pb.UpdateBookmarkRequest{
		BookmarkId:   id,
		BookmarkName: name,
		Uri:          uri,
	}
	return req
}
