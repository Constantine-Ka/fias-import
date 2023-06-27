package service

import (
	model_objectAddr "fias-import_byLondon/model/model-objectAddr"
	"fias-import_byLondon/pkg/repository"
	"io"
)

type InstallServices interface {
	NewTables() bool
}
type FileService interface {
	Unpacking(path string) []string
	ParserAddrObj(reader io.Reader) model_objectAddr.ADDRESSOBJECTS
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
