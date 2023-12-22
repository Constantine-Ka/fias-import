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
	"github.com/spf13/viper"
	"strings"
)

type InsertData struct {
	db  *sqlx.DB
	cfg *viper.Viper
}

func (i2 InsertData) ParamOne(tableName string, i []model.ParamNode) bool {
	//TODO implement me
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `change_id`, `change_id_end`, `type_id`, `value`, `update_date`, `start_date`, `end_date`) VALUES ", tableName)
	var query string
	for _, s := range i {
		//now := time.Now()
		//enddate, err := t	ime.Parse(s.ENDDATE, "2022-10-06")
		//if err != nil {
		//	logger.Error(err)
		//
		//	return false
		//}
		//if now.Compare(enddate) > 0 {
		strings.ReplaceAll(strings.ReplaceAll(s.VALUE, "ж\\", ":"), "\\", "")

		query = fmt.Sprintf(" ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", s.ID, s.OBJECTID, s.CHANGEID, s.CHANGEIDEND, s.TYPEID, s.VALUE, s.UPDATEDATE, s.STARTDATE, s.ENDDATE)

		query = queryPre + query
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
		//}

		query = ""
	}

	return true
}

func (i2 InsertData) Paramtypes(tableName string, i model.PARAMTYPES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `enddate`, `startdate`, `updatedate`, `isactive`, `code`, `description`, `name`, `text`) VALUES ", tableName)
	var query string
	for _, s := range i.PARAMTYPE {

		query = fmt.Sprintf(" ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", s.ID, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.ISACTIVEN, s.CODE, s.DESC, s.NAME, s.Text)
		query = queryPre + query
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
		query = ""
	}
	return true
}

func (i2 InsertData) Ndoctypes(tableName string, i model.NDOCTYPES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `enddate`, `startdate`, `name`, `text`) VALUES ", tableName)
	var query string
	for _, s := range i.NDOCTYPE {

		query = fmt.Sprintf(" ('%d','%s','%s','%s','%s')", s.ID, s.ENDDATE, s.STARTDATE, s.NAME, s.Text)
		query = queryPre + query
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
		query = ""
	}

	return true
}

func (i2 InsertData) Ndockinds(tableName string, i model.NDOCKINDS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `name`, `text`) VALUES ", tableName)
	var query string
	for _, s := range i.NDOCKIND {

		query = fmt.Sprintf(" ('%s','%s','%s')", s.ID, s.NAME, s.Text)
		query = queryPre + query
		_, err := i2.db.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
		query = ""
	}

	return true
}

func (i2 InsertData) Addressobjecttypes(tableName string, i model.DICTALL) bool {
	logger := logging.GetLogger()
	var query string
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `isactive`, `enddate`, `startdate`, `updatedate`, `description`, `shortname`, `name`, `level`, `text`) VALUES ", tableName)
	if len(i.ADDHOUSETYPES.HOUSETYPE) >= 1 {
		for _, s := range i.ADDHOUSETYPES.HOUSETYPE {

			query = fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
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

	if len(i.ADDRESSOBJECTTYPES.ADDRESSOBJECTTYPE) >= 1 {
		for _, s := range i.ADDRESSOBJECTTYPES.ADDRESSOBJECTTYPE {

			query = fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%s','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, s.LEVEL, s.Text)
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
	if len(i.APARTMENTTYPES.APARTMENTTYPE) >= 1 {
		for _, s := range i.APARTMENTTYPES.APARTMENTTYPE {

			query = fmt.Sprintf(" ('%d',if(%t,'1', '0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
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

	if len(i.HOUSETYPES.HOUSETYPE) >= 1 {
		for _, s := range i.HOUSETYPES.HOUSETYPE {

			query = fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, s.DESC, s.SHORTNAME, s.NAME, 0, s.Text)
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
	if len(i.OBJECTLEVELS.OBJECTLEVEL) >= 1 {
		for _, s := range i.OBJECTLEVELS.OBJECTLEVEL {

			query = fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%s','%s')", 0, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", "", s.NAME, s.LEVEL, s.Text)
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

	if len(i.OPERATIONTYPES.OPERATIONTYPE) >= 1 {
		for _, s := range i.OPERATIONTYPES.OPERATIONTYPE {

			query = fmt.Sprintf(" ('%d',if(%t,'1', '0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", s.SHORTNAME, s.NAME, 0, s.Text)
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

	if len(i.ROOMTYPES.ROOMTYPE) >= 1 {
		for _, s := range i.ROOMTYPES.ROOMTYPE {

			query = fmt.Sprintf(" ('%d',IF(%t,'1','0'),'%s','%s','%s','%s','%s','%s','%d','%s')", s.ID, s.ISACTIVE, s.ENDDATE, s.STARTDATE, s.UPDATEDATE, "", s.SHORTNAME, s.NAME, 0, s.Text)
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

func (i2 InsertData) Params(tableName string, data model.PARAMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `change_id`, `change_id_end`, `type_id`, `value`, `update_date`, `start_date`, `end_date`) VALUES ", tableName)
	var query, comma string

	for i, s := range data.PARAM {

		strings.ReplaceAll(strings.ReplaceAll(s.VALUE, "ж\\", ":"), "\\", "")

		query = fmt.Sprintf("%s%s ('%d','%s','%s','%s','%d','%s','%s','%s','%s')", query, comma, s.ID, s.OBJECTID, s.CHANGEID, s.CHANGEIDEND, s.TYPEID, s.VALUE, s.UPDATEDATE, s.STARTDATE, s.ENDDATE)
		if (i != 0 && i%200 == 0) || i == len(data.PARAM)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}
	return true
}

func (i2 InsertData) AdmHierarchy(tableName string, data model_hierarchy.ADMITEMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("insert INTO `%s`(`nalog_id`,  `object_id`, `parent_object_id`, `change_id`, `region_code`, `area_code`, `city_code`,`place_code`, `plan_code`,`street_code`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `path`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.ITEM {
		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%d','%d','%d','%d','%d','%d','%s','%s','%s','%s' ,'%s',IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.REGIONCODE, s.AREACODE, s.CITYCODE, s.PLACECODE, s.PLANCODE, s.STREETCODE, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.PATH)
		if (i != 0 && i%200 == 0) || i == len(data.ITEM)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Panicln(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}

	}

	return true
}

//func (i2 InsertData) AdmHierarchyTwo(tableName string, i model_hierarchy.ADMITEMSTwo) bool {
//	logger := logging.GetLogger()
//	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`PrimaryKey`, `ID`, `OBJECTID`, `PARENTOBJID`, `CHANGEID`, `REGIONCODE`, `AREACODE`, `CITYCODE`, `PLACECODE`, `PLANCODE`, `STREETCODE`, `PREVID`, `NEXTID`, `UPDATEDATE`, `STARTDATE`, `ENDDATE`, `ISACTIVE`, `PATH`, `FK_ITEMS`) VALUES ", "hierarchy_adm_2")
//	var query string
//	wg := sync.WaitGroup{}
//	db0, err := NewDB(i2.cfg)
//	if err != nil {
//		logger.Error(err)
//	}
//	db1, err := NewDB(i2.cfg)
//	if err != nil {
//		logger.Error(err)
//	}
//	db2, err := NewDB(i2.cfg)
//	if err != nil {
//		logger.Error(err)
//	}
//	db3, err := NewDB(i2.cfg)
//	if err != nil {
//		logger.Error(err)
//	}
//
//	for _, s := range i.ITEM {
//		go (s)
//
//	}
//	wg.Wait()
//
//	return true
//}

func (i2 InsertData) MunHierarchy(tableName string, data model_hierarchy.MUNITEMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `parent_object_gid`, `change_id`, `oktmo`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `path`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.ITEM {

		query = fmt.Sprintf("%s%s  ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',IF(%t,'1','0'),'%s','%s','%s')", query, comma, s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.OKTMO, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, "1", s.PATH, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.ITEM)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}

	}
	return true
}

func (i2 InsertData) Apartments(tableName string, data model_apartments.APARTMENTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `apart_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.APARTMENT {

		query = fmt.Sprintf("%s%s  ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.APARTTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.APARTMENT)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) Carplaces(tableName string, data model_carplaces.CARPLACES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query, comma string

	for i, s := range data.CARPLACE {

		query = fmt.Sprintf("%s%s  ('%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.CARPLACE)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) Houses(tableName string, data model_houses.HOUSES) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `house_num`, `house_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)

	var query, comma string

	for i, s := range data.HOUSE {

		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.HOUSENUM, s.HOUSETYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.HOUSE)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) AddrObject(tableName string, data model_objectAddr.ADDRESSOBJECTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `name`, `typename`, `level`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.OBJECT {
		query = fmt.Sprintf("%s%s ('%d','%d','%s','%d','%s','%s','%d','%d','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NAME, s.TYPENAME, s.LEVEL, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.OBJECT)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}

	}
	return true
}
func AddAddrObject(DB *sqlx.DB, tableName string, i model_objectAddr.ADDRESSOBJECTS) bool {
	logger := logging.GetLogger()
	//queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `name`, `typename`, `level`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query string
	for _, s := range i.OBJECT {
		query = fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `name`, `typename`, `level`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ('%d','%d','%s','%d','%s','%s','%d','%d','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", tableName, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NAME, s.TYPENAME, s.LEVEL, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		_, err := DB.Exec(query)
		if err != nil {
			logger.Info(query)
			logger.Error(err)
			return false
		}
		query = ""

	}
	return true
}

func (i2 InsertData) Rooms(tableName string, data model_rooms.ROOMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `room_type`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query, comma string

	for i, s := range data.ROOM {
		//if _ != 0 {
		//	query = query + ", "
		//}
		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.ROOMTYPE, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.ROOM)-1 {
			query = queryPre + query

			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}

	}
	return true
}

func (i2 InsertData) Steads(tableName string, data model_steads.STEADS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("insert INTO `%s`(`id`, `object_id`, `object_gid`, `change_id`, `number`, `oper_type_id`, `prev_id`, `next_id`, `update_date`, `start_date`, `end_date`, `is_active`, `is_actualve`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.STEAD {

		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%s','%s','%d','%d','%s','%s','%s',IF(%t,'1','0'),IF(%t,'1','0'),'%s')", query, comma, s.ID, s.OBJECTID, s.OBJECTGUID, s.CHANGEID, s.NUMBER, s.OPERTYPEID, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.ISACTUAL, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.STEAD)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}

	}

	return true
}

func (i2 InsertData) ObjectDivision(tableName string, data model_objectAddr.ITEMS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`object_id`, `parent_id`, `child_id`, `change_id`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.ITEM {

		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%s')", query, comma, s.ID, s.PARENTID, s.CHILDID, s.CHANGEID, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.ITEM)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) ChangeHistory(tableName string, data model_other.CHANGEHISTORY) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`( `object_id`, `adr_object_id`, `change_id`, `oper_type_id`, `change_date`, `ndoc_id`, `text`) VALUES ", tableName)
	var query, comma string

	for i, s := range data.ITEM {

		query = fmt.Sprintf("%s%s  ('%s','%s','%s','%s','%s','%d','%s')", query, comma, s.OBJECTID, s.ADROBJECTID, s.CHANGEID, s.OPERTYPEID, s.CHANGEDATE, s.NDOCID, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.ITEM)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) NormDocs(tableName string, data model_other.NORMDOCS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`id`, `name`, `number`, `type`, `date`, `update_date`, `k_`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.NORMDOC {

		query = fmt.Sprintf("%s%s  ('%s','%s','%s','%s','%s','%s','%s','%s')", query, comma, s.ID, s.NAME, s.NUMBER, s.TYPE, s.DATE, s.UPDATEDATE, s.KIND, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.NORMDOC)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func (i2 InsertData) ReestrObject(tableName string, data model_other.REESTROBJECTS) bool {
	logger := logging.GetLogger()
	queryPre := fmt.Sprintf("REPLACE INTO `%s`(`object_id`, `object_gid`, `change_id`, `level_id`, `update_date`, `create_date`, `text`) VALUES ", tableName)
	var query, comma string
	for i, s := range data.OBJECT {
		query = fmt.Sprintf("%s%s ('%s','%s','%s','%s','%s','%s','%s')", query, comma, s.OBJECTGUID, s.OBJECTGUID, s.CHANGEID, s.LEVELID, s.UPDATEDATE, s.CREATEDATE, s.Text)
		if (i != 0 && i%200 == 0) || i == len(data.OBJECT)-1 {
			query = queryPre + query
			_, err := i2.db.Exec(query)
			if err != nil {
				logger.Info(query)
				logger.Error(err)
				return false
			}
			comma = ""
			query = ""
		} else {
			comma = ","
		}
	}

	return true
}

func NewInsert(db *sqlx.DB, cfg *viper.Viper) *InsertData {
	return &InsertData{db: db, cfg: cfg}
}

//func addAction(s model_hierarchy.AdmNodeTwo, queryPre string, wg *sync.WaitGroup, ) {
//	wg.Add(1)
//	//dbCel := i2.db
//	//db := *dbCel
//	//db, err := NewDB(i2.cfg)
//	//if err != nil {
//	//	logger.Error(err)
//	//}
//	query := fmt.Sprintf(" ('0','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','1')", s.ID, s.OBJECTID, s.PARENTOBJID, s.CHANGEID, s.REGIONCODE, s.AREACODE, s.CITYCODE, s.PLACECODE, s.PLACECODE, s.STREETCODE, s.PREVID, s.NEXTID, s.UPDATEDATE, s.STARTDATE, s.ENDDATE, s.ISACTIVE, s.PATH)
//
//	query = queryPre + query
//	_, err = i2.db.Exec(query)
//	if err != nil {
//		logger.Info(query)
//		logger.Error(err)
//
//	}
//	//err = db.Close()
//	if err != nil {
//		logger.Error(err)
//	}
//
//	query = ""
//	wg.Done()
//}
