package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
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
	fmt.Println("    dots-cli init -o <pack_name> 		Initializes new package in output directory")
	fmt.Println("    dots-cli add <app1> <app2> 		Adds given apps to package")
	fmt.Println("    dots-cli remote add <remote_url>	Adds new remote to package")
	fmt.Println("    dots-cli pack						Makes package version")
	fmt.Println("    To be continued...")
	return models.CommandResult{}
}
