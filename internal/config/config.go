package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Servers []struct {
		Hostname string `toml:"hostname"`
		Port     uint16 `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"servers"`
}

func GetConfig() *Config {
	var conf Config
	_, err := toml.DecodeFile("config.toml", &conf)

	if err != nil {
		log.Fatal(err)
	}
	if len(conf.Servers) == 0 {
		log.Fatal("No servers found in config file")
	}

	return &conf
}
