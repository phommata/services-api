package config

import (
	"github.com/joeshaw/envdecode"
	"log"
)

type Conf struct {
	Server serverConf
	Db     dbConf
}

type serverConf struct {
	Port int `env:"SERVER_PORT,required"`
}

type dbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}