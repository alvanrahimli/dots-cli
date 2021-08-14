package utils

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/dlog"
	"io"
	"net/http"
	"os"
)

func CopyFile(sourceFileName, destFileName string) error {
	srcStat, statErr := os.Stat(sourceFileName)
	if statErr != nil {
		dlog.Err(statErr.Error())
		return statErr
	}

	// Do not copy non-regular files
	if !srcStat.Mode().IsRegular() {
		isRegularErr := fmt.Errorf("non-regular source file %s (%q)", srcStat.Name(), srcStat.Mode().String())
		dlog.Err(isRegularErr.Error())
		return isRegularErr
	}

	destStat, statErr := os.Stat(destFileName)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			_, createErr := os.Create(destFileName)
			if createErr != nil {
				return createErr
			}
		} else {
			return statErr
		}

		dlog.Err(statErr.Error())
	} else {
		// Do not copy non-regular files
		if !(destStat.Mode().IsRegular()) {
			err := fmt.Errorf("non-regular destination file %s (%q)",
				destStat.Name(), destStat.Mode().String())
			dlog.Err(err.Error())
			return err
		}
		if os.SameFile(srcStat, destStat) {
			return nil
		}
	}

	// Copy file content
	copyErr := copyFileContent(sourceFileName, destFileName)
	if copyErr != nil {
		dlog.Err(copyErr.Error())
	}
	return copyErr
}

func copyFileContent(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		dlog.Err(err.Error())
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		dlog.Err(err.Error())
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		dlog.Err(err.Error())
		return err
	}

	err = out.Sync()
	if err != nil {
		dlog.Err(err.Error())
	}
	return err
}

func DownloadFile(sourceUrl string) (string, error) {
	fileResponse, getErr := http.Get(sourceUrl)
	if getErr != nil {
		dlog.Err(getErr.Error())
		return "", getErr
	}
	defer fileResponse.Body.Close()

	tmpFile, createErr := os.CreateTemp(os.TempDir(), "dots-pack-archive-*.tar.gz")
	if createErr != nil {
		dlog.Err(createErr.Error())
		return "", createErr
	}
	defer tmpFile.Close()

	_, copyErr := io.Copy(tmpFile, fileResponse.Body)
	if copyErr != nil {
		dlog.Err(copyErr.Error())
		return "", copyErr
	}

	return tmpFile.Name(), nil
}

func GetFromUrl(url string) ([]byte, error) {
	fileResponse, getErr := http.Get(url)
	if getErr != nil {
		dlog.Err(getErr.Error())
		return nil, getErr
	}
	defer fileResponse.Body.Close()

	content, err := io.ReadAll(fileResponse.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
