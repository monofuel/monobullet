package monobullet

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey     string `yaml:"apiKey"`
	Realtime   bool   `yaml:"realtime"`
	DeviceName string `yaml:"deviceName"`
}

var config *Config

func ConfigFromFile() {
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(user.HomeDir + "/" + configFilename)
	if err != nil {
		log.Fatalln(err)
	}
	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

func Configuration(c *Config) {
	config = c
}

func Start() {
	if config.Realtime {
		wsConnect()
	}
}
