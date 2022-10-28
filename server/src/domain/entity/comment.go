package entity

// Commentはコミュニティ内のコメントに関する構造体です
type Comment struct {
	Id      int
	Content string
	Like    int
	User    User
}
