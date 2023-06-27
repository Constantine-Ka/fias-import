package service

import "fias-import_byLondon/pkg/repository"

type InstallService struct {
	repo *repository.Repository
}

func (i InstallService) NewTables() bool {
	isSuccess := false
	name := i.repo.Paramtypes("dict_paramtypes", nil)
	isSuccess = name
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Ndoctypes("dict_doctypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Ndockinds("dict_dict_doctinds, nil", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_addhousetypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_addressobjecttypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_apartmenttypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_housetypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_operationtypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_roomtypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Addressobjecttypes("dict_levelstypes", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.AdmHierarchy("hierarchy_adm", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.AdmHierarchy("hierarchy_mun", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Apartments("apartments", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("apartments_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Carplaces("carplaces", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("carplaces_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Houses("houses", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("houses_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Rooms("rooms", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("rooms_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Steads("steads", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("steads_param", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.ChangeHistory("change_history", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.NormDocs("norm_docs", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.ReestrObject("reestr_object", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.AddrObject("object", nil)
	if !isSuccess {
		return false
	}
	isSuccess = i.repo.Params("object_param", nil)
	isSuccess = i.repo.ObjectDivision("object_division", nil)
	if !isSuccess {
		return false
	}

	return isSuccess
}

func NewInstallService(repo *repository.Repository) *InstallService {
	return &InstallService{repo: repo}
}
