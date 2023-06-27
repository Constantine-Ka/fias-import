package model_objectAddr

import "encoding/xml"

type ITEMS struct {
	XMLName xml.Name `xml:"ITEMS"`
	Text    string   `xml:",chardata"`
	ITEM    []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		PARENTID string `xml:"PARENTID,attr"`
		CHILDID  string `xml:"CHILDID,attr"`
		CHANGEID string `xml:"CHANGEID,attr"`
	} `xml:"ITEM"`
}
