package model_hierarchy

import "encoding/xml"

type ADMITEMS struct {
	XMLName xml.Name  `xml:"ITEMS"`
	Text    string    `xml:",chardata"`
	ITEM    []AdmNode `xml:"ITEM"`
}
type ADMITEMSTwo struct {
	XMLName xml.Name     `xml:"ITEMS"`
	Text    string       `xml:",chardata"`
	ITEM    []AdmNodeTwo `xml:"ITEM"`
}
type AdmNode struct {
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
}
type AdmNodeTwo struct {
	PrimaryKey  int    `xml:"PrimaryKey"`
	ID          string `xml:"ID"`
	OBJECTID    string `xml:"OBJECTID"`
	PARENTOBJID string `xml:"PARENTOBJID"`
	CHANGEID    string `xml:"CHANGEID"`
	REGIONCODE  string `xml:"REGIONCODE"`
	AREACODE    string `xml:"AREACODE"`
	CITYCODE    string `xml:"CITYCODE"`
	PLACECODE   string `xml:"PLACECODE"`
	PLANCODE    string `xml:"PLANCODE"`
	STREETCODE  string `xml:"STREETCODE"`
	PREVID      string `xml:"PREVID"`
	NEXTID      string `xml:"NEXTID"`
	UPDATEDATE  string `xml:"UPDATEDATE"`
	STARTDATE   string `xml:"STARTDATE"`
	ENDDATE     string `xml:"ENDDATE"`
	ISACTIVE    string `xml:"ISACTIVE"`
	PATH        string `xml:"PATH"`
	FKITEMS     string `xml:"FK_ITEMS"`
}
