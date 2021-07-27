package apphandlers

import (
	"dots/models"
	"dots/utils"
	"fmt"
	"os"
	"path"
	"strings"
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

func GetApps() map[string]App {
	return map[string]App{
		"i3": I3Wm{},
	}
}

func HandleApp(packageDir string, appName string) (bool, string) {
	apps := GetApps()
	app := apps[appName]
	allDotfiles := app.GetPossibleDotfiles()
	locatedDotfiles := app.LocateDotfiles()

	if len(locatedDotfiles) == 0 {
		return false, fmt.Sprintf("Could not find dotfiles for %s. Checked: %s\n",
			appName, strings.Join(allDotfiles, "\n\t"))
	}

	appSpecificDir := path.Join(packageDir, app.GetName())
	var message string

	// Create app specific directory if it does not exist
	_, statErr := os.Stat(appSpecificDir)
	// If directory does not exist
	if statErr != nil {
		if os.IsNotExist(statErr) {
			mkdirErr := os.MkdirAll(appSpecificDir, os.ModePerm)
			if mkdirErr != nil {
				return false, fmt.Sprintf("Could not create app specific folder: %s", app.GetName())
			}
		}
	}

	// Copy files
	for _, dotfile := range locatedDotfiles {
		dotfileDir := path.Dir(dotfile)
		_, statErr := os.Stat(dotfileDir)
		// Create dotfile parent directory if it does not exists
		if statErr != nil {
			mkdirErr := os.MkdirAll(dotfileDir, os.ModePerm)
			if mkdirErr != nil {
				message += fmt.Sprintf("Could not create folder: %s\n", dotfileDir)
				continue
			}
		}

		relativePath := strings.ReplaceAll(dotfile, app.GetConfigRoot(), "")
		if path.IsAbs(relativePath) {
			relativePath = relativePath[1:]
		}

		finalPathInPackage := path.Join(appSpecificDir, relativePath)
		copyErr := utils.CopyFile(dotfile, finalPathInPackage)
		if copyErr != nil {
			message += fmt.Sprintf("ERROR: %s\n\tCould not copy '%s' to '%s'\n",
				copyErr.Error(), dotfile, relativePath)
			continue
		}
	}

	return true, ""
}

func GetApp(appName string) models.App {
	apps := GetApps()
	return apps[appName].NewApp()
}
