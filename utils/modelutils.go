package utils

import (
	"github.com/alvanrahimli/dots-cli/models"
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
