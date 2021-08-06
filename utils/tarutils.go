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

// CreateTarball creates tarball out of sourceFolder and saves it to tarballFilePath
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

		// Do not include .vers folder to tarball
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

func UnTar(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {
		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)
		// fmt.Println("NOW TARGET: " + target)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.Create(target)
			if err != nil {
				// fmt.Println("Could not create " + target + ". Creating parent directories")
				folders := strings.Split(target, "/")
				mkdirAllErr := os.MkdirAll(strings.Join(folders[:len(folders)-1], "/"), os.ModePerm)
				if mkdirAllErr != nil {
					fmt.Printf("Could not create parent directories for %s. Skipping this file\n", target)
					fmt.Printf("ERROR: %s\n", mkdirAllErr.Error())
					continue
				}
				f, _ = os.Create(target)
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; deferring would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
