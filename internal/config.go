package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
	"os"
)

type Config struct {
	Directories struct {
		Docs   string `toml:"docs"`
		Public string `toml:"public"`
	} `toml:"directories"`
	Errors struct {
		Generic string `toml:"generic"`
	} `toml:"errors"`
	Routes map[string]string `toml:"routes"`
	Server struct {
		BehindProxy bool   `toml:"behind-proxy"`
		Domain      string `toml:"domain"`
		Host        string `toml:"host"`
		Key         string `toml:"key"`
		Port        int    `toml:"port"`
		Ssl         struct {
			Cert string `toml:"cert"`
			Key  string `toml:"key"`
		} `toml:"ssl"`
	} `toml:"server"`
}

// ReadConfig reads and decodes the toml config for the server
func ReadConfig() Config {
	data, err := os.ReadFile("./data/config.toml")
	if err != nil {
		log.Error("failed to read config", err)
	}

	var conf Config
	_, err = toml.Decode(string(data), &conf)
	if err != nil {
		log.Error("failed to decode config", err)
	}

	return conf
}
