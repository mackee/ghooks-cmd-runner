package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type config struct {
	Port    int    `toml:"port"`
	Host    string `toml:"host"`
	Secret  string `toml:"secret"`
	Logfile string `toml:"logfile"`
	Pidfile string `toml:"pidfile"`
	Hook    []hook
}

type hook struct {
	Event          string   `toml:"event"`
	Cmd            string   `toml:"command"`
	Branch         string   `toml:"branch"`
	IncludeActions []string `toml:"include_actions"`
	ExcludeActions []string `toml:"exclude_actions"`
	AccessToken    string   `toml:"access_token"`
}

type hooks struct {
	Hook []hook
}

func loadFile(filename string) (string, error) {
	var err error
	buf, err := ioutil.ReadFile(filename)

	return string(buf), err
}

func loadToml(filename string, c config) (config, error) {
	var config config
	buf, err := loadFile(filename)
	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(buf), &config)
	if err != nil {
		return config, err
	}

	if config.Port == 0 {
		config.Port = c.Port
	}

	if config.Host == "" {
		config.Host = c.Host
	}

	if config.Logfile == "" {
		config.Logfile = c.Logfile
	}

	if config.Pidfile == "" {
		config.Pidfile = c.Pidfile
	}

	return config, err
}

func (h hook) isNotBlankAccessToken() bool {
	return h.AccessToken != ""
}
