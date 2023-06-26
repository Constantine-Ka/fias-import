package repository

import (
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CreatorTable struct {
	db *sqlx.DB
}

func (db *CreatorTable) AdmHierarchy(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `parent_obj_id` INT NULL DEFAULT NULL , `region_code` INT NULL DEFAULT NULL , `update_data` DATE NULL DEFAULT NULL ) ENGINE = InnoDB", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Apartments(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `apart_number` VARCHAR(255) NULL DEFAULT NULL , `apart_type` INT NOT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

// fias_carplaces
func (db *CreatorTable) Carplaces(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `number` VARCHAR(255) NULL DEFAULT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Houses(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `house_number` VARCHAR(255) NULL DEFAULT NULL , `house_type` INT NOT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) MunHierarchy(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `oktmo` VARCHAR(255) NULL DEFAULT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Reestr(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `level_id` VARCHAR(255) NULL DEFAULT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Rooms(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `room_number` VARCHAR(255) NULL DEFAULT NULL , `room_type` INT NOT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) Steards(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NOT NULL , `object_id` INT NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL ,`steads_num` INT NOT NULL , `is_actual` INT NOT NULL , `is_active` INT NOT NULL ) ENGINE = InnoDB;", tableName)
	_, err := db.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
		return false
	}
	return true
}

func (db *CreatorTable) AddrObject(tableName string) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("CREATE TABLE `%s` (`id` INT NULL DEFAULT NULL , `object_id` VARCHAR(255) NULL DEFAULT NULL , `object_gid` VARCHAR(255) NULL DEFAULT NULL , `update_date` DATE NULL DEFAULT NULL , `name` VARCHAR(255) NULL DEFAULT NULL , `type_name` VARCHAR(255) NULL DEFAULT NULL , `level` INT NULL DEFAULT NULL , `is_actual` INT NULL DEFAULT NULL , `is_active` INT NULL DEFAULT NULL ) ENGINE = InnoDB;", tableName)
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
