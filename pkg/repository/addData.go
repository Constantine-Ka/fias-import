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
	"strings"
)

type InsertData struct {
	db *sqlx.DB
}

func (i2 InsertData) ParamOne(tableName string, i []model.ParamNode) bool {
	//TODO implement me
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `change_id`, `change_id_end`, `type_id`, `value`, `update_date`, `start_date`, `end_date`) VALUES ", tableName)
	var query string
	for ind, s := range i {
		strings.ReplaceAll(strings.ReplaceAll(s.VALUE, "ж\\", ":"), "\\", "")
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", s.ID, s.OBJECTID, s.CHANGEID, s.CHANGEIDEND, s.TYPEID, s.VALUE, s.UPDATEDATE, s.STARTDATE, s.ENDDATE)
		if (ind%100) == 0 || ind+1 == len(i) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Paramtypes(tableName string, i model.PARAMTYPES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `enddate`, `startdate`, `updatedate`, `isactive`, `code`, `description`, `name`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.PARAMTYPE {
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", s.ID, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.ISACTIVEN, s.CODE, s.DESC, s.NAME, s.Text)
		if (ind%70) == 0 || ind+1 == len(i.PARAMTYPE) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Ndoctypes(tableName string, i model.NDOCTYPES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `enddate`, `startdate`, `name`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.NDOCTYPE {
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%d','%s','%s','%s','%s')", s.ID, s.ENDDATE, s.STARTDATE, s.NAME, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.NDOCTYPE) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Ndockinds(tableName string, i model.NDOCKINDS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `name`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.NDOCKIND {
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%s','%s','%s')", s.ID, s.NAME, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.NDOCKIND) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Addressobjecttypes(tableName string, i model.DICTALL) bool {
	logger := logging.GetLogger()
	var query string
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `isactive`, `enddate`, `startdate`, `updatedate`, `description`, `shortname`, `name`, `level`, `text`) VALUES ", tableName)
	if len(i.ADDHOUSETYPES.HOUSETYPE) >= 1 {
		for index, s := range i.ADDHOUSETYPES.HOUSETYPE {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
			if (index%70) == 0 || index+1 == len(i.ADDHOUSETYPES.HOUSETYPE) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.ADDRESSOBJECTTYPES.ADDRESSOBJECTTYPE) >= 1 {
		for index, s := range i.ADDRESSOBJECTTYPES.ADDRESSOBJECTTYPE {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%s','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, "1970-01-01", s.Text)
			if (index%70) == 0 || index+1 == len(i.ADDRESSOBJECTTYPES.ADDRESSOBJECTTYPE) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.APARTMENTTYPES.APARTMENTTYPE) >= 1 {
		for index, s := range i.APARTMENTTYPES.APARTMENTTYPE {
			//if (index%70) != 0 && index != 0 {
			query = query + ", "
			//}
			query = query + fmt.Sprintf(" ('%d',if(%t,'1', '0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
			if (index%70) == 0 || index+1 == len(i.APARTMENTTYPES.APARTMENTTYPE) {
				query = queryPre + query[1:len(query)]

				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.HOUSETYPES.HOUSETYPE) >= 1 {
		for index, s := range i.HOUSETYPES.HOUSETYPE {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
			if (index%70) == 0 || index+1 == len(i.HOUSETYPES.HOUSETYPE) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.OBJECTLEVELS.OBJECTLEVEL) >= 1 {
		for index, s := range i.OBJECTLEVELS.OBJECTLEVEL {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%s','%s')", 0, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", "", s.NAME, s.LEVEL, s.Text)
			if (index%70) == 0 || index+1 == len(i.OBJECTLEVELS.OBJECTLEVEL) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.OPERATIONTYPES.OPERATIONTYPE) >= 1 {
		for index, s := range i.OPERATIONTYPES.OPERATIONTYPE {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',if(%t,'1', '0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", s.SHORTNAME, s.NAME, 0, s.Text)
			if (index%70) == 0 || index+1 == len(i.OPERATIONTYPES.OPERATIONTYPE) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}
	if len(i.ROOMTYPES.ROOMTYPE) >= 1 {
		for index, s := range i.ROOMTYPES.ROOMTYPE {

			query = query + ", "

			query = query + fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", s.SHORTNAME, s.NAME, 0, s.Text)
			if (index%70) == 0 || index+1 == len(i.ROOMTYPES.ROOMTYPE) {
				query = queryPre + query[1:len(query)]
				_, err := i2.db.Exec(query)
				if err != nil {
					logger.Info(query)
					logger.Error(err)
					return false
				}
				query = ""
			}
		}
	}

	return true
}

func (i2 InsertData) Params(tableName string, i model.PARAMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `change_id`, `change_id_end`, `type_id`, `value`, `update_date`, `start_date`, `end_date`) VALUES ", tableName)
	var query string
	for ind, s := range i.PARAM {
		strings.ReplaceAll(strings.ReplaceAll(s.VALUE, "ж\\", ":"), "\\", "")
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", s.ID, s.OBJECTID, s.CHANGEID, s.CHANGEIDEND, s.TYPEID, s.VALUE, s.UPDATEDATE, s.STARTDATE, s.ENDDATE)
		if (ind%100) == 0 || ind+1 == len(i.PARAM) {
			query = queryPre + query[1:len(query)]
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
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `parent_object`, `change_id`, `region_code`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `path`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.ITEM {
		if ind != 0 {
			query = query + ", "
		}
		query = query + fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',IF(%t,'1','0'),'%s','%s')", s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.REGIONCODE, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.PATH, s.Text)

		if (ind%100) == 0 || ind+1 == len(i.ITEM) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) MunHierarchy(tableName string, i model_hierarchy.MUNITEMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `parent_object_gid`, `change_id`, `oktmo`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `path`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.ITEM {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',IF(%t,'1','0'),'%s','%s','%s')", s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.OKTMO, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, "1", s.PATH, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.ITEM) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Apartments(tableName string, i model_apartments.APARTMENTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `apart_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.APARTMENT {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.APARTTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.APARTMENT) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Carplaces(tableName string, i model_carplaces.CARPLACES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.CARPLACE {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.CARPLACE) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Houses(tableName string, i model_houses.HOUSES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `house_num`, `house_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.HOUSE {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OPERTYPEID, s.OBJECTGUID, s.CHANGEID, s.HOUSENUM, s.HOUSETYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.HOUSE) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) AddrObject(tableName string, i model_objectAddr.ADDRESSOBJECTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `name`, `typename`, `level`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.OBJECT {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%d','%d','%s','%d','%s','%s','%d','%d','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NAME, s.TYPENAME, s.LEVEL, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.OBJECT) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Rooms(tableName string, i model_rooms.ROOMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `room_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.ROOM {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.ROOMTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.ROOM) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) Steads(tableName string, i model_steads.STEADS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.STEAD {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.STEAD) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `parent_id`, `child_id`, `change_id`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.ITEM {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s')", s.ID, s.PARENTID, s.CHILDID, s.CHANGEID, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.ITEM) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`( `object_id`, `adr_object_id`, `change_id`, `oper_type_id`, `change_date`, `ndoc_id`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.ITEM {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%d','%s')", s.OBJECTID, s.ADROBJECTID, s.CHANGEID, s.OPERTYPEID, s.CHANGEDATE, s.NDOCID, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.ITEM) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) NormDocs(tableName string, i model_other.NORMDOCS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `name`, `number`, `type`, `date`, `update_date`, `kind`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.NORMDOC {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s','%s')", s.ID, s.NAME, s.NUMBER, s.TYPE, s.DATE, s.UPDATEDATE, s.KIND, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.NORMDOC) {
			query = queryPre + query[1:len(query)]
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

func (i2 InsertData) ReestrObject(tableName string, i model_other.REESTROBJECTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`object_id`, `object_gid`, `change_id`, `level_id`, `update_date`, `create_date`, `text`) VALUES ", tableName)
	var query string
	for ind, s := range i.OBJECT {
		if ind != 0 {
			query = query + ", "
		}
		query = fmt.Sprintf(" ('%s','%s','%s','%s','%s','%s','%s')", s.OBJECTGUID, s.OBJECTGUID, s.CHANGEID, s.LEVELID, s.UPDATEDATE, s.CREATEDATE, s.Text)
		if (ind%100) == 0 || ind+1 == len(i.OBJECT) {
			query = queryPre + query[1:len(query)]
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

func NewInsert(db *sqlx.DB) *InsertData {
	return &InsertData{db: db}
}
