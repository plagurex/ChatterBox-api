package utils

import (
	"chatterbox/internal/models"
)

func LoadConfig(path string) models.Config {
	return models.Config{
		IsDebugMode: true,
		Addr:        "localhost:8080",
		DbPath:      "main.db",
	}
}
