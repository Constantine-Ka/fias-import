package model_hierarchy

import "encoding/xml"

type MUNITEMS struct {
	XMLName xml.Name  `xml:"ITEMS"`
	Text    string    `xml:",chardata"`
	ITEM    []MunNode `xml:"ITEM"`
}
type MunNode struct {
	ID          string `xml:"ID,attr" json:"id,omitempty"`
	OBJECTID    string `xml:"OBJECTID,attr" json:"objectid,omitempty"`
	PARENTOBJID string `xml:"PARENTOBJID,attr" json:"parentobjid,omitempty"`
	CHANGEID    string `xml:"CHANGEID,attr" json:"changeid,omitempty"`
	OKTMO       string `xml:"OKTMO,attr" json:"oktmo,omitempty"`
	PREVID      string `xml:"PREVID,attr" json:"previd,omitempty"`
	NEXTID      string `xml:"NEXTID,attr" json:"nextid,omitempty"`
	UPDATEDATE  string `xml:"UPDATEDATE,attr" json:"updatedate,omitempty"`
	STARTDATE   string `xml:"STARTDATE,attr" json:"startdate,omitempty"`
	ENDDATE     string `xml:"ENDDATE,attr" json:"enddate,omitempty"`
	ISACTIVE    bool   `xml:"ISACTIVE,attr" json:"isactive,omitempty"`
	PATH        string `xml:"PATH,attr" json:"path,omitempty"`
	Text        string `xml:"TEXT,omitempty"`
}
