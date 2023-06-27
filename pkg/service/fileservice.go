package service

import (
	"archive/zip"
	"fias-import_byLondon/pkg/repository"
	"fias-import_byLondon/utills/logging"
	"log"
	"regexp"
)

type FileServices struct {
	repo repository.Inserter
}

func (s FileServices) Unpacking(path string) []string {
	logger := logging.GetLogger()
	log.Println("Начало распаковки")
	zipFile, err := zip.OpenReader(path)
	log.Println("Конец распаковки")
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer zipFile.Close()
	var result []string
	for i, f := range zipFile.File {
		if i == 106 {

			result = append(result, f.Name)
			log.Println(f.Name)
			fileReader, err := f.Open()
			if err != nil {
				logger.Fatal(err)
			}
			regexFilename, err := regexp.MatchString(`AS_ADDR_OBJ_\d{8}`, f.Name)
			if err != nil {
				logger.Fatal(err)
			}
			if regexFilename {
				log.Println("😎")
				AddRObj := s.ParserAddrObj(fileReader)
				log.Println(len(AddRObj.OBJECT))
				//log.Println(i)
				for j, objects := range AddRObj.OBJECT {
					if j == 0 {
						log.Println("😛")
						log.Print(objects)
					}

					//s.repo.InsertObjAddr("fias_object", objects)
				}
			}
		}
	}
	//regexFilename, err = regexp.MatchString(`AS_ADDR_OBJ_DIVISION_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	AddRObjDivison := ParserAddrObjDivision(fileReader)
	//}
	//
	//regexFilename, err = regexp.MatchString(`AS_ADM_HIERARCHY_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	AddRAdmHieRarchy := ParserAdmHieRarchy(fileReader)
	//}
	//regexFilename, err = regexp.MatchString(`AS_APARTMENTS_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	AddRApartments := ParserApartments(fileReader)
	//}
	//
	//
	//regexFilename, err = regexp.MatchString(`AS_CARPLACES_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	carplaces := ParserCarplaces(fileReader)
	//}
	//
	//
	//regexFilename, err = regexp.MatchString(`AS_HOUSES_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	houses := ParserHouses(fileReader)
	//}
	//
	//
	//regexFilename, err = regexp.MatchString(`AS_MUN_HIERARCHY_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	MunHieRarchy := ParserMunHieRarchy(fileReader)
	//}
	////regexFilename, err = regexp.MatchString(`AS_NORMATIVE_DOCS_\d{8}`, f.Name)
	////if err != nil {
	////	logger.Fatal(err)
	////}
	//
	//regexFilename, err = regexp.MatchString(`AS_REESTR_OBJECTS_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//if regexFilename {
	//	ReestrObj := ParserReestrObj(fileReader)
	//}
	//regexFilename, err = regexp.MatchString(`AS_ROOMS_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//if regexFilename {
	//	rooms := ParserRooms(fileReader)
	//}
	//
	//regexFilename, err = regexp.MatchString(`AS_STEADS_\d{8}`, f.Name)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//if regexFilename {
	//	steads := ParserSteads(fileReader)
	//}
	//
	//if strings.Contains(f.Name, "ADDR_OBJ_") {
	//	log.Println(f.Name)
	//	//addrObjects := s.ParserAddrObj(fileReader)
	//}
	//}

	return result
}

func NewFileService(repo repository.Inserter) *FileServices {
	return &FileServices{repo: repo}
}
