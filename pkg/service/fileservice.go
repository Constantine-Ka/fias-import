package service

import (
	"archive/zip"
	"fias-import_byLondon/utills/logging"
	"io"
	"os"
	"path/filepath"
)

func Unpacking(path, dest string) []string {
	logger := logging.GetLogger()

	zipFile, err := zip.OpenReader(path)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer zipFile.Close()
	var result []string
	os.MkdirAll("tmp", os.ModePerm)
	for _, file := range zipFile.File {
		result = append(result, file.Name)
		// Get the corresponding file in the destination
		dstFile, err := os.OpenFile(
			filepath.Join(dest, file.Name),
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			file.Mode(),
		)
		if err != nil {
			logger.Error(err)
		}

		// Copy the contents of the zip to the destination
		rc, err := file.Open()
		if err != nil {
			logger.Error(err)
		}
		_, err = io.Copy(dstFile, rc)

		// Close the file handles
		rc.Close()
		dstFile.Close()

		if err != nil {
			logger.Error(err)

		}
	}

	return result
}
