package entitiy

// Userはアプリに登録されたユーザーに関する構造体です
type User struct {
	Id       int
	Name     string
	Mail     string
	Password string
	Communities Communities
}

type Users []User