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
type ITEMS struct {
	XMLName xml.Name `xml:"ITEMS"`
	Text    string   `xml:",chardata"`
	ITEM    []any    `xml:"ITEM"`
}

func ParserUniversal(fileReader *os.File, tablename string, adapterFunc func(element xml.StartElement) any, insertFunc func(string, any)) {
	logger := logging.GetLogger()
	var result ITEMS
	log.Println("Начато открытие файл")
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult interface{}
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.ITEM) != 0 {
					insertFunc(tablename, result)
					//result = model_hierarchy.ADMITEMS{}
				}
				break
			} else {
				log.Print(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "ITEM" {
				tmpResult = adapterFunc(element)
				//tmpResult = utills.XmlElemToAdmHieRarchy(element)
				result.ITEM = append(result.ITEM, tmpResult)
				//result_test.ITEM = append(result_test.ITEM, tmpResult)
				if count == 1000 {
					//go func() {
					//	r.Inserter.AdmHierarchy(tablename, result)
					//}()
					insertFunc(tablename, result)
					result.ITEM = []any{}

					count = 0
				}
			}

		}

	}
	//logger.Println("😃😃😃")
	err := fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	//log.Println("Количество элементов:", )

	//for _, node := range result.ITEM {
	//
	//}

	log.Println("Файл закрыт")
	//return result

}
func ParserParams(fileReader *os.File, tablename string, r *repository.Repository) model.PARAMS {
	logger := logging.GetLogger()
	var result model.PARAMS
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		//var tmpResult model.ParamNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.PARAM) != 0 {
					r.Inserter.Params(tablename, result)
					result = model.PARAMS{}
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
				result.PARAM = append(result.PARAM, tmpResult)

				if count == 1000 {
					r.Inserter.Params(tablename, result)
					result = model.PARAMS{}
					count = 0
				}
			}

		}

	}
	logger.Info("Файл закрыт")
	return result
}
func ParserAddrObj(fileReader *os.File, tablename string, r *repository.Repository) model_objectAddr.ADDRESSOBJECTS {
	logger := logging.GetLogger()
	var a model_objectAddr.ADDRESSOBJECTS
	var result []model_objectAddr.AddrO
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result) != 0 {
					a = model_objectAddr.ADDRESSOBJECTS{OBJECT: result}
					r.Inserter.AddrObject(tablename, a)
					result = []model_objectAddr.AddrO{}
				}
				break
			} else {
				logger.Error(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "OBJECT" {
				count++
				tmpResult := utills.XmlElemToObject(element)
				//if tmpResult.NEXTID == 0 {
				result = append(result, tmpResult)
				a = model_objectAddr.ADDRESSOBJECTS{OBJECT: result}
				r.Inserter.AddrObject(tablename, a)
				result = []model_objectAddr.AddrO{}
				count = 0
				//}
			}

		}

	}
	logger.Info("Файл закрыт")
	return a
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
func ParserAdmHieRarchy(fileReader *os.File, tablename string, r *repository.Repository) model_hierarchy.ADMITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.ADMITEMS
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult model_hierarchy.AdmNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.ITEM) != 0 {
					r.Inserter.AdmHierarchy(tablename, result)
					result = model_hierarchy.ADMITEMS{}
				}
				break
			} else {
				log.Print(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "ITEM" {
				count++
				tmpResult = utills.XmlElemToAdmHieRarchy(element)
				result.ITEM = append(result.ITEM, tmpResult)
				if count == 100 {
					r.Inserter.AdmHierarchy(tablename, result)
					result.ITEM = []model_hierarchy.AdmNode{}
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
	err := fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	log.Println("Файл закрыт")
	//err = xml.Unmarshal(contentBytes, &result)
	//if err != nil {
	//	if err != io.EOF {
	//		logger.Error(err)
	//	}
	//}
	return result
}

func ParserAdmHieRarchyTWO(fileReader *os.File, tablename string, r *repository.Repository, cfg *model.Config) model_hierarchy.ADMITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.ADMITEMS
	//var resultArr []model_hierarchy.ADMITEMS
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult model_hierarchy.AdmNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.ITEM) != 0 {
					r.Inserter.AdmHierarchy(tablename, result)
					result = model_hierarchy.ADMITEMS{}
				}
				break
			} else {
				log.Print(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "ITEM" {
				count++
				tmpResult = utills.XmlElemToAdmHieRarchy(element)
				result.ITEM = append(result.ITEM, tmpResult)
				if count == 1000 {
					r.Inserter.AdmHierarchy(tablename, result)
					result.ITEM = []model_hierarchy.AdmNode{}

					count = 0
				}
			}

		}

	}
	//logger.Println("😃😃😃")
	err := fileReader.Close()
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Файл закрыт")
	return result
}

func ParserApartments(fileReader *os.File, tablename string, r *repository.Repository) model_apartments.APARTMENTS {
	logger := logging.GetLogger()
	var result model_apartments.APARTMENTS
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult model_apartments.ITEM
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.APARTMENT) != 0 {
					r.Inserter.Apartments(tablename, result)
					result = model_apartments.APARTMENTS{}
				}
				break
			} else {
				logger.Error(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "APARTMENT" {
				count++
				tmpResult = utills.XmlElemToAppartments(element)
				result.APARTMENT = append(result.APARTMENT, tmpResult)
				if count == 1000 {
					r.Inserter.Apartments(tablename, result)
					result.APARTMENT = []model_apartments.ITEM{}
					count = 0
				}
			}

		}

	}
	logger.Info("Файл закрыт")
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
func ParserHouses(fileReader *os.File, tablename string, r *repository.Repository) model_houses.HOUSES {
	logger := logging.GetLogger()
	var a model_houses.HOUSES
	var result []model_houses.HousesItem
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		//var tmpResult model.ParamNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result) != 0 {
					a = model_houses.HOUSES{HOUSE: result}
					r.Inserter.Houses(tablename, a)
					result = []model_houses.HousesItem{}
				}
				break
			} else {
				log.Print(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "HOUSE" {
				count++
				tmpResult := utills.XmlElemToHouses(element)
				//if tmpResult.NEXTID == 0 {
				result = append(result, tmpResult)
				if count == 1000 {
					r.Inserter.Houses(tablename, a)
					result = []model_houses.HousesItem{}
					count = 0
				}
				a = model_houses.HOUSES{HOUSE: result}

				//}
			}

		}

	}
	logger.Info("Файл закрыт")
	return a
}
func parserMunHieRarchy(fileReader *os.File, tablename string, r *repository.Repository) model_hierarchy.MUNITEMS {
	logger := logging.GetLogger()
	var a model_hierarchy.MUNITEMS
	var result []model_hierarchy.MunNode
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		//var tmpResult model.ParamNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result) != 0 {
					a = model_hierarchy.MUNITEMS{ITEM: result}
					r.Inserter.MunHierarchy(tablename, a)
					result = []model_hierarchy.MunNode{}
				}
				break
			} else {
				logger.Error(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "ITEM" {
				count++
				tmpResult := utills.XmlElemToMunHieRarchy(element)
				//if tmpResult.NEXTID == "0" {
				result = append(result, tmpResult)
				if count == 1000 {
					a = model_hierarchy.MUNITEMS{ITEM: result}
					r.Inserter.MunHierarchy(tablename, a)
					result = []model_hierarchy.MunNode{}
					count = 0
				}
			}
		}
	}
	logger.Info("Файл закрыт")
	return a
}
func ParserReestrObj(fileReader *os.File, tableName string, r *repository.Repository) model_other.REESTROBJECTS {
	logger := logging.GetLogger()
	var result model_other.REESTROBJECTS
	count := 0
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult model_other.ReestrNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.OBJECT) != 0 {
					r.Inserter.ReestrObject(tableName, result)
					result = model_other.REESTROBJECTS{}
				}
				break
			} else {
				logger.Error(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "OBJECT" {
				count++
				tmpResult = utills.XmlElemToReestr(element)
				result.OBJECT = append(result.OBJECT, tmpResult)
				if count == 1000 {
					r.Inserter.ReestrObject(tableName, result)
					result.OBJECT = []model_other.ReestrNode{}
					count = 0
				}
			}

		}

	}
	logger.Info("Файл закрыт")
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

	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	logger.Info("Файл закрыт")
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

	err = xml.Unmarshal(contentBytes, &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	logger.Info("Файл закрыт")
	return result
}
func ParserChangeH(fileReader *os.File, tablename string, r *repository.Repository) model_other.CHANGEHISTORY {
	logger := logging.GetLogger()
	var result model_other.CHANGEHISTORY
	count := 0
	//byteFile := StreamToString(fileReader)
	decoder := xml.NewDecoder(fileReader)
	for {
		var tmpResult model_other.HistoryNode
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				//TODO
				if len(result.ITEM) != 0 {
					r.Inserter.ChangeHistory(tablename, result)
					result = model_other.CHANGEHISTORY{}
				}
				break
			} else {
				logger.Error(err)
				break
			}
		}
		if element, ok := token.(xml.StartElement); ok {
			if element.Name.Local == "PARAM" {
				count++
				tmpResult = utills.XmlElemToHistory(element)
				result.ITEM = append(result.ITEM, tmpResult)
				if count == 1000 {
					r.Inserter.ChangeHistory(tablename, result)
					result = model_other.CHANGEHISTORY{}
					count = 0
				}
			}

		}

	}
	logger.Info("Файл закрыт")
	//}
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
	err = fileReader.Close()
	if err != nil {
		logger.Error(err)
	}
	var result model.DICTALL
	var ADDRESSOBJECTTYPE model.ADDRESSOBJECTTYPES
	switch t {
	case "ADDHOUSE":
		err = xml.Unmarshal(contentBytes, &result.ADDHOUSETYPES)
	case "ADDR_OBJ":
		log.Println("😦")
		log.Println(string(contentBytes))
		err = xml.Unmarshal(contentBytes, &ADDRESSOBJECTTYPE)
		result.ADDRESSOBJECTTYPES = ADDRESSOBJECTTYPE
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
	//log.Println(ADDRESSOBJECTTYPE)
	logger.Info("Файл закрыт")
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
