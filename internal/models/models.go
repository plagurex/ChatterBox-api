package models

import "database/sql"

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
	PostId   int           `db:"post_id"`
	ParentId sql.NullInt64 `db:"parent_id"`
	UserId   int           `db:"user_id"`
	Text     string
}

type Config struct {
	IsDebugMode bool
	Addr        string
	DbPath      string
}
