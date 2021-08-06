package utils

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
)

func IndexOfApp(appName string, list []models.App) int {
	for i, app := range list {
		if app.Name == appName {
			return i
		}
	}

	return -1
}

func RemoveApp(appName string, list []models.App) ([]models.App, error) {
	index := IndexOfApp(appName, list)
	// If app found
	if index != -1 {
		return append(list[:index], list[index+1:]...), nil
	}

	return nil, fmt.Errorf("could not find app '%s'", appName)
}

func IndexOfRemote(remoteName string, list []models.RemoteAddr) int {
	for i, remote := range list {
		if remote.Name == remoteName {
			return i
		}
	}

	return -1
}

func RemoveRemote(remoteName string, list []models.RemoteAddr) ([]models.RemoteAddr, error) {
	index := IndexOfRemote(remoteName, list)
	// If app found
	if index != -1 {
		return append(list[:index], list[index+1:]...), nil
	}

	return nil, fmt.Errorf("could not find remote '%s'", remoteName)
}
