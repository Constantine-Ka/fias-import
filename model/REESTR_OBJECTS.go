package model

import "encoding/xml"

type REESTROBJECTS struct {
	XMLName xml.Name `xml:"REESTR_OBJECTS"`
	Text    string   `xml:",chardata"`
	OBJECT  []struct {
		Text       string `xml:",chardata"`
		OBJECTID   string `xml:"OBJECTID,attr"`
		OBJECTGUID string `xml:"OBJECTGUID,attr"`
		CHANGEID   string `xml:"CHANGEID,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		LEVELID    string `xml:"LEVELID,attr"`
		CREATEDATE string `xml:"CREATEDATE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
	} `xml:"OBJECT"`
}
