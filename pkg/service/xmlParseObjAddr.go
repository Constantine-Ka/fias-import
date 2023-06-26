package service

import (
	"bytes"
	"encoding/xml"
	"fias-import_byLondon/model"
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
func (s FileServices) ParserAddrObj(fileReader io.Reader) model.ADDRESSOBJECTS {
	logger := logging.GetLogger()
	var result model.ADDRESSOBJECTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	log.Print(len(result.ObjectField))
	return result
}

func (s FileServices) ParserAddrObjDivision(fileReader io.Reader) model.Devisions {
	logger := logging.GetLogger()
	var result model.Devisions
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserAdmHieRarchy(fileReader io.Reader) model.HIERARCHYItem {
	logger := logging.GetLogger()
	var result model.HIERARCHYItem
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserApartments(fileReader io.Reader) model.APARTMENTS {
	logger := logging.GetLogger()
	var result model.APARTMENTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}

func (s FileServices) ParserCarplaces(fileReader io.Reader) model.CARPLACES {
	logger := logging.GetLogger()
	var result model.CARPLACES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserHouses(fileReader io.Reader) model.HOUSES {
	logger := logging.GetLogger()
	var result model.HOUSES
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) arserMunHieRarchy(fileReader io.Reader) model.ITEMSMUN_HIERARCHY {
	logger := logging.GetLogger()
	var result model.ITEMSMUN_HIERARCHY
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserReestrObj(fileReader io.Reader) model.REESTROBJECTS {
	logger := logging.GetLogger()
	var result model.REESTROBJECTS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserRooms(fileReader io.Reader) model.ROOMS {
	logger := logging.GetLogger()
	var result model.ROOMS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
func (s FileServices) ParserSteads(fileReader io.Reader) model.STEADS {
	logger := logging.GetLogger()
	var result model.STEADS
	byteFile := StreamToString(fileReader)
	err := xml.Unmarshal([]byte(byteFile), &result)
	if err != nil {
		if err != io.EOF {
			logger.Error(err)
		}
	}
	return result
}
