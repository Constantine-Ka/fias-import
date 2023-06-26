package repository

import (
	"fias-import_byLondon/model"
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type InsertData struct {
	db *sqlx.DB
}

func (i InsertData) InsertObjAddr(tablename string, data model.Objects) bool {
	logger := logging.GetLogger()
	query := fmt.Sprintf("INSERT INTO `%s`(`id`, `object_id`, `object_gid`, `update_date`, `name`, `type_name`, `level`, `is_actual`, `is_active`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s')", tablename, data.ObjectField.ID, data.ObjectField.OBJECTID, data.ObjectField.OBJECTGUID, data.ObjectField.UPDATEDATE, data.ObjectField.NAME, data.ObjectField.TYPENAME, data.ObjectField.LEVEL, data.ObjectField.ISACTUAL, data.ObjectField.ISACTIVE)
	_, err := i.db.Exec(query)
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
