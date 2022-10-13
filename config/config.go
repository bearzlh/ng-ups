package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Config             *viper.Viper
)

func init() {
	var err error
	Config = viper.New()
	Config.SetConfigFile("config.json") // name of config file (without extension)
	err = Config.ReadInConfig()         // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error mysql config file: %w \n", err))
	}
}
