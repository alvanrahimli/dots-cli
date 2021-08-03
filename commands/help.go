package commands

import (
	"fmt"
	"github.com/dots/models"
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

func (h Help) ExecuteCommand(opts *models.Opts, _ *models.AppConfig) models.CommandResult {
	fmt.Println()
	fmt.Println("    dots init -o <pack_name> 		Initializes new package in output directory")
	fmt.Println("    dots add <app1> <app2> 		Adds given apps to package")
	fmt.Println("    dots remote add <remote_url>	Adds new remote to package")
	fmt.Println("    dots pack						Makes package version")
	fmt.Println("    To be continued...")
	return models.CommandResult{}
}
