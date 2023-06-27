package model_houses

import "encoding/xml"

type HOUSES struct {
	XMLName xml.Name `xml:"HOUSES"`
	Text    string   `xml:",chardata"`
	HOUSE   []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		HOUSENUM   string `xml:"HOUSENUM,attr"`
		HOUSETYPE  string `xml:"HOUSETYPE,attr"`
		OPERTYPEID string `xml:"OPERTYPEID,attr"`
		PREVID     string `xml:"PREVID,attr"`
		NEXTID     string `xml:"NEXTID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   string `xml:"ISACTUAL,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
	} `xml:"HOUSE"`
}
