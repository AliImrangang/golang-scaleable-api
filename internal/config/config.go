package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPSERVER struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPSERVER  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		log.Fatalf("config file does not exist:%s", configPath)

	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg

}

//this configuration for postgres db conectivity
// type Config struct {
// 	Env         string
// 	Addr        string
// 	DBHost      string
// 	DBPort      string
// 	DBUser      string
// 	DBPassword  string
// 	DBName      string
// }

// func MustLoad() *Config {
// 	// In a real implementation, you'd load these from environment variables or a config file
// 	return &Config{
// 		Env:         "development",
// 		Addr:        ":8080",
// 		DBHost:      "localhost",
// 		DBPort:      "5432",
// 		DBUser:      "youruser",
// 		DBPassword:  "yourpassword",
// 		DBName:      "yourdatabase",
// 	}
// }
