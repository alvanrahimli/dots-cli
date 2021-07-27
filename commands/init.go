package commands

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Init command initializes new package in given output folder
type Init struct {
	Options *Opts
}

func (i Init) getArguments() []string {
	return []string{}
}

func (i Init) checkRequirements() (bool, string) {
	// If there is already manifest file here
	_, statErr := os.Stat(path.Join(i.Options.OutputDir, "manifest.json"))
	if statErr != nil {
		return true, ""
	}

	return false, "There is already package in this directory"
}

func (i Init) ExecuteCommand(opts *Opts) CommandResult {
	i.Options = opts
	satisfiesRequirements, message := i.checkRequirements()
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
	manifest := NewManifest()

	// Package name
	if opts.PackageName == "" {
		var packageName string
		fmt.Print("Package name: ")
		_, scanErr := fmt.Scanln(&packageName)
		if scanErr != nil {
			fmt.Printf("ERROR: %s\n", scanErr.Error())
			os.Exit(1)
		}

		opts.PackageName = packageName
	}
	manifest.Name = opts.PackageName

	// Author name
	if opts.AuthorName == "" {
		var authorName string
		fmt.Print("Author name: ")
		_, scanErr := fmt.Scanln(&authorName)
		if scanErr != nil {
			fmt.Printf("ERROR: %s\n", scanErr.Error())
			os.Exit(1)
		}

		opts.AuthorName = authorName
	}
	manifest.Author.Name = opts.AuthorName

	// Author email
	if opts.AuthorEmail == "" {
		var authorEmail string
		fmt.Print("Author email: ")
		_, scanErr := fmt.Scanln(&authorEmail)
		if scanErr != nil {
			fmt.Printf("ERROR: %s\n", scanErr.Error())
			os.Exit(1)
		}

		opts.AuthorEmail = authorEmail
	}
	manifest.Author.Email = opts.AuthorEmail

	// Package version
	if opts.Version == "" {
		var packageVersion Version
		fmt.Print("Package version: ")
		_, scanErr := fmt.Scanf("%d.%d.%d",
			&packageVersion.Major,
			&packageVersion.Minor,
			&packageVersion.Patch)
		if scanErr != nil {
			fmt.Printf("ERROR: %s\n", scanErr.Error())
			os.Exit(1)
		}
		manifest.Version = packageVersion
	} else {
		manifest.Version = NewPackageVersion(opts.Version)
	}

	writeErr := WriteManifestFile(opts.OutputDir, &manifest)
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
	return CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Initialized empty dots package in %s", finalPath),
	}
}
