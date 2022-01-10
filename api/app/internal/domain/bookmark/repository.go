package bookmark

// ブックマークの永続化を担うリポジトリのインターフェース。
type Repository interface {
	// IDを生成する。
	NextID() *ID

	// ブックマークを保存する。
	Save(bookmark *Bookmark) error
}
