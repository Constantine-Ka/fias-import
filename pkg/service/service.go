package service

import (
	"fias-import_byLondon/pkg/repository"
	"github.com/spf13/viper"
)

//var cfg *model.Config

type InstallServices interface {
	NewTables() bool
}
type FileService interface {
	Unpacking(path, typefile string) []string
}

type Service struct {
	//InstallServices
	FileService
}

func NewService(db *repository.Repository, vp *viper.Viper) *Service {
	return &Service{
		//InstallServices: NewInstallService(db),
		FileService: NewFileService(db, vp),
	}
}
