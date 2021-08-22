package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/dlog"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
	"path"
	"strings"
)

type Install struct {
	Options *models.Opts
}

func (i Install) GetArguments() []string {
	return []string{}
}

func (i Install) CheckRequirements() (bool, string) {
	return true, ""
}

func (i Install) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	// Backup old files
	i.Options = opts
	satisfiesRequirements, message := i.CheckRequirements()
	if !satisfiesRequirements {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Push command can not work in this directory:\n\t%s\n", message),
		}
	}

	manifest, err := utils.ReadManifestFile(opts.OutputDir)
	if err != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read manifest. Did you initialize package?",
		}
	}

	// TODO: Handle mismatching versions
	// BACKUP
	for _, app := range manifest.Apps {
		dlog.Info("Backup is in progress for '%s'")

		// Get app dotfiles
		var handler models.Handler
		for appHandler, details := range config.Handlers {
			if appHandler == app.Name {
				handler = details
			}
		}

		// Create .backup
		backupPath := path.Join(i.Options.OutputDir, ".backup")
		_, statErr := os.Stat(backupPath)
		if statErr != nil {
			mkdirErr := os.Mkdir(backupPath, os.ModePerm)
			if mkdirErr != nil {
				return models.CommandResult{
					Code:    1,
					Message: "Could not create backup folder",
				}
			}
		}

		// Copy dotfile to .backup
		for _, dotfile := range handler.Dotfiles {
			appSpecificDir := path.Join(backupPath, app.Name)

			dotfileExpanded := os.ExpandEnv(dotfile)
			rootExpanded := os.ExpandEnv(handler.ConfigRoot)
			relBackupPath := strings.ReplaceAll(dotfileExpanded, rootExpanded, "")
			if path.IsAbs(relBackupPath) {
				relBackupPath = relBackupPath[1:]
			}

			dotfileBackupPath := path.Join(appSpecificDir, relBackupPath)

			_, statErr := os.Stat(path.Dir(dotfileBackupPath))
			if statErr != nil {
				mkdirErr := os.MkdirAll(path.Dir(dotfileBackupPath), os.ModePerm)
				if mkdirErr != nil {
					return models.CommandResult{
						Code:    1,
						Message: fmt.Sprintf("Could not create app backup folder for %s", app.Name),
					}
				}
			}

			copyErr := utils.CopyFile(dotfileExpanded, dotfileBackupPath)
			if copyErr != nil {
				dlog.Err(copyErr.Error())
				if os.IsNotExist(copyErr) {
					fmt.Printf("Could not backup '%s' for '%s'. File not found\n", dotfileExpanded, app.Name)
				}
				continue
			}
		}
	}
	dlog.Info("Successfully created backup for package %s", manifest.Name)

	// INSTALLATION
	for _, app := range manifest.Apps {
		var appHandler models.Handler
		for handlerName, handler := range config.Handlers {
			if handlerName == app.Name {
				appHandler = handler
				break
			}
		}

		for _, dotfile := range appHandler.Dotfiles {
			dotfile = os.ExpandEnv(dotfile)
			rel2Pack := strings.ReplaceAll(dotfile, os.ExpandEnv(appHandler.ConfigRoot), "")
			if path.IsAbs(rel2Pack) {
				rel2Pack = rel2Pack[1:]
			}

			inPackPath := path.Join(i.Options.OutputDir, app.Name, rel2Pack)
			_, statErr := os.Stat(dotfile)
			// If old dotfile found
			if statErr == nil {
				removeErr := os.Remove(dotfile)
				if removeErr != nil {
					dlog.Info("Could not remove %s", dotfile)
				}
			}

			copyErr := utils.CopyFile(inPackPath, dotfile)
			if copyErr != nil {
				dlog.Warn("Could not copy %s to %s\nError: %s", inPackPath, dotfile, copyErr.Error())
			}
		}
	}

	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Successfully installed package %s", manifest.Name),
	}
}
