package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/apphandler"
	"github.com/alvanrahimli/dots-cli/dlog"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
	"path"
	"strings"
)

// Add command when name is specified adds new application to package.
type Add struct {
	Options *models.Opts
}

func (a Add) GetArguments() []string {
	return []string{}
}

func (a Add) CheckRequirements() (bool, string) {
	if len(a.Options.Arguments) < 2 && a.Options.WpPath == "" {
		return false, fmt.Sprintf("%s is not enough arguments for add command.", a.Options.Arguments)
	}

	return true, ""
}

func (a Add) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	a.Options = opts
	// Check if arguments satisfy required arguments for add command
	satisfiesRequirements, message := a.CheckRequirements()
	if !satisfiesRequirements {
		fmt.Printf("Add command can not work in this directory:\n\t%s\n", message)
		os.Exit(1)
	}

	// Read manifest
	manifest, manifestErr := utils.ReadManifestFile(opts.OutputDir)
	if manifestErr != nil {
		fmt.Printf("ERROR: %s\n", manifestErr.Error())
		os.Exit(1)
	}

	// If user adds wallpaper
	if a.Options.WpPath != "" {
		wallpapersDir := path.Join(a.Options.OutputDir, "wallpapers")
		// Check for wallpapers directory
		_, statErr := os.Stat(wallpapersDir)
		if statErr != nil {
			mkdirErr := os.Mkdir(wallpapersDir, os.ModePerm)
			if mkdirErr != nil {
				dlog.Err(mkdirErr.Error())
				return models.CommandResult{
					Code:    1,
					Message: "Could not create wallpapers folder",
				}
			}
		}

		wpName := path.Join(wallpapersDir, path.Base(a.Options.WpPath))
		copyErr := utils.CopyFile(a.Options.WpPath, wpName)
		if copyErr != nil {
			dlog.Err(copyErr.Error())
			fmt.Println("Could not copy wallpaper")
		} else {
			fmt.Println("Wallpaper added")
			relativePath := path.Join("wallpapers", path.Base(a.Options.WpPath))
			manifest.Wallpapers = append(manifest.Wallpapers, relativePath)
		}
	}

	// Check if apps exist in package or not
	possibleAppNames := make([]string, 0)
	for _, appName := range opts.Arguments[1:] {
		appExists := utils.AppExistsInManifest(appName, &manifest)
		if appExists {
			fmt.Printf("App %s already exists in this package\n", appName)
			continue
		}

		possibleAppNames = append(possibleAppNames, appName)
	}

	// Exit app if all apps are in package
	if len(possibleAppNames) == 0 && a.Options.WpPath == "" {
		fmt.Println("All packages are already in package")
		os.Exit(1)
	}

	addedApps := make([]string, 0)
	failedApps := make([]string, 0)
	// Copy files to package
	for _, appName := range possibleAppNames {
		added, message := apphandler.HandleApp(config, opts.OutputDir, appName)
		if added {
			addedApps = append(addedApps, appName)
			manifest.Apps = append(manifest.Apps, models.App{
				Name:    appName,
				Version: config.Handlers[appName].Version,
			})
		} else {
			fmt.Printf("%s\n", message)
			failedApps = append(failedApps, appName)
		}
	}

	// If there are new apps
	if len(addedApps) > 0 || a.Options.WpPath != "" {
		// Remove old manifest
		manifestPath := path.Join(opts.OutputDir, "manifest.json")
		removeErr := os.Remove(manifestPath)
		if removeErr != nil {
			return models.CommandResult{
				Code:    1,
				Message: "Could not remove old manifest file",
			}
		}

		if !manifest.Modified {
			manifest.Modified = true
			manifest.OfferNewVersion()
		}

		// Save manifest file
		manifestWriteErr := utils.WriteManifestFile(opts.OutputDir, &manifest)
		if manifestWriteErr != nil {
			return models.CommandResult{
				Code:    1,
				Message: "Could write updated manifest file",
			}
		}

		var resultMsg string
		if len(addedApps) > 0 {
			resultMsg += fmt.Sprintf("Following apps added to package: %s\n", strings.Join(addedApps, ", "))
		}
		if len(failedApps) > 0 {
			resultMsg += fmt.Sprintf("These apps could not be added: %s\n", strings.Join(failedApps, ", "))
		}

		return models.CommandResult{
			Code:    0,
			Message: resultMsg,
		}
	} else {
		return models.CommandResult{
			Code: 1,
			Message: fmt.Sprintf("\tThese apps could not be added: %s\n",
				strings.Join(failedApps, ", ")),
		}
	}
}
