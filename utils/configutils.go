package utils

import (
	"encoding/json"
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"os"
	"path"
)

func ReadConfig() (*models.AppConfig, error) {
	configPath, confDirErr := os.UserConfigDir()
	if confDirErr != nil {
		return nil, confDirErr
	}

	configPath = path.Join(configPath, "dots-cli", "config.json")
	configBytes, readErr := os.ReadFile(configPath)
	if readErr != nil {
		return nil, readErr
	}

	config := models.AppConfig{}
	jsonErr := json.Unmarshal(configBytes, &config)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if config.Handlers == nil {
		return nil, fmt.Errorf("no handler configurations found")
	}

	return &config, nil
}
