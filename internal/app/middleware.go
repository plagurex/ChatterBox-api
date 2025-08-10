package app

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func DBErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0].Err

			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(404, gin.H{"error": "Not Found"})
				c.Abort()
				return
			}

			panic(err)
		}
	}
}
