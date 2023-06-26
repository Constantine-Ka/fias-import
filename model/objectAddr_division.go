package model

import "encoding/xml"

type ITEM struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"ID,attr"`
	PARENTID string `xml:"PARENTID,attr"`
	CHILDID  string `xml:"CHILDID,attr"`
	CHANGEID string `xml:"CHANGEID,attr"`
}
type Devisions struct {
	XMLName xml.Name `xml:"ITEMS"`
	Text    string   `xml:",chardata"`
	ITEM    []ITEM   `xml:"ITEM"`
}
