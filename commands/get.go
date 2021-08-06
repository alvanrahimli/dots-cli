package commands

import (
	"encoding/json"
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"io"
	"log"
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
	request, reqErr := http.NewRequest("GET", config.Registry+models.PackagesEndpoint+packageName, nil)
	if reqErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not initialize request",
		}
	}

	res, doErr := client.Do(request)
	if doErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Error occurred while getting response",
		}
	}

	if res.StatusCode == http.StatusNotFound {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not find package with name: %s", packageName),
		}
	}

	defer res.Body.Close()
	responseBody, bodyErr := io.ReadAll(res.Body)
	if bodyErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read response data",
		}
	}

	response := models.GetPackagesResponse{}
	jsonErr := json.Unmarshal(responseBody, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
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
		fmt.Print("\n\n\tEnter version: ")

		var version string
		_, scanErr := fmt.Scanln(&version)
		if scanErr != nil {
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
	archiveUrl = config.Registry + archiveUrl
	archiveFilePath, downloadErr := utils.DownloadFile(archiveUrl)
	if downloadErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: fmt.Sprintf("Could not download file from %s", archiveUrl),
		}
	}

	// Extract archiveFile to output directory
	archiveFile, openErr := os.Open(archiveFilePath)
	if openErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not open downloaded archive",
		}
	}

	unTarErr := utils.UnTar(g.Options.OutputDir, archiveFile)
	if unTarErr != nil {
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
