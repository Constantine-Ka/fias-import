package repository

import (
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

type InsertData struct {
	db *sqlx.DB
}

func (i2 InsertData) Params(tableName string, i interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) AdmHierarchy(tableName string, i model_hierarchy.ADMITEMS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) MunHierarchy(tableName string, i model_hierarchy.MUNITEMS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) Apartments(tableName string, i model_apartments.APARTMENTS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) Carplaces(tableName string, i model_carplaces.CARPLACES) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) Houses(tableName string, i model_houses.HOUSES) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) AddrObject(tableName string, i interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) Rooms(tableName string, i model_rooms.ROOMS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) Steads(tableName string, i model_steads.STEADS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) NormDocs(tableName string, i model_other.NORMDOCS) bool {
	//TODO implement me
	panic("implement me")
}

func (i2 InsertData) ReestrObject(tableName string, i model_other.REESTROBJECTS) bool {
	//TODO implement me
	panic("implement me")
}

func NewInsert(db *sqlx.DB) *InsertData {
	return &InsertData{db: db}
}
