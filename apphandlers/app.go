package apphandlers

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
)

type App interface {
	GetPossibleDotfiles() []string
	LocateDotfiles() []string
	GetConfigRoot() string
	GetVersion() string
	AppExists() bool
	GetName() string
	NewApp() models.App
}

func GetExistingDotfiles(dotfiles []string) []string {
	existingDotfiles := make([]string, 0)
	for _, dotfile := range dotfiles {
		_, err := os.Stat(os.ExpandEnv(dotfile))
		if err == nil {
			existingDotfiles = append(existingDotfiles, dotfile)
		}
	}

	return existingDotfiles
}

func HandleApp(config *models.AppConfig, packageDir string, appName string) (bool, string) {
	// If app exists in handlers, get it
	var app *models.Handler
	if foundApp, ok := config.Handlers[appName]; ok {
		app = &foundApp
	} else {
		return false, fmt.Sprintf("Handler for '%s' does not exist. "+
			"\nPlease consider contributing at github.com/alvanrahimli/dots-cli-cli\n", appName)
	}

	locatedDotfiles := GetExistingDotfiles(app.Dotfiles)
	// If none found, return
	if len(locatedDotfiles) == 0 {
		return false, fmt.Sprintf("Could not find dotfiles for %s. Checked: %s\n",
			appName, strings.Join(app.Dotfiles, "\n\t"))
	}

	appSpecificDir := path.Join(packageDir, appName)
	var message string

	// Create app specific directory if it does not exist
	_, statErr := os.Stat(appSpecificDir)
	// If directory does not exist
	if statErr != nil {
		if os.IsNotExist(statErr) {
			mkdirErr := os.MkdirAll(appSpecificDir, os.ModePerm)
			if mkdirErr != nil {
				return false, fmt.Sprintf("Could not create app specific folder: %s", appName)
			}
		}
	}

	// Copy files. dotfile -> abs. path of file
	for _, dotfile := range locatedDotfiles {
		dotfile = os.ExpandEnv(dotfile)

		relativePath := strings.ReplaceAll(dotfile, os.ExpandEnv(app.ConfigRoot), "")
		if path.IsAbs(relativePath) {
			relativePath = relativePath[1:]
		}

		finalPathInPackage := path.Join(appSpecificDir, relativePath)
		finalDirInPackage := path.Dir(finalPathInPackage)
		// Create dotfile parent directory in package if it does not exist
		_, statErr := os.Stat(finalDirInPackage)
		if statErr != nil {
			mkdirErr := os.MkdirAll(finalDirInPackage, os.ModePerm)
			if mkdirErr != nil {
				message += fmt.Sprintf("Could not create folder: %s\n", finalDirInPackage)
				continue
			}
		}

		copyErr := utils.CopyFile(dotfile, finalPathInPackage)
		if copyErr != nil {
			message += fmt.Sprintf("ERROR: %s\n\tCould not copy '%s' to '%s'\n",
				copyErr.Error(), dotfile, relativePath)
			continue
		}
	}

	return message == "", message
}
