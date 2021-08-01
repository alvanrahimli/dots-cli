package commands

import (
	"fmt"
	"github.com/dots/models"
	"github.com/dots/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Init command initializes new package in given output folder
type Init struct {
	Options *models.Opts
}

func (i Init) GetArguments() []string {
	return []string{}
}

func (i Init) CheckRequirements() (bool, string) {
	// If there is already manifest file here
	_, statErr := os.Stat(path.Join(i.Options.OutputDir, "manifest.json"))
	if statErr != nil {
		return true, ""
	}

	return false, "There is already package in this directory"
}

func (i Init) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	i.Options = opts
	satisfiesRequirements, message := i.CheckRequirements()
	if !satisfiesRequirements {
		fmt.Printf("Init command can not work in this directory:\n\t%s\n", message)
		os.Exit(1)
	}

	// If no value is given, use current directory
	if opts.OutputDir == "" {
		opts.OutputDir = "."
	} else {
		mkdirErr := os.MkdirAll(opts.OutputDir, os.ModePerm)
		if mkdirErr != nil {
			fmt.Printf("ERROR: %s\n", mkdirErr.Error())
			os.Exit(1)
		}
	}

	// Initialize manifest object
	manifest := models.NewManifest()

	// Package name
	if opts.PackageName == "" {
		var packageName string
		// Initialize default packageName
		if strings.Contains(opts.OutputDir, "/") {
			packageName = path.Base(opts.PackageName)
		} else {
			packageName = opts.OutputDir
		}

		fmt.Printf("Package name (%s): ", packageName)
		_, scanErr := fmt.Scanln(&packageName)
		if scanErr != nil {
			// As we initialized default value, we dont need to print error
			//fmt.Printf("ERROR: %s\n", scanErr.Error())
			//os.Exit(1)
		}

		opts.PackageName = packageName
	}
	manifest.Name = opts.PackageName

	// Author name
	if opts.AuthorName == "" {
		authorName := config.AuthorName
		fmt.Printf("Author name (%s): ", authorName)
		_, scanErr := fmt.Scanln(&authorName)
		if scanErr != nil {
			// As we initialized default value, we dont need to print error
			//fmt.Printf("ERROR: %s\n", scanErr.Error())
			//os.Exit(1)
		}

		opts.AuthorName = authorName
	}
	manifest.Author.Name = opts.AuthorName

	// Author email
	if opts.AuthorEmail == "" {
		authorEmail := config.AuthorEmail
		fmt.Printf("Author email (%s): ", authorEmail)
		_, scanErr := fmt.Scanln(&authorEmail)
		if scanErr != nil {
			// As we initialized default value, we dont need to print error
			//fmt.Printf("ERROR: %s\n", scanErr.Error())
			//os.Exit(1)
		}

		opts.AuthorEmail = authorEmail
	}
	manifest.Author.Email = opts.AuthorEmail

	// Package version
	if opts.Version == "" {
		packageVersion := models.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		}
		fmt.Printf("Package version (%s): ", packageVersion.ToString())
		_, scanErr := fmt.Scanf("%d.%d.%d",
			&packageVersion.Major,
			&packageVersion.Minor,
			&packageVersion.Patch)
		if scanErr != nil {
			// As we initialized default value, we dont need to print error
			//fmt.Printf("ERROR: %s\n", scanErr.Error())
			//os.Exit(1)
		}
		manifest.Versions = append(manifest.Versions, packageVersion)
	} else {
		manifest.Versions = append(manifest.Versions, utils.NewPackageVersion(opts.Version))
	}

	writeErr := utils.WriteManifestFile(opts.OutputDir, &manifest)
	if writeErr != nil {
		fmt.Printf("ERROR: %s\n", writeErr.Error())
		os.Exit(1)
	}

	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	finalPath := path.Join(filepath.Dir(ex), opts.OutputDir)
	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Initialized empty dots package in %s", finalPath),
	}
}
