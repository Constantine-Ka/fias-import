package repository

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
	"github.com/jmoiron/sqlx"
)

type dict interface {
	Paramtypes(tableName string, i interface{}) bool
	Ndoctypes(tableName string, i model.NDOCTYPES) bool
	Ndockinds(tableName string, i model.NDOCKINDS) bool
	Addressobjecttypes(tableName string, i interface{}) bool
}
type dictInsert interface {
	//Paramtypes(tableName string, i interface{}) bool
	//Ndoctypes(tableName string, i interface{}) bool
	//Ndockinds(tableName string, i interface{}) bool
	//Addressobjecttypes(tableName string, i interface{}) bool
}
type content interface {
	Params(tableName string, i model.PARAMS) bool
	AdmHierarchy(tableName string, i model_hierarchy.ADMITEMS) bool
	MunHierarchy(tableName string, i model_hierarchy.MUNITEMS) bool
	Apartments(tableName string, i model_apartments.APARTMENTS) bool
	Carplaces(tableName string, i model_carplaces.CARPLACES) bool
	Houses(tableName string, i model_houses.HOUSES) bool
	AddrObject(tableName string, i model_objectAddr.ADDRESSOBJECTS) bool
	Rooms(tableName string, i model_rooms.ROOMS) bool
	Steads(tableName string, i model_steads.STEADS) bool
	ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool
	ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool
	NormDocs(tableName string, i model_other.NORMDOCS) bool
	ReestrObject(tableName string, i model_other.REESTROBJECTS) bool
}

type CreateTable interface {
	dict
	content
}
type Inserter interface {
	content
}
type Repository struct {
	CreateTable
	Inserter
	//
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CreateTable: NewCreateTable(db),
		Inserter:    NewInsert(db),
	}
}
