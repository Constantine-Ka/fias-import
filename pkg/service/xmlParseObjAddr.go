package service

import (
	"encoding/xml"
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
	"fias-import_byLondon/utills/logging"
	"io"
	"log"
	"os"
)

type XMLServices struct {
	repo *repository.Repository
}

func ParserParams(fileReader *os.File, tablename string, r *repository.Repository) []model.ParamNode {
	//logger := logging.GetLogger()
	var result []model.ParamNode
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		//var tmpResult model.ParamNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result) != 0 {
					r.Inserter.ParamOne(tablename, result)
					result = []model.ParamNode{}
				}
				break
			} else {
				log.Print(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "PARAM" {
				count++
				tmpResult := utills.XmlElemToParam(element)
				switch tmpResult.TYPEID {
				case 5, 7, 10, 11, 12, 13, 19:
					result = append(result, tmpResult)
				}
				if count == 100 {
					r.Inserter.ParamOne(tablename, result)
					result = []model.ParamNode{}
					count = 0
				}
			}

		}

	}
	//contentBytes, err := io.ReadAll(fileReader)
	//if err != nil {
	//	logger.Error(err)
	//}
	//logger.Println("😃😃😃")
	//err = fileReader.Close()
	//if err != nil {
	//	logger.Error(err)
	//}
	//err = xml.Unmarshal(contentBytes, &result)
	//if err != nil {
	//	if err != io.EOF {
	//		logger.Error(err)
	//	}
	//}
	return result
}
func ParserAddrObj(fileReader *os.File) model_objectAddr.ADDRESSOBJECTS {
	logger := logging.GetLogger()
	var result model_objectAddr.ADDRESSOBJECTS
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}

func ParserAddrObjDivision(fileReader *os.File) model_objectAddr.ITEMS {
	logger := logging.GetLogger()
	var result model_objectAddr.ITEMS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserAdmHieRarchy(fileReader *os.File) model_hierarchy.ADMITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.ADMITEMS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserApartments(fileReader *os.File) model_apartments.APARTMENTS {
	logger := logging.GetLogger()
	var result model_apartments.APARTMENTS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}

func ParserCarplaces(fileReader *os.File) model_carplaces.CARPLACES {
	logger := logging.GetLogger()
	var result model_carplaces.CARPLACES
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserHouses(fileReader *os.File) model_houses.HOUSES {
	logger := logging.GetLogger()
	var result model_houses.HOUSES
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func parserMunHieRarchy(fileReader *os.File) model_hierarchy.MUNITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.MUNITEMS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserReestrObj(fileReader *os.File) model_other.REESTROBJECTS {
	logger := logging.GetLogger()
	var result model_other.REESTROBJECTS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserRooms(fileReader *os.File) model_rooms.ROOMS {
	logger := logging.GetLogger()
	var result model_rooms.ROOMS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserSteads(fileReader *os.File) model_steads.STEADS {
	logger := logging.GetLogger()
	var result model_steads.STEADS
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserChangeH(fileReader *os.File) model_other.CHANGEHISTORY {
	logger := logging.GetLogger()
	logger.Info("😊😊")
	var result model_other.CHANGEHISTORY
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserNDocs(fileReader *os.File) model_other.NORMDOCS {
	logger := logging.GetLogger()
	var result model_other.NORMDOCS
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func parserObjectType(fileReader *os.File, t string) model.DICTALL {
	logger := logging.GetLogger()
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	var result model.DICTALL
	switch t {
	case "ADDHOUSE":
		err = xml.Unmarshal(contentBytes, &result.ADDHOUSETYPES)
	case "ADDR_OBJ":
		log.Println("😦")
		log.Println(string(contentBytes))
		err = xml.Unmarshal(contentBytes, &result.ADDRESSOBJECTTYPE)
	case "APARTMENT":
		err = xml.Unmarshal(contentBytes, &result.APARTMENTTYPES)
	case "HOUSE":
		err = xml.Unmarshal(contentBytes, &result.HOUSETYPES)
	case "LEVELS":
		err = xml.Unmarshal(contentBytes, &result.OBJECTLEVELS)
	case "OPERATION":
		err = xml.Unmarshal(contentBytes, &result.OPERATIONTYPES)
	case "ROOM":
		err = xml.Unmarshal(contentBytes, &result.ROOMTYPES)
	}
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserNdocsType(fileReader *os.File) model.NDOCTYPES {
	logger := logging.GetLogger()
	var result model.NDOCTYPES
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result

}
func ParserNdocsInd(fileReader *os.File) model.NDOCKINDS {
	logger := logging.GetLogger()
	var result model.NDOCKINDS
	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserParamTypes(fileReader *os.File) model.PARAMTYPES {
	logger := logging.GetLogger()
	var result model.PARAMTYPES

	//byteFile := StreamToString(fileReader)
	contentBytes, err := io.ReadAll(fileReader)
	if err != nil {
		logger.Error(err)
	}
	logger.Println("😃😃😃")
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")

	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result

}

func NewXMLService(repo *repository.Repository) *XMLServices {
	return &XMLServices{repo: repo}
}
