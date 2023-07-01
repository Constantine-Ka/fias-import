package service

import (
	"fias-import_byLondon/pkg/repository"
)

type InstallService struct {
	repo *repository.Repository
}

//func (i InstallService) NewTables() bool {
//	isSuccess := false
//	tableList := i.repo.ListTable()
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.ParamTypes) == -1 {
//		name := i.repo.CreateTable.Paramtypes(cfg.Tablename.Dict.ParamTypes, nil)
//		isSuccess = name
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Ndoctypes(cfg.Tablename.Dict.NDocTypes, model.NDOCTYPES{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Ndockinds(cfg.Tablename.Dict.NDockInds, model.NDOCKINDS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.AddHouseTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.AddressObjectTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.ApartmentTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.HouseTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.OperationTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.RoomTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.LevelsTypes, nil)
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.AdmHierarchy(cfg.Tablename.Content.AdmHierarchy, model_hierarchy.ADMITEMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.MunHierarchy(cfg.Tablename.Content.MunHierarchy, model_hierarchy.MUNITEMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Apartments(cfg.Tablename.Content.Apartments, model_apartments.APARTMENTS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.ApartmentsP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Carplaces(cfg.Tablename.Content.Carplaces, model_carplaces.CARPLACES{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.CarplacesP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Houses(cfg.Tablename.Content.Houses, model_houses.HOUSES{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.HousesP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Rooms(cfg.Tablename.Content.Rooms, model_rooms.ROOMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.RoomsP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Steads(cfg.Tablename.Content.Steads, model_steads.STEADS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.SteadsP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.ChangeHistory(cfg.Tablename.Content.ChangeHistory, model_other.CHANGEHISTORY{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.NormDocs(cfg.Tablename.Content.NormDocs, model_other.NORMDOCS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.ReestrObject(cfg.Tablename.Content.ReestrObject, model_other.REESTROBJECTS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.AddrObject(cfg.Tablename.Content.AddrObject, model_objectAddr.ADDRESSOBJECTS{})
//		if !isSuccess {
//			return false
//		}
//	}
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.Params(cfg.Tablename.Content.AddrObjectP, model.PARAMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//	if utills.IndexOf(tableList, cfg.Tablename.Dict.NDocTypes) == -1 {
//		isSuccess = i.repo.CreateTable.ObjectDivision(cfg.Tablename.Content.ObjectDivision, model_objectAddr.ITEMS{})
//		if !isSuccess {
//			return false
//		}
//	}
//
//	return isSuccess
//}

func NewInstallService(repo *repository.Repository) *InstallService {
	return &InstallService{repo: repo}
}
