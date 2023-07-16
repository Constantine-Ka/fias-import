package utills

import (
	"encoding/xml"
	"fias-import_byLondon/model"
	model_hierarchy "fias-import_byLondon/model/model-hierarchy"
	model_other "fias-import_byLondon/model/model-other"
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
func XmlElemToAdmHieRarchy(attrS xml.StartElement) model_hierarchy.AdmNode {
	var tmpResult model_hierarchy.AdmNode
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "ID":
			tmpResult.ID = attr.Value
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "PARENTOBJID":
			tmpResult.PARENTOBJID = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "REGIONCODE":
			tmpResult.REGIONCODE = attr.Value
		case "PREVID":
			tmpResult.PREVID = attr.Value
		case "NEXTID":
			tmpResult.NEXTID = attr.Value
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		case "STARTDATE":
			tmpResult.STARTDATE = attr.Value
		case "ENDDATE":
			tmpResult.ENDDATE = attr.Value
		case "ISACTIVE":
			if attr.Value == "1" {
				tmpResult.ISACTIVE = true
			} else {
				tmpResult.ISACTIVE = false
			}
		case "PATH":
			tmpResult.PATH = attr.Value
		}
	}
	return tmpResult
}
func XmlElemToHistory(attrS xml.StartElement) model_other.HistoryNode {
	var tmpResult model_other.HistoryNode
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "Text":
			tmpResult.Text = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "ADROBJECTID":
			tmpResult.ADROBJECTID = attr.Value
		case "OPERTYPEID":
			tmpResult.OPERTYPEID = attr.Value
		case "CHANGEDATE":
			tmpResult.CHANGEDATE = attr.Value
		case "NDOCID":
			tmpResult.NDOCID, _ = strconv.ParseInt(attr.Value, 10, 64)
		}
	}
	return tmpResult
}
func XmlElemToReestr(attrS xml.StartElement) model_other.ReestrNode {
	var tmpResult model_other.ReestrNode
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "Text":
			tmpResult.Text = attr.Value
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "OBJECTGUID":
			tmpResult.OBJECTGUID = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "ISACTIVE":
			tmpResult.ISACTIVE = attr.Value == "1"
		case "LEVELID":
			tmpResult.LEVELID = attr.Value
		case "CREATEDATE":
			tmpResult.CREATEDATE = attr.Value
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		}
	}
	return tmpResult
}
