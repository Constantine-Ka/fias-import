package model_other

import "encoding/xml"

type CHANGEHISTORY struct {
	XMLName xml.Name      `xml:"ITEMS"`
	Text    string        `xml:",chardata"`
	ITEM    []HistoryNode `xml:"ITEM"`
}

type HistoryNode struct {
	Text        string `xml:",chardata"`
	CHANGEID    string `xml:"CHANGEID,attr"`
	OBJECTID    string `xml:"OBJECTID,attr"`
	ADROBJECTID string `xml:"ADROBJECTID,attr"`
	OPERTYPEID  string `xml:"OPERTYPEID,attr"`
	CHANGEDATE  string `xml:"CHANGEDATE,attr"`
	NDOCID      int64  `xml:"NDOCID,attr"`
}
