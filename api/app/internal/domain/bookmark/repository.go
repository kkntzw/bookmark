package bookmark

// ブックマークの永続化を担うリポジトリのインターフェース。
type Repository interface {
	// IDを生成する。
	NextID() *ID

	// IDからブックマークを検索する。
	//
	// 該当するブックマークが存在しない場合はnilを返却する。
	FindByID(id *ID) (*Bookmark, error)

	// ブックマークを保存する。
	Save(bookmark *Bookmark) error
}
