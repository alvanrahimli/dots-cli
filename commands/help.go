package commands

import (
	"dots/models"
	"fmt"
)

// Help command prints help message
type Help struct {
}

func (h Help) GetArguments() []string {
	return []string{}
}

func (h Help) CheckRequirements() (bool, string) {
	return true, ""
}

func (h Help) ExecuteCommand(opts *models.Opts) models.CommandResult {
	fmt.Println()
	fmt.Println("   Yes, this is help message")
	return models.CommandResult{}
}
