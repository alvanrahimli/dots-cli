package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
	"path"
)

type Remove struct {
	Options *models.Opts
}

func (r Remove) GetArguments() []string {
	return []string{}
}

func (r Remove) CheckRequirements() (bool, string) {
	if len(r.Options.Arguments) < 2 {
		return false, fmt.Sprintf("%s is not enough arguments for add command.", r.Options.Arguments)
	}

	return true, ""
}

func (r Remove) ExecuteCommand(opts *models.Opts, _ *models.AppConfig) models.CommandResult {
	r.Options = opts
	// Check if arguments satisfy required arguments for add command
	satisfiesRequirements, reqMessage := r.CheckRequirements()
	if !satisfiesRequirements {
		fmt.Printf("Add command can not work in this directory:\n\t%s\n", reqMessage)
		os.Exit(1)
	}

	// Read manifest
	manifest, manifestErr := utils.ReadManifestFile(opts.OutputDir)
	if manifestErr != nil {
		return models.CommandResult{
			Code: 1,
			Message: fmt.Sprintf("Could not read manifest file (%s)",
				manifestErr.Error()),
		}
	}

	// Check if app exists in manifest
	existingApps := make([]string, 0)
	for _, appName := range r.Options.Arguments[1:] {
		appExists := utils.AppExistsInManifest(appName, &manifest)
		if !appExists {
			fmt.Printf("App %s does exist in this package\n", appName)
			continue
		}

		existingApps = append(existingApps, appName)
	}

	if len(existingApps) == 0 {
		return models.CommandResult{
			Code:    1,
			Message: "Could not find any app in manifest",
		}
	}

	var message string
	// Iterate existing
	for _, app := range existingApps {
		// Check if app specific dir exists
		appSpecificDir := path.Join(opts.OutputDir, app)
		if _, scanErr := os.Stat(appSpecificDir); scanErr != nil {
			message += fmt.Sprintf("Could not find folder for '%s'\n", app)
			continue
		}

		// Remove from manifest

		modifiedApps, err := utils.RemoveApp(app, manifest.Apps)
		if err != nil {
			return models.CommandResult{
				Code:    1,
				Message: fmt.Sprintf("Could not remove app '%s' from manifest", app),
			}
		}

		// Update manifest apps
		manifest.Apps = modifiedApps

		// If exists, remove directory
		if removeErr := os.RemoveAll(appSpecificDir); removeErr != nil {
			message += fmt.Sprintf("Could not delete folder '%s'\n", appSpecificDir)
			continue
		}
	}

	// Removed successfully
	if len(message) == 0 {
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
			// Offer new version
			manifest.OfferNewVersion()
		}

		// Save new manifest
		saveErr := utils.WriteManifestFile(opts.OutputDir, &manifest)
		if saveErr != nil {
			return models.CommandResult{
				Code:    1,
				Message: fmt.Sprintf("Could not save manifest file (%s)", saveErr.Error()),
			}
		}
		return models.CommandResult{
			Code:    0,
			Message: "Successfully removed app from manifest",
		}
	} else {
		// Could not remove app
		return models.CommandResult{
			Code:    1,
			Message: message,
		}
	}
}
