package entity

// Communityはアプリ内のコミュニティに関する構造体です
type Community struct {
	Id      int
	Name    string
	ImgUrl  string
	Content string
	Recipe  Recipe
	Users   Users
}

type Communities []*Community
