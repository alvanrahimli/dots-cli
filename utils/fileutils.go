package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CopyFile(sourceFileName, destFileName string) error {
	srcStat, statErr := os.Stat(sourceFileName)
	if statErr != nil {
		return statErr
	}

	// Do not copy non-regular files
	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("non-regular source file %s (%q)", srcStat.Name(), srcStat.Mode().String())
	}

	destStat, statErr := os.Stat(destFileName)
	if statErr != nil {
		if !os.IsNotExist(statErr) {
			return statErr
		}
	} else {
		// Do not copy non-regular files
		if !(destStat.Mode().IsRegular()) {
			return fmt.Errorf("non-regular destination file %s (%q)",
				destStat.Name(), destStat.Mode().String())
		}
		if os.SameFile(srcStat, destStat) {
			return nil
		}
	}

	// Copy file content
	copyErr := copyFileContent(sourceFileName, destFileName)
	return copyErr
}

func copyFileContent(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	err = out.Sync()
	return err
}

func DownloadFile(sourceUrl string) (string, error) {
	fileResponse, getErr := http.Get(sourceUrl)
	if getErr != nil {
		return "", getErr
	}
	defer fileResponse.Body.Close()

	tmpFile, createErr := os.CreateTemp(os.TempDir(), "dots-pack-archive-*.tar.gz")
	if createErr != nil {
		return "", createErr
	}
	defer tmpFile.Close()

	_, copyErr := io.Copy(tmpFile, fileResponse.Body)
	if copyErr != nil {
		return "", copyErr
	}

	return tmpFile.Name(), nil
}
