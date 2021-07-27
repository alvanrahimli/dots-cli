package utils

import (
	"dots/models"
	"encoding/json"
	"os"
	"path"
)

func ReadManifestFile(dir string) (models.Manifest, error) {
	manifestFileAddr := path.Join(dir, "manifest.json")
	_, statErr := os.Stat(manifestFileAddr)
	if statErr != nil {
		return models.Manifest{}, statErr
	}

	manifestBytes, readErr := os.ReadFile(manifestFileAddr)
	if readErr != nil {
		return models.Manifest{}, readErr
	}

	manifest := models.Manifest{}
	jsonErr := json.Unmarshal(manifestBytes, &manifest)
	if jsonErr != nil {
		return models.Manifest{}, jsonErr
	}

	return manifest, nil
}

func AppExistsInManifest(appName string, manifest *models.Manifest) bool {
	for _, app := range manifest.Apps {
		if app.Name == appName {
			return true
		}
	}

	return false
}
