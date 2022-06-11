package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
	S3AccessKey string `mapstructure:"S3_ACCESS_KEY"`
	S3SecretKey string `mapstructure:"S3_SECRET_KEY"`
}

var Conf *Config

func Load() {
	// tell viper where our config file is
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// override values that it has read from config file with the values of the corresponding environment variables if they exist
	viper.AutomaticEnv()

	// read in config values
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// unmarshal config to struct
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
}