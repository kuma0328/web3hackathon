package entitiy

// Commentはコミュニティ内のコメントに関する構造体です
type Comment struct {
	Id          int
	Content     string
	Like        int
	User        User
	Community   Community
	CommentType CommentType
}

// CommentTypeはComentの属性に関する情報を保持した構造体です
type CommentType struct {
	Id   int
	Type string
}
