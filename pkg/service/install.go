package service

import "fias-import_byLondon/pkg/repository"

type InstallService struct {
	repo repository.CreateTable
}

func (i InstallService) NewTables() bool {
	isSuccess := false
	isSuccess = i.repo.AddrObject("fias_object")
	isSuccess = i.repo.AdmHierarchy("fias_adm_hierarchy")
	isSuccess = i.repo.Apartments("fias_apartment")
	isSuccess = i.repo.Carplaces("fias_car_places")
	isSuccess = i.repo.Houses("fias_houses")
	isSuccess = i.repo.MunHierarchy("fias_mun_hierarchy")
	isSuccess = i.repo.Reestr("fias_reestr")
	isSuccess = i.repo.Rooms("fias_rooms")
	isSuccess = i.repo.Steards("fias_steards")
	return isSuccess
}

func NewInstallService(repo repository.CreateTable) *InstallService {
	return &InstallService{repo: repo}
}
