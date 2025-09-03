package models

import "database/sql"

type User struct {
	Id       int    `db:"id" json:"id" binding:"-"`
	Username string `db:"username" json:"username"`
}

type Post struct {
	Id     int    `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Text   string `db:"text" json:"text"`
	UserId int    `db:"user_id" json:"user_id"`
}

type Comment struct {
	Id       int           `db:"id" json:"id"`
	PostId   int           `db:"post_id" json:"post_id"`
	ParentId sql.NullInt64 `db:"parent_id" json:"parent_id"`
	UserId   int           `db:"user_id" json:"user_id"`
	Text     string        `db:"text" json:"text"`
}

type Config struct {
	DebugMode  bool   `json:"debug_mode"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DBHost     string `json:"db_host"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
}
