package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// CreateTarball creates tarball from sourceFolder and saves it to tarballFilePath
//goland:noinspection ALL
func CreateTarball(sourceFolder, tarballFilePath string) error {
	file, err := os.Create(tarballFilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not create tarball file '%s', got error '%s'",
			tarballFilePath, err.Error()))
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	walkErr := filepath.Walk(sourceFolder, func(path string, info fs.FileInfo, err error) error {
		// Skip folders
		if info.IsDir() {
			return nil
		}

		// TODO: Ignore .vers folder
		if strings.Contains(path, ".vers") {
			return nil
		}

		// Normalize path, so it is not absolute anymore
		relativePath := normalizePath(sourceFolder, path)

		tarWriterErr := addFileToTarWriter(path, relativePath, tarWriter)
		if tarWriterErr != nil {
			return tarWriterErr
		}

		return nil
	})

	if walkErr != nil {
		return walkErr
	}

	return nil
}

func addFileToTarWriter(filePath, relativePath string, tarWriter *tar.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not open file '%s', got error '%s'",
			filePath, err.Error()))
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get stat for file '%s', got error '%s'",
			filePath, err.Error()))
	}

	header := &tar.Header{
		Name:    relativePath,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not write header for file '%s', got error '%s'", filePath, err.Error()))
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not copy the file '%s' data to the tarball, got error '%s'", filePath, err.Error()))
	}

	return nil
}

func normalizePath(packFolder, dotfile string) string {
	dotfile = strings.Replace(dotfile, packFolder, "", -1)
	if dotfile[:1] == "/" {
		dotfile = dotfile[1:]
	}

	return dotfile
}
