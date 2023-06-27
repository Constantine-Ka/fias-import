package model_steads

import "encoding/xml"

type STEADS struct {
	XMLName xml.Name `xml:"STEADS"`
	Text    string   `xml:",chardata"`
	STEAD   []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		NUMBER     string `xml:"NUMBER,attr"`
		OPERTYPEID string `xml:"OPERTYPEID,attr"`
		PREVID     string `xml:"PREVID,attr"`
		NEXTID     string `xml:"NEXTID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   string `xml:"ISACTUAL,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
	} `xml:"STEAD"`
}
