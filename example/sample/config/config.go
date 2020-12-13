package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yodstar/goutil/database/sqlbuilder"
)

// Config
type Config struct {
	Listen   string
	RootDir  string
	BaseURL  string
	DBConfig []*sqlbuilder.Conf
	DuerOS   struct {
		AppID  string
		Secret string
	}
	Logger struct {
		Outfile string
		Filter  string
		Level   string
	}
}

// Config.LoadFile
func (c *Config) LoadFile(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

var CONF = &Config{}

// LoadFile
func LoadFile(path string) (err error) {
	return CONF.LoadFile(path)
}
