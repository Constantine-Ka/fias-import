package service

import (
	"archive/zip"
	"fias-import_byLondon/utills/logging"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unpacking(path, dest string) []string {
	logger := logging.GetLogger()

	zipFile, err := zip.OpenReader(path)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer zipFile.Close()
	var result []string
	dst := "output"
	os.MkdirAll("dst", os.ModePerm)
	for _, f := range zipFile.File {
		result = append(result, f.Name)
		filePath := filepath.Join(dst, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return result
}
