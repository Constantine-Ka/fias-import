package service

import (
	"fias-import_byLondon/model"
	"fias-import_byLondon/pkg/repository"
	"io"
)

type InstallServices interface {
	NewTables() bool
}
type FileService interface {
	Unpacking(path string) []string
	ParserAddrObj(reader io.Reader) model.ADDRESSOBJECTS
}

type Service struct {
	InstallServices
	FileService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		InstallServices: NewInstallService(db),
		FileService:     NewFileService(db),
	}
}
