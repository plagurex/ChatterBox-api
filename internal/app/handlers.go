package app

import (
	m "chatterbox/internal/models"
	"strconv"

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
	comments := make([]m.Comment, 0, 100)
	postId, _ := strconv.Atoi(c.Param("post_id"))
	if err := a.DfsComments(&comments, nil, params.PerPage*(params.Page+1), postId); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, comments[params.PerPage*params.Page:])
}

func (a *App) GetCommentHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ? and id = ?"
	handleGetOne[m.Comment](a, c, query, c.Param("post_id"), c.Param("comment_id"))
}

func (a *App) GetAllRepliesHandler(c *gin.Context) {
	query := "SELECT * FROM Comments WHERE post_id = ? and parent_id = ?"
	handleGetList[m.Comment](a, c, query, c.Param("post_id"), c.Param("comment_id"))
}

func (a *App) GetAllUsersHandler(c *gin.Context) {
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
	query := "SELECT id, username FROM Users LIMIT ? OFFSET ?"
	handleGetList[m.User](a, c, query, params.PerPage, params.Page*params.PerPage)
}

func (a *App) GetUserHandler(c *gin.Context) {
	query := "SELECT id, username FROM Users WHERE id = ?"
	handleGetOne[m.User](a, c, query, c.Param("user_id"))
}

func (a *App) AddUserHandler(c *gin.Context) {
	var item m.User
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	query := "INSERT INTO Users (username) VALUES (:username)"
	_, err := a.db.NamedExec(query, item)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, gin.H{
		"message": "created",
	})
}

func (a *App) AddPostHandler(c *gin.Context) {
	var item m.Post
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	query := "INSERT INTO Posts (username) VALUES (:username)"
	_, err := a.db.NamedExec(query, item)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, gin.H{
		"message": "created",
	})
}

func (a *App) AddCommentHandler(c *gin.Context) {
	var item m.User
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	query := "INSERT INTO Users (username) VALUES (:username)"
	_, err := a.db.NamedExec(query, item)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, gin.H{
		"message": "created",
	})
}
