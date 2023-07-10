package main

import (
	"fias-import_byLondon/model"
	"fias-import_byLondon/pkg/handler"
	"fias-import_byLondon/pkg/repository"
	"fias-import_byLondon/pkg/service"
	"fias-import_byLondon/utills/logging"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfg *model.Config

func main() {
	//--------------------------------------
	//Инициализация библиотек
	logger := logging.GetLogger()
	logger.Info("Подготовка")
	vp := viper.New()
	vp.AddConfigPath(".")
	//vp.SetConfigName("config.yaml")
	err := vp.ReadInConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("fatal error config file: %w",
			err))
	}
	sftpClient, sshClient := handler.CreateConnection(*vp)
	sftpClient.Close()
	sshClient.Close()
	//sftpClient.Wait()
	//config.GetConf()

	//-------------------------------`-------
	//Инициализация Флагов
	file := flag.String("file", vp.GetString("basic.filepath"), "Путь к файлу")
	prefix := flag.String("prefix", vp.GetString("basic.prefix"), "Что парсить")
	//---Варианты префиксов
	// @param dict
	// @param addrobj
	// @param addrobjdiv
	// @param addrobjp
	// @param hierarchyadm
	// @param hierarchymun
	// @param appartments
	// @param appartmentsp
	// @param carplaces
	// @param carplacesp
	// @param history
	// @param houses
	// @param housesp
	// @param normativedocs
	// @param reestrobj
	// @param rooms
	// @param roomsp
	// @param steads
	// @param steadsp

	//--------
	flag.Parse()
	//-------------------------------------
	//Инициализация Баз данных
	db, err := repository.NewDB(vp)
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(0)
	if err != nil {
		logger.Fatalf("failed to initialize db:%s",
			err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//handlers := handler.NewHandler(services, sshClient)
	//ospTablename := vp.GetString("tablename.osp")
	//--------------------------------------
	//Счётчики
	countAll := 0
	countValid := 0
	countInvalid := 0
	//----------------------------------------------------
	//Сама программа
	logger.Infoln("Запущено")

	//services.InstallServices.NewTables()
	_ = services.Unpacking(*file, *prefix, sftpClient, sshClient)
	sftpClient.Close()
	sshClient.Close()
	//for _, name := range names {
	//	log.Println(name)
	//}
	//----------------------------------------------------
	//Подведение итогов
	logger.Info("Всего: ", countAll)
	logger.Info("Успешно: ", countValid)
	logger.Info("Неудачно: ", countInvalid)
	logger.Info("Завершено")

	//----------------------------------------------------

}
