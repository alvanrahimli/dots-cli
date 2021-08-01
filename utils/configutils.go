package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dots/models"
	"os"
	"path"
)

func ReadConfig() *models.AppConfig {
	configPath, confDirErr := os.UserConfigDir()
	if confDirErr != nil {
		fmt.Printf("ERROR: %s\n", confDirErr.Error())
		os.Exit(1)
	}

	configPath = path.Join(configPath, "dots", "config.json")
	configBytes, readErr := os.ReadFile(configPath)
	if readErr != nil {
		if os.IsNotExist(readErr) {
			fmt.Println("No config file found.")
			// TODO: what?
			return &models.AppConfig{}
		} else {
			fmt.Printf("ERROR: %s\n", readErr.Error())
			os.Exit(1)
		}

	}

	config := models.AppConfig{}
	err := json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	if config.Handlers == nil {
		fmt.Println("No handler configurations found. Quitting.")
		os.Exit(1)
	}

	return &config
}
