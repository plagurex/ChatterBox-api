package app

import (
	m "chatterbox/internal/models"

	"github.com/gin-gonic/gin"
)

func (a *App) GetAllPostsHandler(c *gin.Context) {
	params := struct {
		PerPage int `form:"per_page" binding:"omitempty,max=100,min=1"`
		Page    int `form:"page" binding:"omitempty,min=0"`
	}{
		PerPage: 20,
		Page:    0,
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	query := "SELECT * FROM Posts LIMIT ? OFFSET ?"
	handleGetList[m.Post](a, c, query, params.PerPage, params.Page*params.PerPage)
}

func (a *App) GetPostHandler(c *gin.Context) {
	query := "SELECT * FROM Posts WHERE id = ?"
	handleGetOne[m.Post](a, c, query, c.Param("post_id"))
}

func (a *App) GetAllCommentsHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ?"
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
