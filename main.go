package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

var configFilename = ".monobullet"

type Config struct {
	ApiKey string `yaml:"apiKey"`
}

var config Config

func init() {
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

func main() {

}
