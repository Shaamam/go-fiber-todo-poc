package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var AppConfig Config

// Config represents the structure of the configuration file
type Config struct {
	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Port     string `mapstructure:"port"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func LoadConfig() {

	env := flag.String("env", "dev", "Environment to use (e.g., dev, qa, prod)")
	flag.Parse()

	viper.SetConfigName(fmt.Sprintf("config.%s", *env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // Adjust the path as needed

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
