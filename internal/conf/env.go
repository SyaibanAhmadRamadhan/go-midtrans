package conf

import (
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-midtrans/internal/util"
	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	AppMode        string         `mapstructure:"APP_MODE" json:"APP_MODE"`
	WebConf        WebConf        `json:"WEB_CONF" mapstructure:"WEB_CONF"`
	PaymentGateway PaymentGateway `json:"PAYMENT_GATEWAY" mapstructure:"PAYMENT_GATEWAY"`
}

func LoadEnv() {
	var conf Conf
	dirRetryList := []string{``, `../`, `../../`, `../../../`}
	for _, dirPrefix := range dirRetryList {
		viper.SetConfigFile(dirPrefix + "env.json")
		viper.SetConfigType("json")

		err := viper.ReadInConfig()
		if err == nil {
			viper.SetConfigFile(dirPrefix + "env.override.json")
			viper.SetConfigType("json")
			err = viper.MergeInConfig()
			util.PanicIF(err)

			err = viper.Unmarshal(&conf)
			util.PanicIF(err)

			Config = &conf
			fmt.Println(Config)
			return
		}
	}
	panic(`cannot load .env and .env.override`)
}
