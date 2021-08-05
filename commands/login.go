package commands

import (
	"encoding/json"
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/alvanrahimli/dots-cli/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Login command signs user in and saves token
type Login struct {
	Options *models.Opts
}

func (l Login) GetArguments() []string {
	return []string{}
}

func (l Login) CheckRequirements() (bool, string) {
	return true, ""
}

func (l Login) ExecuteCommand(opts *models.Opts, config *models.AppConfig) models.CommandResult {
	var email, password string

	// Ask for email
	fmt.Print("Enter email: ")
	_, scanErr := fmt.Scanln(&email)
	if scanErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read username",
		}
	}

	// Ask for password
	fmt.Print("Enter password: ")
	_, scanErr = fmt.Scanln(&password)
	if scanErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not read password",
		}
	}

	// Send POST request
	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	client := &http.Client{}
	req, reqErr := http.NewRequest("POST", models.LoginEndpoint, strings.NewReader(data.Encode()))
	if reqErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not initialize request",
		}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, resErr := client.Do(req)
	if resErr != nil {
		fmt.Println(resErr.Error())
		return models.CommandResult{
			Code:    1,
			Message: "Could not get response",
		}
	}

	if res.StatusCode == http.StatusUnauthorized {
		return models.CommandResult{
			Code:    1,
			Message: "Invalid email or password provided",
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

	loginResponse := models.LoginResponse{}
	jsonErr := json.Unmarshal(responseBody, &loginResponse)
	if jsonErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not parse response data. Check app version",
		}
	}

	if config.AuthorName != loginResponse.Data.User.Username {
		fmt.Printf("Signed User's name (%s) does not match with entered author name (%s)",
			loginResponse.Data.User.Username, config.AuthorName)
		fmt.Println("Would you like to change it? [Y/n] ")
		var choice string
		_, scanErr = fmt.Scanln(&choice)
		if scanErr != nil || choice == "y" || choice == "Y" {
			config.AuthorName = loginResponse.Data.User.Username
		}
	}

	if config.AuthorEmail != loginResponse.Data.User.Email {
		fmt.Printf("Signed User's email (%s) does not match with entered author email (%s)",
			loginResponse.Data.User.Email, config.AuthorEmail)
		fmt.Println("Would you like to change it? [Y/n] ")
		var choice string
		_, scanErr = fmt.Scanln(&choice)
		if scanErr != nil || choice == "y" || choice == "Y" {
			config.AuthorEmail = loginResponse.Data.User.Email
		}
	}

	config.AuthorToken = loginResponse.Data.Token
	saveErr := utils.SaveConfig(config)
	if saveErr != nil {
		return models.CommandResult{
			Code:    1,
			Message: "Could not save token",
		}
	}

	return models.CommandResult{
		Code:    0,
		Message: fmt.Sprintf("User '%s' signed in", loginResponse.Data.User.Username),
	}
}
