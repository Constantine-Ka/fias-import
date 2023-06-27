package model_carplaces

import "encoding/xml"

type PARAMS struct {
	XMLName xml.Name `xml:"PARAMS"`
	Text    string   `xml:",chardata"`
	PARAM   []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"ID,attr"`
		OBJECTID    string `xml:"OBJECTID,attr"`
		CHANGEID    string `xml:"CHANGEID,attr"`
		CHANGEIDEND string `xml:"CHANGEIDEND,attr"`
		TYPEID      string `xml:"TYPEID,attr"`
		VALUE       string `xml:"VALUE,attr"`
		UPDATEDATE  string `xml:"UPDATEDATE,attr"`
		STARTDATE   string `xml:"STARTDATE,attr"`
		ENDDATE     string `xml:"ENDDATE,attr"`
	} `xml:"PARAM"`
}
