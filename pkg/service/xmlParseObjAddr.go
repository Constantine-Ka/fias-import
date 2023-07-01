package service

import (
	"bytes"
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
	"fias-import_byLondon/utills/logging"
	"io"
	"log"
)

type XMLServices struct {
	repo *repository.Repository
}

func streamToByte(stream io.Reader) []byte {
	logger := logging.GetLogger()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
func ParserParams(fileReader io.Reader) model.PARAMS {
	logger := logging.GetLogger()
	var result model.PARAMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserAddrObj(fileReader io.Reader) model_objectAddr.ADDRESSOBJECTS {
	logger := logging.GetLogger()
	var result model_objectAddr.ADDRESSOBJECTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}

func ParserAddrObjDivision(fileReader io.Reader) model_objectAddr.ITEMS {
	logger := logging.GetLogger()
	var result model_objectAddr.ITEMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserAdmHieRarchy(fileReader io.Reader) model_hierarchy.ADMITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.ADMITEMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserApartments(fileReader io.Reader) model_apartments.APARTMENTS {
	logger := logging.GetLogger()
	var result model_apartments.APARTMENTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}

func ParserCarplaces(fileReader io.Reader) model_carplaces.CARPLACES {
	logger := logging.GetLogger()
	var result model_carplaces.CARPLACES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserHouses(fileReader io.Reader) model_houses.HOUSES {
	logger := logging.GetLogger()
	var result model_houses.HOUSES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func parserMunHieRarchy(fileReader io.Reader) model_hierarchy.MUNITEMS {
	logger := logging.GetLogger()
	var result model_hierarchy.MUNITEMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserReestrObj(fileReader io.Reader) model_other.REESTROBJECTS {
	logger := logging.GetLogger()
	var result model_other.REESTROBJECTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserRooms(fileReader io.Reader) model_rooms.ROOMS {
	logger := logging.GetLogger()
	var result model_rooms.ROOMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserSteads(fileReader io.Reader) model_steads.STEADS {
	logger := logging.GetLogger()
	var result model_steads.STEADS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserChangeH(fileReader io.Reader) model_other.CHANGEHISTORY {
	logger := logging.GetLogger()
	var result model_other.CHANGEHISTORY
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserNDocs(fileReader io.Reader) model_other.NORMDOCS {
	logger := logging.GetLogger()
	var result model_other.NORMDOCS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func parserObjectType(fileReader io.Reader, t string) model.DICTALL {
	logger := logging.GetLogger()
	byteFile := StreamToString(fileReader)
	var result model.DICTALL
	var err error
	switch t {
	case "ADDHOUSE":
		err = xml.Unmarshal([]byte(byteFile), &result.ADDHOUSETYPES)
	case "ADDR_OBJ":
		log.Println("😦")
		log.Println(byteFile)
		err = xml.Unmarshal([]byte(byteFile), &result.ADDRESSOBJECTTYPE)
	case "APARTMENT":
		err = xml.Unmarshal([]byte(byteFile), &result.APARTMENTTYPES)
	case "HOUSE":
		err = xml.Unmarshal([]byte(byteFile), &result.HOUSETYPES)
	case "LEVELS":
		err = xml.Unmarshal([]byte(byteFile), &result.OBJECTLEVELS)
	case "OPERATION":
		err = xml.Unmarshal([]byte(byteFile), &result.OPERATIONTYPES)
	case "ROOM":
		err = xml.Unmarshal([]byte(byteFile), &result.ROOMTYPES)
	}
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserNdocsType(fileReader io.Reader) model.NDOCTYPES {
	logger := logging.GetLogger()
	var result model.NDOCTYPES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result

}
func ParserNdocsInd(fileReader io.Reader) model.NDOCKINDS {
	logger := logging.GetLogger()
	var result model.NDOCKINDS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func ParserParamTypes(fileReader io.Reader) model.PARAMTYPES {
	logger := logging.GetLogger()
	var result model.PARAMTYPES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
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
