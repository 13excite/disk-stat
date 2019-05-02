package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//DefaultConfigPath is path to toml config,
//usage if have not -c --config key when run script
const DefaultConfigPath = "./config.toml"

//Config struct represents api,database server and credentials settings
type Config struct {
	Server   string
	Database string
	Listen   string
}

// Read and parse the configuration file
func (c *Config) Read(filePath string) {
	if filePath == "" {
		filePath = DefaultConfigPath
	}
	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		log.Fatal(err)
	}
}
