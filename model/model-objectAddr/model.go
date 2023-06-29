package model_objectAddr

import "encoding/xml"

type ADDRESSOBJECTS struct {
	XMLName xml.Name `xml:"ADDRESSOBJECTS"`
	Text    string   `xml:",chardata"`
	OBJECT  []struct {
		Text       string `xml:",chardata"`
		ID         int64  `xml:"ID,attr"`
		OBJECTID   int64  `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   int64  `xml:"CHANGEID,attr"`
		NAME       string `xml:"NAME,attr"`
		TYPENAME   string `xml:"TYPENAME,attr"`
		LEVEL      int64  `xml:"LEVEL,attr"`
		OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
		PREVID     int64  `xml:"PREVID,attr"`
		NEXTID     int64  `xml:"NEXTID,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		ISACTUAL   int8   `xml:"ISACTUAL,attr"`
		ISACTIVE   int8   `xml:"ISACTIVE,attr"`
	} `xml:"OBJECT"`
}
