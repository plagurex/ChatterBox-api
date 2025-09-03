package app

import (
	"chatterbox/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type App struct {
	r  *gin.Engine
	db *sqlx.DB
}

func NewApp() *App {
	return &App{r: gin.Default()}
}

func (a *App) Run(config models.Config) error {
	fmt.Printf("Starting on: http://%s:%d\n", config.Host, config.Port)

	var err error
	a.db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName))
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
	a.r.GET("/posts/:post_id/comments/:comment_id/replies", a.GetAllRepliesHandler)
	a.r.GET("/users", a.GetAllUsersHandler)
	a.r.GET("/users/:user_id/", a.GetUserHandler)

	a.r.POST("/users", a.AddUserHandler)
	return a.r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
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

func (a *App) DfsComments(slice *[]models.Comment, base *models.Comment, limit int, postId int) error {
	replies := []models.Comment{}
	var err error
	if base == nil {
		err = a.db.Select(&replies, "SELECT * FROM Comments WHERE post_id = ? and parent_id IS NULL", postId)
	} else {
		err = a.db.Select(&replies, "SELECT * FROM Comments WHERE parent_id = ?", base.Id)
	}
	if err != nil {
		return err
	}
	for _, v := range replies {
		if len(*slice) >= limit {
			break
		}
		*slice = append(*slice, v)
		a.DfsComments(slice, &v, limit, postId)
	}

	return nil
}
