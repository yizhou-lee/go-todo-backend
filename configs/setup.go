package configs

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig reads the configuration file.
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
