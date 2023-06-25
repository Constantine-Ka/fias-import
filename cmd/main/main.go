package main

import (
	"fias-import_byLondon/pkg/service"
	"fias-import_byLondon/utills/logging"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//--------------------------------------
	//Инициализация библиотек
	logger := logging.GetLogger()
	logger.Info("Подготовка")
	vp := viper.New()
	vp.AddConfigPath(".") // optionally look for config in the working directory
	vp.SetConfigName("config.yaml")
	err := vp.ReadInConfig() // Find and read the config file
	if err != nil {          // Handle errors reading the config file
		logrus.Error(fmt.Errorf("fatal error config file: %w",
			err))
	}

	//-------------------------------`-------
	//Инициализация Баз данных
	//db, err := repository.NewDB(vp)
	//db.SetConnMaxLifetime(0)
	//db.SetMaxOpenConns(0)
	//if err != nil {
	//	logger.Fatalf("failed to initialize db:%s",
	//		err.Error())
	//}
	//repos := repository.NewRepository(db)
	//services := service.NewService(repos)
	//handlers := handler.NewHandler(services)
	//ospTablename := vp.GetString("tablename.osp")
	//--------------------------------------
	//Счётчики
	countAll := 0
	countValid := 0
	countInvalid := 0
	//----------------------------------------------------
	//Сама программа
	logger.Infoln("Запущено")
	names := service.Unpacking("gar_xml .zip", "./tmp")
	for _, name := range names {
		log.Println(name)
	}
	//----------------------------------------------------
	//Подведение итогов
	logger.Info("Всего: ", countAll)
	logger.Info("Успешно: ", countValid)
	logger.Info("Неудачно: ", countInvalid)
	logger.Info("Завершено")
	//----------------------------------------------------

}
