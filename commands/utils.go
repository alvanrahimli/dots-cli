package commands

import (
	"dots/models"
	"encoding/json"
	"os"
	"path"
	"strconv"
	"strings"
)

func NewPackageVersion(versionStr string) models.Version {
	versionNumbersStr := strings.Split(versionStr, ".")
	majorVersion, _ := strconv.Atoi(versionNumbersStr[0])
	minorVersion, _ := strconv.Atoi(versionNumbersStr[1])
	patchVersion, _ := strconv.Atoi(versionNumbersStr[2])

	return models.Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}
}

func WriteManifestFile(outDir string, manifest *models.Manifest) error {
	manifestBytes, marshallErr := json.MarshalIndent(manifest, "", "  ")
	if marshallErr != nil {
		return marshallErr
	}

	fileDir := path.Join(outDir, "manifest.json")
	writeErr := os.WriteFile(fileDir, manifestBytes, os.ModePerm)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
