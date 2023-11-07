package utills

import (
	"encoding/xml"
	"fias-import_byLondon/model"
	model_apartments "fias-import_byLondon/model/model-apartments"
	model_hierarchy "fias-import_byLondon/model/model-hierarchy"
	model_houses "fias-import_byLondon/model/model-houses"
	model_objectAddr "fias-import_byLondon/model/model-objectAddr"
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

func XmlElemToAdmHieRarchyTwo(attrS xml.StartElement) model_hierarchy.AdmNodeTwo {
	var tmpResult model_hierarchy.AdmNodeTwo
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
		case "AREACODE":
			tmpResult.AREACODE = attr.Value
		case "CITYCODE":
			tmpResult.CITYCODE = attr.Value
		case "PLACECODE":
			tmpResult.PLACECODE = attr.Value
		case "PLANCODE":
			tmpResult.PLACECODE = attr.Value
		case "STREETCODE":
			tmpResult.PLACECODE = attr.Value
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
				tmpResult.ISACTIVE = "1"
			} else {
				tmpResult.ISACTIVE = "0"
			}
		case "PATH":
			tmpResult.PATH = attr.Value
		}
	}
	return tmpResult
}

func XmlElemToMunHieRarchy(attrS xml.StartElement) model_hierarchy.MunNode {
	var tmpResult model_hierarchy.MunNode
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
		case "OKTMO":
			tmpResult.OKTMO = attr.Value
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
func XmlElemToObject(attrS xml.StartElement) model_objectAddr.AddrO {
	var tmpResult model_objectAddr.AddrO
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "ID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.ID = num
		case "OBJECTID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.OBJECTID = num
		case "OBJECTGUID":
			tmpResult.OBJECTGUID = attr.Value
		case "CHANGEID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.CHANGEID = num
		case "NAME":
			tmpResult.NAME = attr.Value
		case "TYPENAME":
			tmpResult.TYPENAME = attr.Value
		case "LEVEL":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.LEVEL = num
		case "OPERTYPEID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.OPERTYPEID = num
		case "PREVID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.PREVID = num
		case "NEXTID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_objectAddr.AddrO{}
			}
			tmpResult.NEXTID = num
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		case "STARTDATE":
			tmpResult.STARTDATE = attr.Value
		case "ENDDATE":
			tmpResult.ENDDATE = attr.Value
		case "ISACTUAL":
			tmpResult.ISACTUAL = attr.Value == "1"
		case "ISACTIVE":
			tmpResult.ISACTIVE = attr.Value == "1"
		}
	}
	return tmpResult
}
func XmlElemToHouses(attrS xml.StartElement) model_houses.HousesItem {
	var tmpResult model_houses.HousesItem
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "ID":
			tmpResult.ID = attr.Value
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "OBJECTGUID":
			tmpResult.OBJECTGUID = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "HOUSENUM":
			tmpResult.HOUSENUM = attr.Value
		case "HOUSETYPE":
			tmpResult.HOUSETYPE = attr.Value
		case "OPERTYPEID":
			tmpResult.OPERTYPEID = attr.Value
		case "PREVID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_houses.HousesItem{}
			}
			tmpResult.PREVID = num
		case "NEXTID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_houses.HousesItem{}
			}
			tmpResult.NEXTID = num
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		case "STARTDATE":
			tmpResult.STARTDATE = attr.Value
		case "ENDDATE":
			tmpResult.ENDDATE = attr.Value
		case "ISACTUAL":
			tmpResult.ISACTUAL = attr.Value == "1"
		case "ISACTIVE":
			tmpResult.ISACTIVE = attr.Value == "1"
		}
	}
	return tmpResult
}
func XmlElemToAppartments(attrS xml.StartElement) model_apartments.ITEM {
	var tmpResult model_apartments.ITEM
	for _, attr := range attrS.Attr {
		switch attr.Name.Local {
		case "ID":
			tmpResult.ID = attr.Value
		case "OBJECTID":
			tmpResult.OBJECTID = attr.Value
		case "OBJECTGUID":
			tmpResult.OBJECTGUID = attr.Value
		case "CHANGEID":
			tmpResult.CHANGEID = attr.Value
		case "NUMBER":
			tmpResult.NUMBER = attr.Value
		case "APARTTYPE":
			tmpResult.APARTTYPE = attr.Value
		case "OPERTYPEID":
			tmpResult.OPERTYPEID = attr.Value
		case "PREVID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_apartments.ITEM{}
			}
			tmpResult.PREVID = num
		case "NEXTID":
			num, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return model_apartments.ITEM{}
			}
			tmpResult.NEXTID = num
		case "UPDATEDATE":
			tmpResult.UPDATEDATE = attr.Value
		case "STARTDATE":
			tmpResult.STARTDATE = attr.Value
		case "ENDDATE":
			tmpResult.ENDDATE = attr.Value
		case "ISACTUAL":
			tmpResult.ISACTUAL = attr.Value == "1"
		case "ISACTIVE":
			tmpResult.ISACTIVE = attr.Value == "1"
		}
	}
	return tmpResult
}
