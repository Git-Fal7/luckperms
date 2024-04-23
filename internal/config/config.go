package config

import (
	"log"

	"github.com/spf13/viper"
)

var ViperConfig = viper.New()

func InitConfig() {
	ViperConfig.AddConfigPath(".")
	ViperConfig.SetConfigName("luckperms")
	ViperConfig.SetConfigType("yaml")

	ViperConfig.SetDefault("database.hostname", "localhost")
	ViperConfig.SetDefault("database.port", 5432)
	ViperConfig.SetDefault("database.username", "admin")
	ViperConfig.SetDefault("database.password", "adminpassword")
	ViperConfig.SetDefault("database.database", "minecraft")

	err := ViperConfig.ReadInConfig()
	if err != nil {
		// Create config file
		log.Println("Couldn't find luckperms.yml, creating a new config file")
		ViperConfig.WriteConfigAs("./luckperms.yml")
	}

}
