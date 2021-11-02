package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	viper := viper.New()
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../..")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return viper
}
