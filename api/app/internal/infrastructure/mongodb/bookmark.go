package mongodb

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ブックマークの永続化を担うリポジトリの具象型。
type bookmarkRepository struct {
	collection *mongo.Collection // コレクション
}

// ブックマークの永続化を担うリポジトリを生成する。
func NewBookmarkRepository(collection *mongo.Collection) repository.Bookmark {
	return &bookmarkRepository{
		collection: collection,
	}
}

// ブックマークに関するドキュメント。
type BookmarkDocument struct {
	ID   string   `bson:"_id"`  // ID
	Name string   `bson:"name"` // ブックマーク名
	URI  string   `bson:"uri"`  // URI
	Tags []string `bson:"tags"` // タグ一覧
}

// IDを生成する。
//
// バージョン4のUUIDを16進表記で生成する。
func (r *bookmarkRepository) NextID() *entity.ID {
	uuid, _ := uuid.NewRandom()
	id, _ := entity.NewID(uuid.String())
	return id
}

// ブックマークを保存する。
//
// nilを指定した場合はエラーを返却する。
// ドキュメントの保存に失敗した場合はエラーを返却する。
//
//	db.bookmarks.updateOne(
//	  {_id: "ID"},
//	  {
//	    $set: {_id: "ID", name: "Name", tags: ["1", "2", "3"], uri: "URI"},
//	    $currentDate: {lastModified: true}
//	  },
//	  {upsert: true}
//	)
func (r *bookmarkRepository) Save(bookmark *entity.Bookmark) error {
	if bookmark == nil {
		return fmt.Errorf("argument \"bookmark\" is nil")
	}
	ctx := context.Background()
	id := bookmark.ID()
	name := bookmark.Name()
	uri := bookmark.URI()
	tags := make([]string, len(bookmark.Tags()))
	for i, tag := range bookmark.Tags() {
		tags[i] = tag.Value()
	}
	document := BookmarkDocument{
		ID:   id.Value(),
		Name: name.Value(),
		URI:  uri.String(),
		Tags: tags,
	}
	update := bson.M{"$set": document, "$currentDate": bson.M{"lastModified": true}}
	opts := options.Update().SetUpsert(true)
	if _, err := r.collection.UpdateByID(ctx, id.Value(), update, opts); err != nil {
		return fmt.Errorf("failed at collection.UpdateByID: %w", err)
	}
	return nil
}

// ブックマーク一覧を検索する。
//
// ブックマークが存在しない場合は空のスライスを返却する。
//
// ドキュメントの検索に失敗した場合はエラーを返却する。
// ドキュメントのデコードに失敗した場合はエラーを返却する。
//
//	db.bookmarks.find({})
func (r *bookmarkRepository) FindAll() ([]entity.Bookmark, error) {
	ctx := context.Background()
	filter := bson.D{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed at collection.Find: %w", err)
	}
	var documents []BookmarkDocument
	if err := cursor.All(ctx, &documents); err != nil {
		return nil, fmt.Errorf("failed at cursor.All: %w", err)
	}
	bookmarks := make([]entity.Bookmark, len(documents))
	for i, document := range documents {
		id, _ := entity.NewID(document.ID)
		name, _ := entity.NewName(document.Name)
		uri, _ := entity.NewURI(document.URI)
		tags := make([]entity.Tag, len(document.Tags))
		for i, v := range document.Tags {
			tag,  _ := entity.NewTag(v)
			tags[i] = *tag
		}
		bookmark, _ := entity.NewBookmark(id, name, uri, tags)
		bookmarks[i] = *bookmark
	}
	return bookmarks, nil
}

// IDからブックマークを検索する。
//
// 該当するブックマークが存在しない場合はnilを返却する。
//
// nilを指定した場合はエラーを返却する。
// ドキュメントの検索に失敗した場合はエラーを返却する。
//
//	db.bookmarks.findOne({_id: "ID"})
func (r *bookmarkRepository) FindByID(id *entity.ID) (*entity.Bookmark, error) {
	if id == nil {
		return nil, fmt.Errorf("argument \"id\" is nil")
	}
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: id.Value()}}
	result := r.collection.FindOne(ctx, filter)
	var document BookmarkDocument
	err := result.Decode(&document)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed at collection.FindOne: %w", err)
	}
	name, _ := entity.NewName(document.Name)
	uri, _ := entity.NewURI(document.URI)
	tags := make([]entity.Tag, len(document.Tags))
	for i, v := range document.Tags {
		tag, _ := entity.NewTag(v)
		tags[i] = *tag
	}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark, nil
}
