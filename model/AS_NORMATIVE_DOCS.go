package model

import "encoding/xml"

type NORMDOCS struct {
	XMLName xml.Name `xml:"NORMDOCS"`
	Text    string   `xml:",chardata"`
	NORMDOC []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		NAME       string `xml:"NAME,attr"`
		DATE       string `xml:"DATE,attr"`
		NUMBER     string `xml:"NUMBER,attr"`
		TYPE       string `xml:"TYPE,attr"`
		KIND       string `xml:"KIND,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
	} `xml:"NORMDOC"`
}
