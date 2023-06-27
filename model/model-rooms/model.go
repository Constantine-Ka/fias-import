package model_rooms

import "encoding/xml"

type ROOMS struct {
	XMLName xml.Name `xml:"ROOMS"`
	Text    string   `xml:",chardata"`
	ROOM    []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		NUMBER     string `xml:"NUMBER,attr"`
		ROOMTYPE   string `xml:"ROOMTYPE,attr"`
		OPERTYPEID string `xml:"OPERTYPEID,attr"`
		PREVID     string `xml:"PREVID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   string `xml:"ISACTUAL,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		NEXTID     string `xml:"NEXTID,attr"`
	} `xml:"ROOM"`
}
