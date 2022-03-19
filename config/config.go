package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var config *Config
var once sync.Once

type MySQL struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

type Redis struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

type Config struct {
	MySQL MySQL `yaml:"mysql"`
	Reids Redis `yaml:"redis"`
}

func GetInstance() *Config {
	once.Do(initConf)
	return config
}

func initConf() {
	content, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	config = &Config{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		panic(err)
	}
	if err := validator.Validate(config); err != nil {
		panic(err)
	}
}
