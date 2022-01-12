package command

// ブックマーク登録用のコマンド。
type RegisterBookmark struct {
	Name string   // ブックマーク名
	URI  string   // URI
	Tags []string // タグ一覧
}
