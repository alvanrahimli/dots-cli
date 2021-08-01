package commands

import "github.com/dots/models"

// Remote command adds/removes registries for package in current folder
type Remote struct {
}

func (r Remote) GetArguments() []string {
	return []string{}
}

func (r Remote) CheckRequirements() (bool, string) {
	return true, ""
}

func (r Remote) ExecuteCommand(opts *models.Opts) models.CommandResult {
	return models.CommandResult{}
}
