package repository

import (
	"fias-import_byLondon/model"
	"github.com/jmoiron/sqlx"
)

type CreateTable interface {
	AddrObject(tableName string) bool
	AdmHierarchy(tableName string) bool
	Apartments(tableName string) bool
	Carplaces(tableName string) bool
	Houses(tableName string) bool
	MunHierarchy(tableName string) bool
	Reestr(tableName string) bool
	Rooms(tableName string) bool
	Steards(tablename string) bool
}
type Inserter interface {
	InsertObjAddr(tablename string, data model.Objects) bool
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
