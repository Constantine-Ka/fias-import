package service

import (
	"bytes"
	"encoding/xml"
	model_apartments "fias-import_byLondon/model/model-apartments"
	model_carplaces "fias-import_byLondon/model/model-carplaces"
	model_hierarchy "fias-import_byLondon/model/model-hierarchy"
	model_houses "fias-import_byLondon/model/model-houses"
	model_objectAddr "fias-import_byLondon/model/model-objectAddr"
	model_other "fias-import_byLondon/model/model-other"
	model_rooms "fias-import_byLondon/model/model-rooms"
	model_steads "fias-import_byLondon/model/model-steads"
	"fias-import_byLondon/utills/logging"
	"io"
	"log"
)

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
func (s FileServices) ParserAddrObj(fileReader io.Reader) model_objectAddr.ADDRESSOBJECTS {
	logger := logging.GetLogger()
	var result model_objectAddr.ADDRESSOBJECTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	log.Print(len(result.OBJECT))
	return result
}

func (s FileServices) ParserAddrObjDivision(fileReader io.Reader) model_objectAddr.ITEMS {
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
func (s FileServices) ParserAdmHieRarchy(fileReader io.Reader) model_hierarchy.ADMITEMS {
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
func (s FileServices) ParserApartments(fileReader io.Reader) model_apartments.APARTMENTS {
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

func (s FileServices) ParserCarplaces(fileReader io.Reader) model_carplaces.CARPLACES {
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
func (s FileServices) ParserHouses(fileReader io.Reader) model_houses.HOUSES {
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
func (s FileServices) arserMunHieRarchy(fileReader io.Reader) model_hierarchy.MUNITEMS {
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
func (s FileServices) ParserReestrObj(fileReader io.Reader) model_other.REESTROBJECTS {
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
func (s FileServices) ParserRooms(fileReader io.Reader) model_rooms.ROOMS {
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
func (s FileServices) ParserSteads(fileReader io.Reader) model_steads.STEADS {
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
