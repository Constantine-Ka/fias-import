package service

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
	"fias-import_byLondon/pkg/repository"
	"fias-import_byLondon/utills"
	"fias-import_byLondon/utills/config"
	"fias-import_byLondon/utills/logging"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"strings"
)

type FileServices struct {
	repo *repository.Repository
	VP   *viper.Viper
}

func (s *FileServices) Unpacking(path, filetype string) []string {
	//logger := logging.GetLogger()
	//cfg := config.GetConf()
	//log.Println("Начало распаковки")
	//zipFile, err := zip.OpenReader(path)
	//log.Println("Конец распаковки")
	//if err != nil {
	//	logger.Fatal(err.Error())
	//}
	//defer zipFile.Close()
	//log.Println(ListTable)

	filesInfo := getListFile(path)

	var result []string

	for _, f := range filesInfo {

		//if i == 6 {
		//break
		ListTable := s.repo.Global.ListTable()
		//result = append(result, f)
		//fileReader, err := f.Open()
		//if err != nil {
		//	logger.Fatal(err)
		//}

		if !strings.HasSuffix(f, ".bak") {

			s.fileHandler(f, filetype, ListTable)

		}

		//}
	}

	return result
}

func getListFile(dir string) []string {
	logger := logging.GetLogger()
	log.Println(dir)
	//log.Println(stat.IsDir())
	//dirList, err := sftpClient.ReadDir(dir)
	dirList, err := os.ReadDir(dir)
	if err != nil {
		logger.Error(err)
	}
	var dirArr []string
	for _, info := range dirList {
		if info.IsDir() {
			//tempdir, _ := sftpClient.ReadDir(dir + "/" + info.Name())
			tempdir, _ := os.ReadDir(dir + "/" + info.Name())
			for _, fileInfo := range tempdir {
				dirArr = append(dirArr, dir+"/"+info.Name()+"/"+fileInfo.Name())
			}
		} else {
			dirArr = append(dirArr, dir+"/"+info.Name())
		}
	}
	//log.Println(dirArr)
	return dirArr
}

func (s *FileServices) fileHandler(f, filetype string, ListTable []string) {
	logger := logging.GetLogger()
	cfg := config.GetConf()
	if regexp.MustCompile(`AS_ADDR_OBJ_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "addrobj") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if len(ListTable) == 0 || utills.IndexOf(ListTable, cfg.Tablename.Content.AddrObject) == -1 {
			s.repo.CreateTable.AddrObject(cfg.Tablename.Content.AddrObject, model_objectAddr.ADDRESSOBJECTS{})
		}

		data := ParserAddrObj(fileReader, cfg.Tablename.Content.AddrObject, s.repo)
		for _, d := range data.OBJECT {
			if d.LEVEL == 1 {
				log.Println(d.NAME)
			}
		}

		//data := ParserAddrObj(fileReader, cfg.Tablename.Content.AddrObject, s.repo)

		_ = os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_ADDR_OBJ_DIVISION_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "addrobjdiv") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.ObjectDivision) == -1 {
			s.repo.CreateTable.ObjectDivision(cfg.Tablename.Content.ObjectDivision, model_objectAddr.ITEMS{})
		}
		s.repo.Inserter.ObjectDivision(cfg.Tablename.Content.ObjectDivision, ParserAddrObjDivision(fileReader))
		_ = os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_ADDR_OBJ_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "addrobjp") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.AddrObjectP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.AddrObjectP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.AddrObjectP, s.repo)

		_ = os.Rename(f, f+".bak")
		//s.repo.Inserter.Params(cfg.Tablename.Content.AddrObjectP, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_ADM_HIERARCHY_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "hierarchyadm") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.AdmHierarchy) == -1 {
			s.repo.CreateTable.AdmHierarchy(cfg.Tablename.Content.AdmHierarchy, model_hierarchy.ADMITEMS{})
		}
		//ParserUniversal(fileReader, cfg.Tablename.Content.AdmHierarchy, utills.XmlElemToAdmHieRarchy(), s.repo.Inserter.AdmHierarchy(tablename, result))
		ParserAdmHieRarchyTWO(fileReader, cfg.Tablename.Content.AdmHierarchy, s.repo, cfg)
		err = os.Rename(f, f+".bak")
		if err != nil {
			log.Print(err)
		}
		//s.repo.Inserter.AdmHierarchy(cfg.Tablename.Content.AdmHierarchy, ParserAdmHieRarchy(fileReader))
	} else if regexp.MustCompile(`AS_APARTMENTS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "appartments") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.Apartments) == -1 {
			s.repo.CreateTable.Apartments(cfg.Tablename.Content.Apartments, model_apartments.APARTMENTS{})
		}
		ParserApartments(fileReader, cfg.Tablename.Content.Apartments, s.repo)
		//s.repo.Inserter.Apartments(cfg.Tablename.Content.Apartments, ParserApartments(fileReader))
		_ = os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_APARTMENTS_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "appartmentsp") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.ApartmentsP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.ApartmentsP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.ApartmentsP, s.repo)

		_ = os.Rename(f, f+".bak")
		//s.repo.Inserter.Params(cfg.Tablename.Content.ApartmentsP, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_CARPLACES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "carplaces") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.Carplaces) == -1 {
			s.repo.CreateTable.Carplaces(cfg.Tablename.Content.Carplaces, model_carplaces.CARPLACES{})
		}
		s.repo.Inserter.Carplaces(cfg.Tablename.Content.Carplaces, ParserCarplaces(fileReader))
	} else if regexp.MustCompile(`AS_CARPLACES_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "carplacesp") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.CarplacesP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.CarplacesP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.CarplacesP, s.repo)
		os.Rename(f, f+".bak")

		//s.repo.Inserter.Params(, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_CHANGE_HISTORY_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "history") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.ChangeHistory) == -1 {
			s.repo.CreateTable.ChangeHistory(cfg.Tablename.Content.ChangeHistory, model_other.CHANGEHISTORY{})
		}

		ParserChangeH(fileReader, cfg.Tablename.Content.ChangeHistory, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.ChangeHistory(cfg.Tablename.Content.ChangeHistory, ParserChangeH(fileReader))

	} else if regexp.MustCompile(`AS_HOUSES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "houses") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.Houses) == -1 {
			s.repo.CreateTable.Houses(cfg.Tablename.Content.Houses, model_houses.HOUSES{})
		}
		ParserHouses(fileReader, cfg.Tablename.Content.Houses, s.repo)
		//s.repo.Inserter.Houses(cfg.Tablename.Content.Houses, ParserHouses(fileReader))
		_ = os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_HOUSES_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "housesp") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.HousesP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.HousesP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.HousesP, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.Params(, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_MUN_HIERARCHY_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "hierarchymun") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.MunHierarchy) == -1 {
			s.repo.CreateTable.MunHierarchy(cfg.Tablename.Content.MunHierarchy, model_hierarchy.MUNITEMS{})
		}
		parserMunHieRarchy(fileReader, cfg.Tablename.Content.MunHierarchy, s.repo)
		//s.repo.Inserter.MunHierarchy(cfg.Tablename.Content.MunHierarchy, parserMunHieRarchy(fileReader))
		err = os.Rename(f, f+".bak")
		if err != nil {
			log.Print(err)
		}
	} else if regexp.MustCompile(`AS_NORMATIVE_DOCS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "normativedocs") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.NormDocs) == -1 {
			s.repo.CreateTable.NormDocs(cfg.Tablename.Content.NormDocs, model_other.NORMDOCS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.NormDocs, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.NormDocs(, ParserNDocs(fileReader))
	} else if regexp.MustCompile(`AS_REESTR_OBJECTS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "reestrobj") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.ReestrObject) == -1 {
			s.repo.CreateTable.ReestrObject(cfg.Tablename.Content.ReestrObject, model_other.REESTROBJECTS{})
			s.repo.Global.SetPrimaryKey(cfg.Tablename.Content.ReestrObject, "id")
		}
		ParserReestrObj(fileReader, cfg.Tablename.Content.ReestrObject, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.ReestrObject(cfg.Tablename.Content.ReestrObject, ParserReestrObj(fileReader))
	} else if regexp.MustCompile(`AS_ROOMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "rooms") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.Rooms) == -1 {
			s.repo.CreateTable.Rooms(cfg.Tablename.Content.Rooms, model_rooms.ROOMS{})
		}
		s.repo.Inserter.Rooms(cfg.Tablename.Content.Rooms, ParserRooms(fileReader))
		os.Rename(f, f+".bak")

	} else if regexp.MustCompile(`AS_ROOMS_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "roomsp") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.RoomsP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.RoomsP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.RoomsP, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.Params(, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_STEADS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "steads") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.Steads) == -1 {
			s.repo.CreateTable.Steads(cfg.Tablename.Content.Steads, model_steads.STEADS{})
		}
		s.repo.Inserter.Steads(cfg.Tablename.Content.Steads, ParserSteads(fileReader))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_STEADS_PARAMS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "steads") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Content.SteadsP) == -1 {
			s.repo.CreateTable.Params(cfg.Tablename.Content.SteadsP, model.PARAMS{})
		}
		ParserParams(fileReader, cfg.Tablename.Content.SteadsP, s.repo)
		os.Rename(f, f+".bak")
		//s.repo.Inserter.Params(cfg.Tablename.Content.SteadsP, ParserParams(fileReader))
	} else if regexp.MustCompile(`AS_ADDHOUSE_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.AddHouseTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.AddHouseTypes, model.DICTALL{})
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.AddHouseTypes, parserObjectType(fileReader, "ADDHOUSE"))
	} else if regexp.MustCompile(`AS_ADDR_OBJ_TYPES_\d{8}`).MatchString(f) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.AddressObjectTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.AddressObjectTypes, model.DICTALL{})
		}

		//log.Println(parserObjectType(fileReader, "ADDR_OBJ"))
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.AddressObjectTypes, parserObjectType(fileReader, "ADDR_OBJ"))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_APARTMENT_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.ApartmentTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.ApartmentTypes, model.DICTALL{})
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.ApartmentTypes, parserObjectType(fileReader, "APARTMENT"))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_HOUSE_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.HouseTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.HouseTypes, model.DICTALL{})
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.HouseTypes, parserObjectType(fileReader, "HOUSE"))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_NORMATIVE_DOCS_KINDS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.NDockInds) == -1 {
			s.repo.CreateTable.Ndockinds(cfg.Tablename.Dict.NDockInds, model.NDOCKINDS{})
		}
		s.repo.Inserter.Ndockinds(cfg.Tablename.Dict.NDockInds, ParserNdocsInd(fileReader))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_NORMATIVE_DOCS_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.NDocTypes) == -1 {
			s.repo.CreateTable.Ndoctypes(cfg.Tablename.Dict.NDocTypes, model.NDOCTYPES{})
		}
		s.repo.Inserter.Ndoctypes(cfg.Tablename.Dict.NDocTypes, ParserNdocsType(fileReader))

		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_OBJECT_LEVELS_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.LevelsTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.LevelsTypes, model.DICTALL{})
			s.repo.Global.SetPrimaryKey(cfg.Tablename.Dict.LevelsTypes, "id")
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.LevelsTypes, parserObjectType(fileReader, "LEVELS"))
	} else if regexp.MustCompile(`AS_OPERATION_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.OperationTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.OperationTypes, model.DICTALL{})
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.OperationTypes, parserObjectType(fileReader, "OPERATION"))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_PARAM_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.ParamTypes) == -1 {
			s.repo.CreateTable.Paramtypes(cfg.Tablename.Dict.ParamTypes, model.PARAMTYPES{})
		}
		s.repo.Inserter.Paramtypes(cfg.Tablename.Dict.ParamTypes, ParserParamTypes(fileReader))
		os.Rename(f, f+".bak")
	} else if regexp.MustCompile(`AS_ROOM_TYPES_\d{8}`).Match([]byte(f)) && (filetype == "all" || filetype == "dict") {
		log.Println("Обрабатывается файл:", f)
		//fileReader, err := sftpC.Open(f)
		log.Println("Открыт файл: ", f)
		fileReader, err := os.Open(f)
		if err != nil {
			logger.Error(err)
		}
		if utills.IndexOf(ListTable, cfg.Tablename.Dict.RoomTypes) == -1 {
			s.repo.CreateTable.Addressobjecttypes(cfg.Tablename.Dict.RoomTypes, model.DICTALL{})
		}
		s.repo.Inserter.Addressobjecttypes(cfg.Tablename.Dict.RoomTypes, parserObjectType(fileReader, "ROOM"))
		os.Rename(f, f+".bak")
	}
}
func NewFileService(repo *repository.Repository, vp *viper.Viper) *FileServices {
	return &FileServices{repo: repo, VP: vp}
}
