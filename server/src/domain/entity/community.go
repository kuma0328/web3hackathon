package entitiy

// Communityはアプリ内のコミュニティに関する構造体です
type Community struct {
	Id     int
	Name   string
	ImgUrl string
	User   User
}
