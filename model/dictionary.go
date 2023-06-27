package model

import "encoding/xml"

type ADDHOUSETYPES struct {
	XMLName   xml.Name `xml:"HOUSETYPES"`
	Text      string   `xml:",chardata"`
	HOUSETYPE []struct {
		Text       string `xml:",chardata"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		DESC       string `xml:"DESC,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
		NAME       string `xml:"NAME,attr"`
		ID         string `xml:"ID,attr"`
	} `xml:"HOUSETYPE"`
}

type ADDRESSOBJECTTYPES struct {
	XMLName           xml.Name `xml:"ADDRESSOBJECTTYPES"`
	Text              string   `xml:",chardata"`
	ADDRESSOBJECTTYPE []struct {
		Text       string `xml:",chardata"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		DESC       string `xml:"DESC,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
		NAME       string `xml:"NAME,attr"`
		LEVEL      string `xml:"LEVEL,attr"`
		ID         string `xml:"ID,attr"`
	} `xml:"ADDRESSOBJECTTYPE"`
}

type APARTMENTTYPES struct {
	XMLName       xml.Name `xml:"APARTMENTTYPES"`
	Text          string   `xml:",chardata"`
	APARTMENTTYPE []struct {
		Text       string `xml:",chardata"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		DESC       string `xml:"DESC,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
		NAME       string `xml:"NAME,attr"`
		ID         string `xml:"ID,attr"`
	} `xml:"APARTMENTTYPE"`
}

type HOUSETYPES struct {
	XMLName   xml.Name `xml:"HOUSETYPES"`
	Text      string   `xml:",chardata"`
	HOUSETYPE []struct {
		Text       string `xml:",chardata"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		DESC       string `xml:"DESC,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
		NAME       string `xml:"NAME,attr"`
		ID         string `xml:"ID,attr"`
	} `xml:"HOUSETYPE"`
}
type NDOCKINDS struct {
	XMLName  xml.Name `xml:"NDOCKINDS"`
	Text     string   `xml:",chardata"`
	NDOCKIND []struct {
		Text string `xml:",chardata"`
		NAME string `xml:"NAME,attr"`
		ID   string `xml:"ID,attr"`
	} `xml:"NDOCKIND"`
}

type NDOCTYPES struct {
	XMLName  xml.Name `xml:"NDOCTYPES"`
	Text     string   `xml:",chardata"`
	NDOCTYPE []struct {
		Text      string `xml:",chardata"`
		ENDDATE   string `xml:"ENDDATE,attr"`
		STARTDATE string `xml:"STARTDATE,attr"`
		NAME      string `xml:"NAME,attr"`
		ID        string `xml:"ID,attr"`
	} `xml:"NDOCTYPE"`
}

type OBJECTLEVELS struct {
	XMLName     xml.Name `xml:"OBJECTLEVELS"`
	Text        string   `xml:",chardata"`
	OBJECTLEVEL []struct {
		Text       string `xml:",chardata"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		NAME       string `xml:"NAME,attr"`
		LEVEL      string `xml:"LEVEL,attr"`
	} `xml:"OBJECTLEVEL"`
}

type OPERATIONTYPES struct {
	XMLName       xml.Name `xml:"OPERATIONTYPES"`
	Text          string   `xml:",chardata"`
	OPERATIONTYPE []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID,attr"`
		NAME       string `xml:"NAME,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
	} `xml:"OPERATIONTYPE"`
}

type PARAMTYPES struct {
	XMLName   xml.Name `xml:"PARAMTYPES"`
	Text      string   `xml:",chardata"`
	PARAMTYPE []struct {
		Text       string `xml:",chardata"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		CODE       string `xml:"CODE,attr"`
		DESC       string `xml:"DESC,attr"`
		NAME       string `xml:"NAME,attr"`
		ID         string `xml:"ID,attr"`
	} `xml:"PARAMTYPE"`
}

type ROOMTYPES struct {
	XMLName  xml.Name `xml:"ROOMTYPES"`
	Text     string   `xml:",chardata"`
	ROOMTYPE []struct {
		Text       string `xml:",chardata"`
		UPDATEDATE string `xml:"UPDATEDATE,attr"`
		ENDDATE    string `xml:"ENDDATE,attr"`
		STARTDATE  string `xml:"STARTDATE,attr"`
		ISACTIVE   string `xml:"ISACTIVE,attr"`
		DESC       string `xml:"DESC,attr"`
		NAME       string `xml:"NAME,attr"`
		ID         string `xml:"ID,attr"`
		SHORTNAME  string `xml:"SHORTNAME,attr"`
	} `xml:"ROOMTYPE"`
}
