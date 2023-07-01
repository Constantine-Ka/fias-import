package config

import (
	"fias-import_byLondon/model"
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	*model.Config
}

func GetConf() *model.Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &model.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
