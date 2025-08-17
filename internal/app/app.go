package app

import (
	"chatterbox/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	r  *gin.Engine
	db *sqlx.DB
}

func NewApp() *App {
	return &App{r: gin.Default()}
}

func (a *App) Run(config models.Config) error {
	var err error
	a.db, err = sqlx.Open("sqlite3", config.DbPath)
	if err != nil {
		return fmt.Errorf("DB open failed: %w", err)
	}
	defer a.db.Close()
	if err := a.db.Ping(); err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
	}

	a.r.Use(ErrorMiddleware())

	a.r.GET("/posts", a.GetAllPostsHandler)
	a.r.GET("/posts/:post_id", a.GetPostHandler)
	a.r.GET("/posts/:post_id/comments", a.GetAllCommentsHandler)
	a.r.GET("/posts/:post_id/comments/:comment_id", a.GetCommentHandler)
	a.r.GET("/posts/:post_id/comments/:comment_id/replies", a.GetRepliesHandler)
	a.r.GET("/users", a.GetAllUsersHandler)
	a.r.GET("/users/:user_id/", a.GetUserHandler)

	return a.r.Run(config.Addr)
}

func handleGetOne[T any](a *App, c *gin.Context, query string, args ...interface{}) {
	var item T
	if err := a.db.Get(&item, query, args...); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, item)
}

func handleGetList[T any](a *App, c *gin.Context, query string, args ...interface{}) {
	items := []T{}
	if err := a.db.Select(&items, query, args...); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, items)

}
