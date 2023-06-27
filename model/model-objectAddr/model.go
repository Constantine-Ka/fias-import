package model_objectAddr

import "encoding/xml"

type ADDRESSOBJECTS struct {
	XMLName xml.Name `xml:"ADDRESSOBJECTS"`
	Text    string   `xml:",chardata"`
	OBJECT  []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		NAME       string `xml:"NAME,attr"`
		TYPENAME   string `xml:"TYPENAME,attr"`
		LEVEL      string `xml:"LEVEL,attr"`
		OPERTYPEID string `xml:"OPERTYPEID,attr"`
		PREVID     string `xml:"PREVID,attr"`
		NEXTID     string `xml:"NEXTID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   string `xml:"ISACTUAL,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
	} `xml:"OBJECT"`
}
