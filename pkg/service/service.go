package service

import (
	"fias-import_byLondon/pkg/repository"
)

//var cfg *model.Config

type InstallServices interface {
	NewTables() bool
}
type FileService interface {
	Unpacking(path string) []string
}

type Service struct {
	InstallServices
	FileService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		//InstallServices: NewInstallService(db),
		FileService: NewFileService(db),
	}
}
