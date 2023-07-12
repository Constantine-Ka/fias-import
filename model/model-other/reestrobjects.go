package model_other

import "encoding/xml"

type REESTROBJECTS struct {
	XMLName xml.Name     `xml:"REESTR_OBJECTS"`
	Text    string       `xml:",chardata"`
	OBJECT  []ReestrNode `xml:"OBJECT"`
}
type ReestrNode struct {
	Text       string `xml:",chardata"`
	OBJECTID   string `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   string `xml:"CHANGEID,attr"`
	ISACTIVE   bool   `xml:"ISACTIVE,attr"`
	LEVELID    string `xml:"LEVELID,attr"`
	CREATEDATE string `xml:"CREATEDATE,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
}
