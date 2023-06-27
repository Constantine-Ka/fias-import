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

type CreatorTable struct {
	db *sqlx.DB
}

func (db *CreatorTable) AddrObject(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(    `id` INT NOT NULL,    `object_id` VARCHAR(255) NULL DEFAULT NULL,    `object_gid` VARCHAR(255) NULL DEFAULT NULL,    `change_id` INT(255) NULL DEFAULT NULL,    `name` VARCHAR(20) NULL DEFAULT NULL,    `typename` VARCHAR(20) NULL DEFAULT NULL,    `level` VARCHAR(20) NULL DEFAULT NULL,    `oper_type_id` INT(255) NULL DEFAULT NULL,    `prev_id` INT NULL DEFAULT NULL,    `next_id` INT NULL DEFAULT NULL,    `update_date` DATE NULL DEFAULT NULL,    `start_date` DATE NULL DEFAULT NULL,    `end_date` DATE NULL DEFAULT NULL,    `is_active` INT NULL DEFAULT '1',    `is_actualve` INT NULL DEFAULT '1',    `text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) MunHierarchy(tableName string, i model_hierarchy.MUNITEMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`parent_object_gid` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`oktmo` VARCHAR(20) NULL DEFAULT NULL,`prev_id` INT NULL DEFAULT NULL,`next_id` INT NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`start_date` DATE NULL DEFAULT NULL,`end_date` DATE NULL DEFAULT NULL,`is_active` INT NULL DEFAULT '1',`is_actualve` INT NULL DEFAULT '1',`path` VARCHAR(255) NULL DEFAULT NULL,`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) AdmHierarchy(tableName string, i model_hierarchy.ADMITEMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` VARCHAR(255) NULL DEFAULT NULL , `parent_object` VARCHAR(255) NULL DEFAULT NULL , `change_id` INT(255) NULL DEFAULT NULL , `region_code` VARCHAR(20) NULL DEFAULT NULL , `prev_id` INT NULL DEFAULT NULL , `next_id` INT NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `start_date` DATE NULL DEFAULT NULL , `end_date` DATE NULL DEFAULT NULL , `is_active` INT NULL DEFAULT '1' , `path` VARCHAR(255) NULL DEFAULT NULL , `text` VARCHAR(255) NULL DEFAULT NULL ) ENGINE = InnoDB;\n", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Apartments(tableName string, i model_apartments.APARTMENTS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(    `id` INT NOT NULL,    `object_id` VARCHAR(255) NULL DEFAULT NULL,    `object_gid` VARCHAR(255) NULL DEFAULT NULL,    `change_id` INT(255) NULL DEFAULT NULL,    `number` VARCHAR(20) NULL DEFAULT NULL,    `apart_type` VARCHAR(20) NULL DEFAULT NULL,    `oper_type_id` INT(255) NULL DEFAULT NULL,    `prev_id` INT NULL DEFAULT NULL,    `next_id` INT NULL DEFAULT NULL,    `update_date` DATE NULL DEFAULT NULL,    `start_date` DATE NULL DEFAULT NULL,    `end_date` DATE NULL DEFAULT NULL,    `is_active` INT NULL DEFAULT '1',    `is_actualve` INT NULL DEFAULT '1',    `text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Carplaces(tableName string, i model_carplaces.CARPLACES) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(    `id` INT NOT NULL,    `object_id` VARCHAR(255) NULL DEFAULT NULL,    `object_gid` VARCHAR(255) NULL DEFAULT NULL,    `change_id` INT(255) NULL DEFAULT NULL,    `number` VARCHAR(20) NULL DEFAULT NULL,    `oper_type_id` INT(255) NULL DEFAULT NULL,    `prev_id` INT NULL DEFAULT NULL,    `next_id` INT NULL DEFAULT NULL,    `update_date` DATE NULL DEFAULT NULL,    `start_date` DATE NULL DEFAULT NULL,    `end_date` DATE NULL DEFAULT NULL,    `is_active` INT NULL DEFAULT '1',    `is_actualve` INT NULL DEFAULT '1',    `text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Houses(tableName string, i model_houses.HOUSES) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`object_gid` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`number` VARCHAR(20) NULL DEFAULT NULL,`house_num` VARCHAR(20) NULL DEFAULT NULL,`house_type` VARCHAR(20) NULL DEFAULT NULL,`oper_type_id` INT(255) NULL DEFAULT NULL,`prev_id` INT NULL DEFAULT NULL,`next_id` INT NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`start_date` DATE NULL DEFAULT NULL,`end_date` DATE NULL DEFAULT NULL,`is_active` INT NULL DEFAULT '1',`is_actualve` INT NULL DEFAULT '1',`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) ObjectDivision(tableName string, i model_objectAddr.ITEMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`parent_id` VARCHAR(255) NULL DEFAULT NULL,`child_id` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Rooms(tableName string, i model_rooms.ROOMS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`object_gid` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`number` VARCHAR(20) NULL DEFAULT NULL,`room_type` VARCHAR(20) NULL DEFAULT NULL,`oper_type_id` INT(255) NULL DEFAULT NULL,`prev_id` INT NULL DEFAULT NULL,`next_id` INT NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`start_date` DATE NULL DEFAULT NULL,`end_date` DATE NULL DEFAULT NULL,`is_active` INT NULL DEFAULT '1',`is_actualve` INT NULL DEFAULT '1',`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Steads(tableName string, i model_steads.STEADS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`object_gid` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`number` VARCHAR(20) NULL DEFAULT NULL,`oper_type_id` INT(255) NULL DEFAULT NULL,`prev_id` INT NULL DEFAULT NULL,`next_id` INT NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`start_date` DATE NULL DEFAULT NULL,`end_date` DATE NULL DEFAULT NULL,`is_active` INT NULL DEFAULT '1',`is_actualve` INT NULL DEFAULT '1',`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) ChangeHistory(tableName string, i model_other.CHANGEHISTORY) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`adr_object_id` INT(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`oper_type_id` INT(20) NULL DEFAULT NULL,`change_date` DATE NULL DEFAULT NULL,`ndoc_id` INT(255) NULL DEFAULT NULL,`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) NormDocs(tableName string, i model_other.NORMDOCS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NULL,`name` VARCHAR(255) NULL DEFAULT NULL,`number` INT(255) NULL DEFAULT NULL,`type` VARCHAR(255) NULL DEFAULT NULL,`date` DATE NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`kind` VARCHAR(255) NULL DEFAULT NULL,`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) ReestrObject(tableName string, i model_other.REESTROBJECTS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s`(`id` INT NOT NULL,`object_id` VARCHAR(255) NULL DEFAULT NULL,`object_gid` VARCHAR(255) NULL DEFAULT NULL,`change_id` INT(255) NULL DEFAULT NULL,`level_id` INT(20) NULL DEFAULT NULL,`update_date` DATE NULL DEFAULT NULL,`create_date` DATE NULL DEFAULT NULL,`text` VARCHAR(255) NULL DEFAULT NULL) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

func (db *CreatorTable) Paramtypes(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `enddate` DATE NULL DEFAULT NULL , `startdate` DATE NULL DEFAULT NULL , `updatedate` DATE NULL DEFAULT NULL , `isactive` INT NOT NULL DEFAULT '1' , `code` TEXT NULL DEFAULT NULL , `description` VARCHAR(255) NOT NULL , `name` VARCHAR(255) NOT NULL , `text` VARCHAR(512) NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Ndoctypes(tableName string, i model.NDOCTYPES) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NULL DEFAULT NULL , `enddate` DATE NULL DEFAULT NULL , `startdate` DATE NULL DEFAULT NULL , `name` VARCHAR(255) NULL DEFAULT NULL , `text` VARCHAR(255) NULL DEFAULT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}
func (db *CreatorTable) Ndockinds(tableName string, i model.NDOCKINDS) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `name` VARCHAR(255) NULL DEFAULT NULL , `text` VARCHAR(255) NULL DEFAULT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

// ADDHOUSETYPES,ADDRESSOBJECTTYPES, APARTMENTTYPES, HOUSETYPES, OPERATIONTYPES, ROOMTYPES, OBJECTLEVELS

func (db *CreatorTable) Addressobjecttypes(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `isactive` INT NULL  DEFAULT NULL, `enddate` DATE NULL , `startdate` DATE NULL DEFAULT NULL , `updatedate` DATE NULL DEFAULT NULL , `description` VARCHAR(255)  NULL DEFAULT NULL , `shortname` VARCHAR(255) NULL DEFAULT NULL , `name` VARCHAR(512) NULL DEFAULT NULL , `level` INT NULL DEFAULT NULL , `text` VARCHAR(225) NULL DEFAULT NULL ) ENGINE = InnoDB;\n", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}

	return true
}

// General
func (db *CreatorTable) Params(tableName string, i interface{}) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` VARCHAR(255) NULL DEFAULT NULL , `change_id` VARCHAR(255) NULL DEFAULT NULL , `change_id_end` VARCHAR(255) NULL DEFAULT NULL , `type_id` INT NULL DEFAULT NULL , `value` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `start_date` DATE NULL DEFAULT NULL , `end_date` DATE NULL DEFAULT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func NewCreateTable(db *sqlx.DB) *CreatorTable {
	return &CreatorTable{db: db}
}
