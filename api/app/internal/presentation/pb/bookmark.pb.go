// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: bookmark.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ブックマークを表すメッセージ。
type Bookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ブックマークIDを表すフィールド。
	BookmarkId string `protobuf:"bytes,1,opt,name=bookmark_id,json=bookmarkId,proto3" json:"bookmark_id,omitempty"`
	// ブックマーク名を表すフィールド。
	BookmarkName string `protobuf:"bytes,2,opt,name=bookmark_name,json=bookmarkName,proto3" json:"bookmark_name,omitempty"`
	// URIを表すフィールド。
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// タグ一覧を表すフィールド。
	Tags []*Tag `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Bookmark) Reset() {
	*x = Bookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookmark) ProtoMessage() {}

func (x *Bookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookmark.ProtoReflect.Descriptor instead.
func (*Bookmark) Descriptor() ([]byte, []int) {
	return file_bookmark_proto_rawDescGZIP(), []int{0}
}

func (x *Bookmark) GetBookmarkId() string {
	if x != nil {
		return x.BookmarkId
	}
	return ""
}

func (x *Bookmark) GetBookmarkName() string {
	if x != nil {
		return x.BookmarkName
	}
	return ""
}

func (x *Bookmark) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Bookmark) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// タグを表すメッセージ。
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// タグ名を表すフィールド。
	//
	// 必須項目。
	// 空白は不正とする。
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_bookmark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_bookmark_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

// CreateBookmark 用のリクエストメッセージ。
type CreateBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ブックマーク名を表すフィールド。
	//
	// 必須項目。
	// 空白は不正とする。
	BookmarkName string `protobuf:"bytes,1,opt,name=bookmark_name,json=bookmarkName,proto3" json:"bookmark_name,omitempty"`
	// URIを表すフィールド。
	//
	// 必須項目。
	// 空白は不正とする。
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// タグ一覧を表すフィールド。
	Tags []*Tag `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateBookmarkRequest) Reset() {
	*x = CreateBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookmarkRequest) ProtoMessage() {}

func (x *CreateBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookmark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookmarkRequest.ProtoReflect.Descriptor instead.
func (*CreateBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_bookmark_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookmarkRequest) GetBookmarkName() string {
	if x != nil {
		return x.BookmarkName
	}
	return ""
}

func (x *CreateBookmarkRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *CreateBookmarkRequest) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_bookmark_proto protoreflect.FileDescriptor

var file_bookmark_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x21, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x20, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x71, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x32, 0x57, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x49, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookmark_proto_rawDescOnce sync.Once
	file_bookmark_proto_rawDescData = file_bookmark_proto_rawDesc
)

func file_bookmark_proto_rawDescGZIP() []byte {
	file_bookmark_proto_rawDescOnce.Do(func() {
		file_bookmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookmark_proto_rawDescData)
	})
	return file_bookmark_proto_rawDescData
}

var file_bookmark_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bookmark_proto_goTypes = []interface{}{
	(*Bookmark)(nil),              // 0: bookmark.Bookmark
	(*Tag)(nil),                   // 1: bookmark.Tag
	(*CreateBookmarkRequest)(nil), // 2: bookmark.CreateBookmarkRequest
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_bookmark_proto_depIdxs = []int32{
	1, // 0: bookmark.Bookmark.tags:type_name -> bookmark.Tag
	1, // 1: bookmark.CreateBookmarkRequest.tags:type_name -> bookmark.Tag
	2, // 2: bookmark.Bookmarker.CreateBookmark:input_type -> bookmark.CreateBookmarkRequest
	3, // 3: bookmark.Bookmarker.CreateBookmark:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bookmark_proto_init() }
func file_bookmark_proto_init() {
	if File_bookmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookmark); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookmark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookmark_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookmarkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookmark_proto_goTypes,
		DependencyIndexes: file_bookmark_proto_depIdxs,
		MessageInfos:      file_bookmark_proto_msgTypes,
	}.Build()
	File_bookmark_proto = out.File
	file_bookmark_proto_rawDesc = nil
	file_bookmark_proto_goTypes = nil
	file_bookmark_proto_depIdxs = nil
}