package service

import (
	"archive/zip"
	"fias-import_byLondon/pkg/repository"
	"fias-import_byLondon/utills/logging"
	"log"
	"regexp"
)

type FileServices struct {
	repo *repository.Repository
}

func (s FileServices) Unpacking(path string) []string {
	logger := logging.GetLogger()
	log.Println("Начало распаковки")
	zipFile, err := zip.OpenReader(path)
	log.Println("Конец распаковки")
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer zipFile.Close()
	var result []string
	for i, f := range zipFile.File {
		if i == 106 {

			result = append(result, f.Name)
			log.Println(f.Name)
			fileReader, err := f.Open()
			if err != nil {
				logger.Fatal(err)
			}
			if regexp.MustCompile(`AS_ADDR_OBJ_\d{8}`).Match([]byte(f.Name)) {
				s.repo.Inserter.AddrObject(cfg.Tablename.Content.AddrObject, s.ParserAddrObj(fileReader))
			} else if regexp.MustCompile(`AS_ADDR_OBJ_DIVISION\d{8}`).Match([]byte(f.Name)) {
				s.repo.Inserter.ObjectDivision(cfg.Tablename.Content.ObjectDivision, s.ParserAddrObjDivision(fileReader))
			}

		}
	}

	return result
}

func NewFileService(repo *repository.Repository) *FileServices {
	return &FileServices{repo: repo}
}
