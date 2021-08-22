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

func (h Help) ExecuteCommand(_ *models.Opts, _ *models.AppConfig) models.CommandResult {
	fmt.Println()
	fmt.Println("    dots-cli init <pack_name> 	        Initializes new package in")
	fmt.Println("    dots-cli add <app1> <app2>         Adds given apps to package")
	fmt.Println("    dots-cli add -w <image_path>       Adds specified wallpaper to package")
	fmt.Println("    dots-cli remote add <remote_url>   Adds new remote to package")
	fmt.Println("    dots-cli pack                      Makes package version")
	fmt.Println("    dots-cli push <remote_name>        Pushes package to specified remote")
	fmt.Println("    dots-cli get <package_name>        Downloads & Extracts package from registry")
	fmt.Println("    dots-cli install                   Installs downloaded package")
	fmt.Println("    dots-cli uninstall                 Uninstalls package")
	fmt.Println("    dots-cli update-db                 Updates handler database")
	fmt.Println()

	return models.CommandResult{}
}
