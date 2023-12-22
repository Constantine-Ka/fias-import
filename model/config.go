package model

import "github.com/spf13/viper"

type Config struct {
	Db struct {
		Driver   string `yaml:"driver"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"db"`
	Tablename struct {
		Dict struct {
			ParamTypes         string `yaml:"ParamTypes"`
			NDocTypes          string `yaml:"NDocTypes"`
			NDockInds          string `yaml:"NDockInds"`
			AddHouseTypes      string `yaml:"AddHouseTypes"`
			AddressObjectTypes string `yaml:"AddressObjectTypes"`
			ApartmentTypes     string `yaml:"ApartmentTypes"`
			HouseTypes         string `yaml:"HouseTypes"`
			OperationTypes     string `yaml:"OperationTypes"`
			RoomTypes          string `yaml:"RoomTypes"`
			LevelsTypes        string `yaml:"LevelsTypes"`
		} `yaml:"dict"`
		Content struct {
			AdmHierarchy   string `yaml:"AdmHierarchy"`
			MunHierarchy   string `yaml:"MunHierarchy"`
			Apartments     string `yaml:"Apartments"`
			ApartmentsP    string `yaml:"ApartmentsP"`
			Carplaces      string `yaml:"Carplaces"`
			CarplacesP     string `yaml:"CarplacesP"`
			Houses         string `yaml:"Houses"`
			HousesP        string `yaml:"HousesP"`
			Rooms          string `yaml:"Rooms"`
			RoomsP         string `yaml:"RoomsP"`
			Steads         string `yaml:"Steads"`
			SteadsP        string `yaml:"SteadsP"`
			ChangeHistory  string `yaml:"ChangeHistory"`
			NormDocs       string `yaml:"NormDocs"`
			ReestrObject   string `yaml:"ReestrObject"`
			AddrObject     string `yaml:"AddrObject"`
			AddrObjectP    string `yaml:"AddrObjectP"`
			ObjectDivision string `yaml:"ObjectDivision"`
		} `yaml:"content"`
	} `yaml:"tablename"`
	VP *viper.Viper
}
