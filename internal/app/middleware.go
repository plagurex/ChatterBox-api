package app

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0].Err

			switch {
			case errors.Is(err, sql.ErrNoRows):
				c.JSON(404, gin.H{"error": "Not Found"})
				c.Abort()
			case errs[0].Type == gin.ErrorTypeBind:
				c.JSON(400, gin.H{"error": "Bad Request", "message": err.Error()})
				c.Abort()
			default:
				c.JSON(501, gin.H{
					"message": "Internal Server Error",
				})
			}
		}
	}
}
