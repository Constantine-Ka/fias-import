package repository

import (
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type GlobalSetting struct {
	db *sqlx.DB
}

func (g GlobalSetting) SetPrimaryKey(table, field string) {
	logger := logging.GetLogger()
	query := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` INT(11) NOT NULL AUTO_INCREMENT, add PRIMARY KEY (`%s`);", table, field, field, field)
	_, err := g.db.Exec(query)
	if err != nil {
		logger.Info(query)
		logger.Error(err)
	}
	//query = fmt.Sprintf("ALTER TABLE `%s` ADD UNIQUE(`%s`)", table, field)
	//_, err = g.db.Exec(query)
	//if err != nil {
	//	logger.Info("SHOW TABLES")
	//	logger.Error(err)
	//}
}

func (g GlobalSetting) ListTable() []string {

	logger := logging.GetLogger()
	query := fmt.Sprintf("SHOW TABLES")
	var result []string
	err := g.db.Select(&result, query)
	if err != nil {
		logger.Info("SHOW TABLES")
		logger.Error(err)
	}
	log.Println()
	return result
}

func NewGlobal(db *sqlx.DB) *GlobalSetting {
	return &GlobalSetting{db: db}
}
