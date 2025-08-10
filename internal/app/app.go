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

func (a App) Run(config models.Config) error {
	var err error
	a.db, err = sqlx.Open("sqlite3", config.DbPath)
	if err != nil {
		return fmt.Errorf("DB open failed: %w", err)
	}
	defer a.db.Close()
	if err := a.db.Ping(); err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
	}

	return a.r.Run(config.Addr)
}
