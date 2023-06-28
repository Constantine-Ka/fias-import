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
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type InsertData struct {
	db *sqlx.DB
}

func (i2 InsertData) Params(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) AdmHierarchy(tableName string, i model_hierarchy.ADMITEMS) bool {
	logger := logging.GetLogger()
	for _, s := range i.ITEM {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `parent_object`, `change_id`, `region_code`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `path`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.REGIONCODE, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.PATH, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Panic(err)
			return false
		}
	}

	return true
}

func (i2 InsertData) MunHierarchy(tableName string, i model_hierarchy.MUNITEMS) bool {
	logger := logging.GetLogger()
	for _, s := range i.ITEM {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `parent_object_gid`, `change_id`, `oktmo`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `path`, `text`) VALUES ('[value-1]','[value-2]','[value-3]','[value-4]','[value-5]','[value-6]','[value-7]','[value-8]','[value-9]','[value-10]','[value-11]','[value-12]','[value-13]','[value-14]')", tableName, s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.OKTMO, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, "1", s.PATH, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) Apartments(tableName string, i model_apartments.APARTMENTS) bool {
	logger := logging.GetLogger()
	for _, s := range i.APARTMENT {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `apart_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('[value-1]','[value-2]','[value-3]','[value-4]','[value-5]','[value-6]','[value-7]','[value-8]','[value-9]','[value-10]','[value-11]','[value-12]','[value-13]','[value-14]','[value-15]')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.APARTTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) Carplaces(tableName string, i model_carplaces.CARPLACES) bool {
	logger := logging.GetLogger()
	for _, s := range i.CARPLACE {
		query := fmt.Sprintf("", tableName)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) Houses(tableName string, i model_houses.HOUSES) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) AddrObject(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) Rooms(tableName string, i model_rooms.ROOMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) Steads(tableName string, i model_steads.STEADS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) NormDocs(tableName string, i model_other.NORMDOCS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (i2 InsertData) ReestrObject(tableName string, i model_other.REESTROBJECTS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("", tableName)
	_, err := i2.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func NewInsert(db *sqlx.DB) *InsertData {
	return &InsertData{db: db}
}
