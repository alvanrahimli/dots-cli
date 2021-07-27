package commands

import "dots/models"

// Push command pushes package to already added registry
// Registries can be added using `dots remote add` command
type Push struct {
}

func (p Push) GetArguments() []string {
	return []string{}
}

func (p Push) CheckRequirements() (bool, string) {
	return true, ""
}

func (p Push) ExecuteCommand(opts *models.Opts) models.CommandResult {
	return models.CommandResult{}
}
