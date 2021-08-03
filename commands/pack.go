package commands

import (
	"fmt"
	"github.com/dots/models"
	"github.com/dots/utils"
	"os"
	"path"
)

type Pack struct {
	Options *models.Opts
}

func (p Pack) GetArguments() []string {
	return []string{}
}

func (p Pack) CheckRequirements() (bool, string) {
	if len(p.Options.Arguments) < 1 {
		return false, fmt.Sprintf("%s is not enough arguments for add command.", p.Options.Arguments)
	}

	return true, ""
}

func (p Pack) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	p.Options = opts
	// Check if arguments satisfy required arguments for add command
	satisfiesRequirements, message := p.CheckRequirements()
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

	if !manifest.Modified {
		return models.CommandResult{
			Code:    0,
			Message: "Package is not modified, there is no need to pack again.",
		}
	}

	versFolderPath := path.Join(p.Options.OutputDir, ".vers")

	// Check if .vers directory exists, if not, create it
	if _, statErr := os.Stat(versFolderPath); statErr != nil {
		mkdirErr := os.Mkdir(versFolderPath, os.ModePerm)
		if mkdirErr != nil {
			return models.CommandResult{
				Code:    1,
				Message: fmt.Sprintf("Could not create .vers folder. (%s)", mkdirErr.Error()),
			}
		}
	}

	tarFileName := fmt.Sprintf("%s-%s-dots.tar.gz", manifest.Name, manifest.LastVersion().ToFormattedString())
	tarballPath := path.Join(versFolderPath, tarFileName)
	tarballErr := utils.CreateTarball(opts.OutputDir, tarballPath)
	if tarballErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not create tarball",
		}
	}

	manifest.Modified = false

	manifestWriteErr := utils.WriteManifestFile(opts.OutputDir, &manifest)
	if manifestWriteErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could write updated manifest file",
		}
	}

	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Created package at %s", tarballPath),
	}
}
