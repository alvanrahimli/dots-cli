package commands

import (
	"encoding/json"
	"fmt"
	"github.com/alvanrahimli/dots-cli/dlog"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"io"
	"net/http"
	"os"
)

type Get struct {
	Options *models.Opts
}

func (g Get) GetArguments() []string {
	return []string{}
}

func (g Get) CheckRequirements() (bool, string) {
	if len(g.Options.Arguments) < 2 {
		return false, "Insufficient amount of arguments entered for 'get'"
	}

	return true, ""
}

func (g Get) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	g.Options = opts
	satisfiesRequirements, message := g.CheckRequirements()
	if !satisfiesRequirements {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("List command can not:\n\t%s\n", message),
		}
	}

	packageName := g.Options.Arguments[1]

	client := &http.Client{}
	// config.Registry + models.PackagesEndpoint + packageName
	getUrl := fmt.Sprintf("%s/%s/%s", config.Registry, models.PackagesEndpoint, packageName)
	dlog.Info("Sending HTTP GET to %s", getUrl)
	request, reqErr := http.NewRequest("GET", getUrl, nil)
	if reqErr != nil {
		dlog.Err(reqErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not initialize request",
		}
	}

	res, doErr := client.Do(request)
	if doErr != nil {
		dlog.Err(doErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Error occurred while getting response",
		}
	}

	if res.StatusCode == http.StatusNotFound {
		dlog.Err("Could not find package %s\n", packageName)
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not find package with name: %s", packageName),
		}
	}

	defer res.Body.Close()
	responseBody, bodyErr := io.ReadAll(res.Body)
	if bodyErr != nil {
		dlog.Err(bodyErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not read response data",
		}
	}

	response := models.GetPackagesResponse{}
	jsonErr := json.Unmarshal(responseBody, &response)
	if jsonErr != nil {
		dlog.Err(jsonErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not parse response data. Check app version",
		}
	}

	// Select package version to download
	var archiveUrl string
	if g.Options.Version != "" {
		for _, pack := range response.Data.Packages {
			if pack.Version == g.Options.Version {
				archiveUrl = pack.ArchiveName
			}
		}
	} else {
		fmt.Println("Multiple versions found. Please enter version number you want to get")
		for i, pack := range response.Data.Packages {
			fmt.Printf("\n\t%d) %s", i+1, pack.Version)
		}
		fmt.Print("\n\nEnter version: ")

		var version string
		_, scanErr := fmt.Scanln(&version)
		if scanErr != nil {
			dlog.Err(scanErr.Error())
			return models.CommandResult{
				Code:    1,
				Message: "Could not read version number",
			}
		}

		for _, pack := range response.Data.Packages {
			if pack.Version == version {
				archiveUrl = pack.ArchiveName
			}
		}
	}

	// Download archive to /tmp
	archiveUrl = fmt.Sprintf("%s/%s", config.Registry, archiveUrl)
	archiveFilePath, downloadErr := utils.DownloadFile(archiveUrl)
	if downloadErr != nil {
		dlog.Err(downloadErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not download file from %s", archiveUrl),
		}
	}

	// Extract archiveFile to output directory
	archiveFile, openErr := os.Open(archiveFilePath)
	if openErr != nil {
		dlog.Err(openErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not open downloaded archive",
		}
	}

	unTarErr := utils.UnTar(g.Options.OutputDir, archiveFile)
	if unTarErr != nil {
		dlog.Err(unTarErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not un-tar archive",
		}
	}

	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("Extracted package archive to '%s'", g.Options.OutputDir),
	}
}
