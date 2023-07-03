package model_hierarchy

import "encoding/xml"

type ADMITEMS struct {
	XMLName xml.Name `xml:"ITEMS"`
	Text    string   `xml:",chardata"`
	ITEM    []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"ID,attr"`
		OBJECTID    string `xml:"OBJECTID,attr"`
		PARENTOBJID string `xml:"PARENTOBJID,attr"`
		CHANGEID    string `xml:"CHANGEID,attr"`
		REGIONCODE  string `xml:"REGIONCODE,attr"`
		PREVID      string `xml:"PREVID,attr"`
		NEXTID      string `xml:"NEXTID,attr"`
		UPDATEDATE  string `xml:"UPDATEDATE,attr"`
		STARTDATE   string `xml:"STARTDATE,attr"`
		ENDDATE     string `xml:"ENDDATE,attr"`
		ISACTIVE    bool   `xml:"ISACTIVE,attr"`
		PATH        string `xml:"PATH,attr"`
	} `xml:"ITEM"`
}
