package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Database struct {
		Adapter string `yaml:"adapter"`
		Host    string `yaml:"host"`
	} `yaml:"databaser"`
}

func getConf(mode string) (*Config, error) {
	yamlPath := "config." + mode + ".yml"
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
