package model

import "encoding/xml"

type HIERARCHYItem struct {
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
	ISACTIVE    string `xml:"ISACTIVE,attr"`
	PATH        string `xml:"PATH,attr"`
}

type HIERARCHY struct {
	XMLName xml.Name        `xml:"ITEMS"`
	Text    string          `xml:",chardata"`
	ITEM    []HIERARCHYItem `xml:"ITEM"`
}
