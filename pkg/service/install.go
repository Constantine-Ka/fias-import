package service

import (
	"fias-import_byLondon/model"
	model_apartments "fias-import_byLondon/model/model-apartments"
	model_carplaces "fias-import_byLondon/model/model-carplaces"
	model_hierarchy "fias-import_byLondon/model/model-hierarchy"
	model_houses "fias-import_byLondon/model/model-houses"
	model_objectAddr "fias-import_byLondon/model/model-objectAddr"
	model_other "fias-import_byLondon/model/model-other"
	model_rooms "fias-import_byLondon/model/model-rooms"
	model_steads "fias-import_byLondon/model/model-steads"
	"fias-import_byLondon/pkg/repository"
)

type InstallService struct {
	repo *repository.Repository
}

func (i InstallService) NewTables() bool {
	isSuccess := false
	name := i.repo.CreateTable.Paramtypes("dict_paramtypes", nil)
	isSuccess = name
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Ndoctypes("dict_doctypes", model.NDOCTYPES{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Ndockinds("dict_doctkinds", model.NDOCKINDS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_addhousetypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_addressobjecttypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_apartmenttypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_housetypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_operationtypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_roomtypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Addressobjecttypes("dict_levelstypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.AdmHierarchy("hierarchy_adm", model_hierarchy.ADMITEMS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.MunHierarchy("hierarchy_mun", model_hierarchy.MUNITEMS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Apartments("apartments", model_apartments.APARTMENTS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("apartments_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Carplaces("carplaces", model_carplaces.CARPLACES{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("carplaces_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Houses("houses", model_houses.HOUSES{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("houses_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Rooms("rooms", model_rooms.ROOMS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("rooms_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Steads("steads", model_steads.STEADS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("steads_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.ChangeHistory("change_history", model_other.CHANGEHISTORY{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.NormDocs("norm_docs", model_other.NORMDOCS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.ReestrObject("reestr_object", model_other.REESTROBJECTS{})
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.AddrObject("object", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.CreateTable.Params("object_param", nil)
	isSuccess = i.repo.CreateTable.ObjectDivision("object_division", model_objectAddr.ITEMS{})
	if !isSuccess {
		return false
	}

	return isSuccess
}

func NewInstallService(repo *repository.Repository) *InstallService {
	return &InstallService{repo: repo}
}
