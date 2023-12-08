package config

import (
	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	URL HostConfig `gcfg:"URL"`
	DB  DBConfig   `gcfg:"DB"`
}

type HostConfig struct {
	Host string `gcfg:"Host"`
}

type DBConfig struct {
	Host       string `gcfg:"Host"`
	Credential string `gcfg:"Credential"`
	Port       string `gcfg:"Port"`
	User       string `gcfg:"User"`
	DBName     string `gcfg:"DBName"`
}

func ReadModuleConfig(cfg interface{}, path string) bool {
	ext := "config.ini"

	fname := path + ext
	err := gcfg.ReadFileInto(cfg, fname)
	return err == nil
}
