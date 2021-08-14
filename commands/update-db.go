package commands

import (
	"encoding/json"
	"github.com/alvanrahimli/dots-cli/dlog"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
)

type UpdateDb struct {
	Options *models.Opts
}

func (u UpdateDb) GetArguments() []string {
	return []string{}
}

func (u UpdateDb) CheckRequirements() (bool, string) {
	return true, ""
}

func (u UpdateDb) ExecuteCommand(_ *models.Opts, config *models.AppConfig) models.CommandResult {
	response, err := utils.GetFromUrl(config.ConfigUrl)
	if err != nil {
		dlog.Err(err.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not get handlers",
		}
	}

	handlers := make(map[string]models.Handler, 0)
	unmarshallErr := json.Unmarshal(response, &handlers)
	if unmarshallErr != nil {
		dlog.Err(unmarshallErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not parse response data",
		}
	}

	config.Handlers = handlers
	configSaveErr := utils.SaveConfig(config)
	if configSaveErr != nil {
		dlog.Err(configSaveErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not save config file",
		}
	}

	dlog.Info("Database updated")
	return models.CommandResult{
		Code:    0,
		Message: "App database updated successfully",
	}
}
