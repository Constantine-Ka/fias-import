package model_carplaces

import "encoding/xml"

type CARPLACES struct {
	XMLName  xml.Name `xml:"CARPLACES"`
	Text     string   `xml:",chardata"`
	CARPLACE []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		NUMBER     string `xml:"NUMBER,attr"`
		OPERTYPEID string `xml:"OPERTYPEID,attr"`
		PREVID     int64  `xml:"PREVID,attr"`
		NEXTID     int64  `xml:"NEXTID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   string `xml:"ISACTUAL,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
	} `xml:"CARPLACE"`
}
