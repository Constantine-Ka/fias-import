package utills

import (
	"encoding/xml"
	"fias-import_byLondon/model"
	"log"
	"strconv"
)

func IndexOf(arr []string, need string) int {
	if len(arr) == 0 {
		return -1
	}
	for ind, val := range arr {
		if val == need {
			return ind
		}
	}
	log.Println(000)
	return -1
}
func ComparisonInvalid(slice1, slice2 []string) []string {
	invalidElem := []string{}
	for _, s1 := range slice1 {
		if IndexOf(slice2, s1) == -1 {
			invalidElem = append(invalidElem, s1)
		}
	}
	return invalidElem

}

func XmlElemToParam(attrS xml.StartElement) model.ParamNode {
	var tmpResult model.ParamNode
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "TYPEID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model.ParamNode{}
			}
			tmpResult.TYPEID = num
		case "ID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model.ParamNode{}
			}
			tmpResult.ID = num
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "CHANGEIDEND":
			tmpResult.CHANGEIDEND = attr.Value
		case "VALUE":
			tmpResult.VALUE = attr.Value
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		case "STARTDATE":
			tmpResult.STARTDATE = attr.Value
		case "ENDDATE":
			tmpResult.ENDDATE = attr.Value
		}
	}
	return tmpResult
}
