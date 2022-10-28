package entity

// Postはコミュニティ内の投稿に関する構造体です
type Post struct {
	Id      int
	Content string
	ImgUrl  string
	Like    int
	User    User
}
