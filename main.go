package monobullet

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

// Config has all the details for the pushbullet client to start
type Config struct {
	APIKey     string `yaml:"apiKey"`
	DeviceName string `yaml:"deviceName"`
	Debug      bool   `yaml:"debug"`
}

var config *Config

// ConfigFromFile loads the config from ~/.monobullet
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

// Configuration for passing the config struct in directly
func Configuration(c *Config) {
	config = c
}

// Start the websocket connection for realtime
func Start(ctx context.Context) {
	wsConnect(ctx)
}

func AddOwnDevice() (*Device, error) {
	if "" == config.DeviceName {
		var err error
		config.DeviceName, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
	}

	device, err := getOwnDevice()
	if err == DeviceMissingError {
		return addDevice(&Device{
			Nickname: config.DeviceName,
		})
	} else if err != nil {
		return nil, err
	}
	return device, nil
}
