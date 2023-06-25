package repository

import (
	"fias-import_byLondon/utills/logging"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(viper *viper.Viper) (*sqlx.DB, error) {

	cfg := Config{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	log.Println(cfg.DBName)
	logger := logging.GetLogger()
	var err error
	var db *sqlx.DB
	if cfg.Driver == "postgres" {
		db, err = sqlx.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	} else {
		db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	}
	if err != nil {
		logger.Panicln("Ошибка подключения к БД. ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Panicln("Ошибка при пинге БД. ", err)
		return nil, err
	}
	return db, nil
}
