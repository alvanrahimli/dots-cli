package apphandlers

import (
	"github.com/alvanrahimli/dots-cli/models"
	"os"
	"os/exec"
)

type I3Wm struct {
}

func (i3 I3Wm) GetPossibleDotfiles() []string {
	return []string{
		os.ExpandEnv("$HOME/.config/i3/config"),
		os.ExpandEnv("$HOME/.i3/config"),
	}
}

func (i3 I3Wm) LocateDotfiles() []string {
	existingDotfiles := make([]string, 0)

	dotfiles := i3.GetPossibleDotfiles()
	for _, dotfile := range dotfiles {
		_, err := os.Stat(dotfile)
		if err == nil {
			existingDotfiles = append(existingDotfiles, dotfile)
		}
	}

	return existingDotfiles
}

func (i3 I3Wm) GetVersion() string {
	return "1.0.3"
}

func (i3 I3Wm) GetName() string {
	return "i3"
}

func (i3 I3Wm) NewApp() models.App {
	return models.App{
		Name:    i3.GetName(),
		Version: i3.GetVersion(),
	}
}

func (i3 I3Wm) AppExists() bool {
	appPath, err := exec.LookPath(i3.GetName())
	return appPath != "" && err == nil
}

func (i3 I3Wm) GetConfigRoot() string {
	return os.ExpandEnv("$HOME/.config/i3")
}
