package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"os/exec"
)

type List struct {
	Options *models.Opts
}

func (l List) GetArguments() []string {
	return []string{}
}

func (l List) CheckRequirements() (bool, string) {
	if len(l.Options.Arguments) < 2 {
		return false, "Insufficient amount of arguments entered for 'list'"
	}

	cmd := l.Options.Arguments[1]
	if !(cmd == "all" || cmd == "added") {
		return false, "Second argument can only be all|installed|added"
	}

	return true, ""
}

func (l List) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	l.Options = opts
	satisfiesRequirements, message := l.CheckRequirements()
	if !satisfiesRequirements {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("List command can not:\n\t%s\n", message),
		}
	}

	// Read manifest
	manifest, err := utils.ReadManifestFile(opts.OutputDir)
	if err != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read manifest. Did you initialize package?",
		}
	}

	subCmd := l.Options.Arguments[1]
	switch subCmd {
	case "all":
		i := 1
		for name, handler := range config.Handlers {
			if l.Options.Installed {
				_, err := exec.LookPath(name)
				if err == nil {
					fmt.Printf("\n\t%d. %s (v: %s)", i, name, handler.Version)
				}
			} else {
				fmt.Printf("\n\t%d. %s (v: %s)", i, name, handler.Version)
			}

			i++
		}
		break
	case "added":
		if len(manifest.Apps) == 0 {
			fmt.Println("\n\tNo apps found in package")
			break
		}
		for i, app := range manifest.Apps {
			fmt.Printf("\n\t%d. %s (v: %s)", i+1, app.Name, app.Version)
		}
		break
	}
	fmt.Print("\n\n")

	return models.CommandResult{
		Code:    0,
		Message: "Apps listed",
	}
}
