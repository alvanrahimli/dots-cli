package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"net/http"
	"net/url"
)

// Remote command adds/removes registries for package in current folder
type Remote struct {
	Options *models.Opts
}

func (r Remote) GetArguments() []string {
	return []string{}
}

func (r Remote) CheckRequirements() (bool, string) {
	// [0]remote [1]add [2]origin [3]address

	if !(r.Options.Arguments[1] == "add" || r.Options.Arguments[1] == "remove") {
		return false, "Second argument must be 'add' or 'remove'"
	}

	if r.Options.Arguments[1] == "add" {
		if len(r.Options.Arguments) < 4 {
			return false, "Insufficient amount of arguments entered for 'add'"
		}
	} else if r.Options.Arguments[1] == "remove" {
		if len(r.Options.Arguments) < 3 {
			return false, "Insufficient amount of arguments entered for 'remove'"
		}
	}

	return true, ""
}

func (r Remote) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	// remote add origin https://dots.registry/

	r.Options = opts
	satisfiesRequirements, message := r.CheckRequirements()
	if !satisfiesRequirements {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Init command can not work in this directory:\n\t%s\n", message),
		}
	}

	manifest, err := utils.ReadManifestFile(opts.OutputDir)
	if err != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read manifest. Did you initialize package?",
		}
	}

	var response models.CommandResult
	if r.Options.Arguments[1] == "add" {
		response = r.AddRemote(&manifest)
	} else if r.Options.Arguments[1] == "remove" {
		response = r.RemoveRemote(&manifest)
	}

	return response
}

func (r Remote) AddRemote(manifest *models.Manifest) models.CommandResult {
	// Duplicate remote name check
	for _, remote := range manifest.Remotes {
		if remote.Name == r.Options.Arguments[2] {
			return models.CommandResult{
				Code:    1,
				Message: fmt.Sprintf("Remote with same name (%s) already exists", remote.Name),
			}
		}
	}

	newRemote := models.RemoteAddr{
		Name: r.Options.Arguments[2],
		Url:  r.Options.Arguments[3],
	}

	// Verify remote address
	remoteUrl, urlErr := url.Parse(newRemote.Url)
	if urlErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not parse remote url",
		}
	}

	client := &http.Client{}
	pingUrl := fmt.Sprintf("%s://%s", remoteUrl.Scheme, remoteUrl.Host) + "/ping"
	req, reqErr := http.NewRequest("GET", pingUrl, nil)
	if reqErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not initialize request",
		}
	}

	req.Header.Add("DOTS_CLI_VERSION", models.AppVersion)
	response, doErr := client.Do(req)
	if doErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Error occurred while ping request to '%s'", pingUrl),
		}
	}

	if response.StatusCode == http.StatusTeapot {
		// 418 I'm teapot, when remote address is not dots-server
		return models.CommandResult{
			Code:    1,
			Message: "There is no dots-server on remote machine",
		}
	} else if response.StatusCode == http.StatusBadRequest {
		// 400 Bad Request, when cli-remote version mismatch
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Remote server does not work with dots-cli version: %s", models.AppVersion),
		}
	} else if response.StatusCode != http.StatusOK {
		// Any other non-200 responses
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Error occurred while ping request to '%s'", pingUrl),
		}
	}

	// Add remote to manifest
	manifest.Remotes = append(manifest.Remotes, newRemote)

	manifestErr := utils.WriteManifestFile(r.Options.OutputDir, manifest)
	if manifestErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not save manifest file",
		}
	}

	return models.CommandResult{
		Code:    0,
		Message: "Remote added successfully",
	}
}

func (r Remote) RemoveRemote(manifest *models.Manifest) models.CommandResult {
	remoteName := r.Options.Arguments[2]
	remoteFound := false
	for _, remote := range manifest.Remotes {
		if remote.Name == r.Options.Arguments[2] {
			remoteFound = true
		}
	}

	if !remoteFound {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not find remote '%s' in manifest", remoteName),
		}
	}

	modifiedRemotes, removeErr := utils.RemoveRemote(remoteName, manifest.Remotes)
	if removeErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not remove remote '%s' in manifest", remoteName),
		}
	}

	manifest.Remotes = modifiedRemotes
	manifestErr := utils.WriteManifestFile(r.Options.OutputDir, manifest)
	if manifestErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not save manifest file",
		}
	}
	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Remote '%s' removed from manifest", remoteName),
	}
}
