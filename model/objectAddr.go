package model

type AddrObject struct {
	ID         string `xml:"ID,attr"`
	OBJECTID   string `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   string `xml:"CHANGEID,attr"`
	NAME       string `xml:"NAME,attr"`
	TYPENAME   string `xml:"TYPENAME,attr"`
	LEVEL      int64  `xml:"LEVEL,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   bool   `xml:"ISACTUAL,attr"`
	ISACTIVE   bool   `xml:"ISACTIVE,attr"`
}
type Objects struct {
	ObjectField AddrObject `xml:"OBJECT"`
}

type ADDRESSOBJECTS struct {
	ObjectField []Objects `xml:"OBJECT"`
}
