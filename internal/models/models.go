package models

type User struct {
	Id       int
	Username string
}

type Post struct {
	Id     int
	Title  string
	Text   string
	UserId int `db:"user_id"`
}

type Comment struct {
	Id       int
	PostId   int `db:"post_id"`
	ParentId int `db:"paresnt_id"`
	UserId   int `db:"user_id"`
}

type Config struct {
	IsDebugMode bool
	Addr        string
	DbPath      string
}
