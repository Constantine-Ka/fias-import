package model_apartments

import "encoding/xml"

type APARTMENTS struct {
	XMLName   xml.Name `xml:"APARTMENTS"`
	Text      string   `xml:",chardata"`
	APARTMENT []ITEM   `xml:"APARTMENT"`
}
type ITEM struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"ID,attr"`
	OBJECTID   string `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   string `xml:"CHANGEID,attr"`
	NUMBER     string `xml:"NUMBER,attr"`
	APARTTYPE  string `xml:"APARTTYPE,attr"`
	OPERTYPEID string `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   bool   `xml:"ISACTUAL,attr"`
	ISACTIVE   bool   `xml:"ISACTIVE,attr"`
}
