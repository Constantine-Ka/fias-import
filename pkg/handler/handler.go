package handler

import "fias-import_byLondon/pkg/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
