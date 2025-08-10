package app

import (
	m "chatterbox/internal/models"

	"github.com/gin-gonic/gin"
)

func (a *App) GetAllPostsHandler(c *gin.Context) {
	query := "SELECT * FROM Posts"
	handleGetList[m.Post](a, c, query)
}

func (a *App) GetPostHandler(c *gin.Context) {
	query := "SELECT * FROM Posts WHERE id = ?"
	handleGetOne[m.Post](a, c, query, c.Param("post_id"))
}

func (a *App) GetAllCommentsHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ? and parent_id IS NULL"
	handleGetList[m.Comment](a, c, query, c.Param("post_id"))
}

func (a *App) GetCommentHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ? and id = ?"
	handleGetOne[m.Comment](a, c, query, c.Param("post_id"), c.Param("comment_id"))
}

func (a *App) GetRepliesHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ? and parent_id = ?"
	handleGetList[m.Comment](a, c, query, c.Param("post_id"), c.Param("comment_id"))
}

func (a *App) GetAllUsersHandler(c *gin.Context) {
	query := "SELECT id, username FROM Users"
	handleGetList[m.User](a, c, query)
}

func (a *App) GetUserHandler(c *gin.Context) {
	query := "SELECT id, username FROM Users WHERE id = ?"
	handleGetOne[m.User](a, c, query, c.Param("user_id"))
}
