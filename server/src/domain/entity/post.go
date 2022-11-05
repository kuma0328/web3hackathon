package entity

// Postはコミュニティ内の投稿に関する構造体です
type Post struct {
	Id          int
	CommunityId int
	Content     string
	Img         []byte
	Like        int
	UserId      int
	User        User
}

type Posts []*Post
