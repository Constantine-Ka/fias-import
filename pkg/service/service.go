package service

import (
	"fias-import_byLondon/pkg/repository"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

//var cfg *model.Config

type InstallServices interface {
	NewTables() bool
}
type FileService interface {
	Unpacking(path, typefile string, clientsftp *sftp.Client, clientssh *ssh.Client) []string
}

type Service struct {
	//InstallServices
	FileService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		//InstallServices: NewInstallService(db),
		FileService: NewFileService(db),
	}
}
