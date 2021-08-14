package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
	"path"
	"strings"
)

type Uninstall struct {
	Options *models.Opts
}

func (u Uninstall) GetArguments() []string {
	return []string{}
}

func (u Uninstall) CheckRequirements() (bool, string) {
	return true, ""
}

func (u Uninstall) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	u.Options = opts
	// Check if arguments satisfy required arguments for add command
	satisfiesRequirements, requirementMessage := u.CheckRequirements()
	if !satisfiesRequirements {
		fmt.Printf("Add command can not work in this directory:\n\t%s\n", requirementMessage)
		os.Exit(1)
	}

	// Read manifest
	manifest, manifestErr := utils.ReadManifestFile(opts.OutputDir)
	if manifestErr != nil {
		fmt.Printf("ERROR: %s\n", manifestErr.Error())
		os.Exit(1)
	}

	// We can add isInstalled in manifest in future
	_, statErr := os.Stat(path.Join(u.Options.OutputDir, ".backup"))
	if statErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: ".backup folder not found to restore",
		}
	}

	uninstalledApps := make([]string, 0)
	failedDotfiles := make([]string, 0)
	// Copy files from backup to ConfigRoot/file_path
	for _, app := range manifest.Apps {
		var handler models.Handler
		var handlerFound = false
		for hName, h := range config.Handlers {
			if hName == app.Name {
				handler = h
				handlerFound = true
				break
			}
		}

		if !handlerFound {
			fmt.Printf("Could not find handler for %s", app.Name)
			continue
		}

		for _, originalDotfile := range handler.Dotfiles {
			originalDotfile = os.ExpandEnv(originalDotfile)

			// Check if .backup folder exists
			backupRoot := path.Join(u.Options.OutputDir, ".backup")
			backupRel := path.Join(backupRoot, app.Name, strings.ReplaceAll(originalDotfile, os.ExpandEnv(handler.ConfigRoot), ""))
			if path.IsAbs(backupRel) {
				backupRel = backupRel[1:]
			}

			// Check if dotfile exists
			_, statErr = os.Stat(originalDotfile)
			if statErr != nil {
				if os.IsNotExist(statErr) {
					fmt.Printf("'%s' not found\n", originalDotfile)
				} else {
					fmt.Printf("Error: '%s' %s\n", originalDotfile, statErr.Error())
				}
				continue
			}

			// Copy file
			copyErr := utils.CopyFile(backupRel, originalDotfile)
			if copyErr != nil {
				fmt.Printf("Could not backup file: %s. (Err: %s)\n", originalDotfile, copyErr.Error())
				fmt.Printf("You can manually restore file from: %s\n", backupRel)
				failedDotfiles = append(failedDotfiles, originalDotfile)
			}
		}

		uninstalledApps = append(uninstalledApps, app.Name)
	}

	var message string
	if len(uninstalledApps) > 0 {
		message = fmt.Sprintf("These apps have been uninstalled and previous states are restored: %s\n",
			strings.Join(uninstalledApps, ", "))
	}
	if len(failedDotfiles) > 0 {
		message += fmt.Sprintf("Could not restore following dotfiles: \n\t%s",
			strings.Join(failedDotfiles, "\n\t"))
	}

	return models.CommandResult{
		Code:    0,
		Message: message,
	}
}
