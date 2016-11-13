package main

import "github.com/monofuel/monobullet"

type Config struct {
	ApiKey     string `yaml:"apiKey"`
	Realtime   bool   `yaml:"realtime"`
	DeviceName string `yaml:"deviceName"`
}

var config Config

func main() {
	monobullet.ConfigFromFile()
	monobullet.Start()
}
