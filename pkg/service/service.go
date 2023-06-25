package service

import (
	"fias-import_byLondon/pkg/repository"
)

type ZipService interface {
	Unpacking(path string) []string
}

type Service struct {
	ZipService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
