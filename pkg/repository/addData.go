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
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type InsertData struct {
	db *sqlx.DB
}

func (i2 InsertData) Params(tableName string, i model.PARAMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `change_id`, `change_id_end`, `type_id`, `value`, `update_date`, `start_date`, `end_date`) VALUES ", tableName)
	var query string
	for ind, s := range i.PARAM {
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%s','%s')", s.ID, s.OBJECTID, s.CHANGEID, s.CHANGEIDEND, s.TYPEID, s.VALUE, s.UPDATEDATE, s.STARTDATE, s.ENDDATE)
		if (ind%70) == 0 || ind+1 == len(i.PARAM) {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			query = ""
		}
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
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `parent_object_gid`, `change_id`, `oktmo`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `path`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.OKTMO, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, "1", s.PATH, s.Text)
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
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `apart_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.APARTTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
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
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTUAL, s.Text)
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
	for _, s := range i.HOUSE {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `house_num`, `house_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OPERTYPEID, s.OBJECTGUID, s.CHANGEID, s.HOUSENUM, s.HOUSETYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) AddrObject(tableName string, i model_objectAddr.ADDRESSOBJECTS) bool {
	logger := logging.GetLogger()
	for _, s := range i.OBJECT {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `name`, `typename`, `level`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%d','%d','%s','%d','%s','%s','%d','%d','%d','%d','%s','%s','%s','%d','%d','%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NAME, s.TYPENAME, s.LEVEL, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) Rooms(tableName string, i model_rooms.ROOMS) bool {
	logger := logging.GetLogger()
	for _, s := range i.ROOM {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `room_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.ROOMTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) Steads(tableName string, i model_steads.STEADS) bool {
	logger := logging.GetLogger()
	for _, s := range i.STEAD {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool {
	logger := logging.GetLogger()
	for _, s := range i.ITEM {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `parent_id`, `child_id`, `change_id`, `text`) VALUES ('%s','%s','%s','%s','%s')", tableName, s.ID, s.PARENTID, s.CHILDID, s.CHANGEID, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool {
	logger := logging.GetLogger()
	for _, s := range i.ITEM {
		query := fmt.Sprintf("INSERT INTO `%s`( `object_id`, `adr_object_id`, `change_id`, `oper_type_id`, `change_date`, `ndoc_id`, `text`) VALUES ('%s','%s','%s','value-5]','%s','%s','%s')", tableName, s.OBJECTID, s.ADROBJECTID, s.CHANGEID, s.OPERTYPEID, s.CHANGEDATE, s.NDOCID, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) NormDocs(tableName string, i model_other.NORMDOCS) bool {
	logger := logging.GetLogger()
	for _, s := range i.NORMDOC {
		query := fmt.Sprintf("INSERT INTO `%s`(`id`, `name`, `number`, `type`, `date`, `update_date`, `kind`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s')", tableName, s.ID, s.NAME, s.NUMBER, s.TYPE, s.DATE, s.UPDATEDATE, s.KIND, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func (i2 InsertData) ReestrObject(tableName string, i model_other.REESTROBJECTS) bool {
	logger := logging.GetLogger()
	for _, s := range i.OBJECT {
		query := fmt.Sprintf("INSERT INTO `%s`(`object_id`, `object_gid`, `change_id`, `level_id`, `update_date`, `create_date`, `text`) VALUES ('%s','%s','%s','%s','%s','%s','%s')", tableName, s.OBJECTGUID, s.OBJECTGUID, s.CHANGEID, s.LEVELID, s.UPDATEDATE, s.CREATEDATE, s.Text)
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
	}
	return true
}

func NewInsert(db *sqlx.DB) *InsertData {
	return &InsertData{db: db}
}
