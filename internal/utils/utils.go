package utils

import (
	"chatterbox/internal/models"
	"encoding/json"
	"os"
)

func LoadConfig(path string) models.Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config models.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
