package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var confiPath string

	confiPath = os.Getenv("CONFIG_PATH")

	if confiPath == "" {
		falgs := flag.String("config", "", "path to deconfiguration file")
		flag.Parse()

		confiPath = *falgs

		if confiPath == "" {
			log.Fatal("Config path is not set at all.")
		}
	}

	if _, err := os.Stat(confiPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exists: %s.", confiPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(confiPath, &cfg)

	if err != nil {
		log.Fatalf("can't read config file %s", err.Error())
	}

	return &cfg
}
