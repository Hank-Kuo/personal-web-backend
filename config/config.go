package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Version string `yaml:"version"`

	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Database struct {
		Host    string `yaml:"host"`
		Adapter string `yaml:"adapter"`
	} `yaml:"database"`
}

func GetConf(mode string) *Config {
	yamlPath := "../config/config." + mode + ".yml"
	yamlFile, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		panic(err)
	}

	c := &Config{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}
