package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"./converter"
)

type DatabaseConfig struct {
	Adapter  string
	DBName   string
	Username string
	Password string
}

type SoftwareConfig struct {
	Name    string
	Version string
}

type Config struct {
	Database DatabaseConfig
	From     SoftwareConfig
	To       SoftwareConfig
}

func loadConfig(filename string) (*Config, error) {
	configFile, err := ioutil.ReadFile("./" + filename + ".json")
	if err != nil {
		return nil, err
	}

	var config = &Config{}
	err = json.Unmarshal(configFile, config)
	return config, err
}

func main() {
	log.Print("Loading config.json")
	config, err := loadConfig("config")
	if err != nil {
		log.Fatal(err)
	}

	fromSoft, exists := converter.Lookup(config.From.Name, config.From.Version)
	if !exists {
		log.Fatal("Unable to find version '%s' of software '%s", config.From.Version, config.From.Name)
	}
	toSoft, exists := converter.Lookup(config.To.Name, config.To.Version)
	if !exists {
		log.Fatal("Unable to find version '%s' of software '%s", config.To.Version, config.To.Name)
	}

	log.Print("Converting from '%s.%s' to '%s.%s'", config.From.Name, config.From.Version, config.To.Name, config.To.Version)

	converter := converter.NewConverter()
	converter.From(fromSoft)
	converter.To(toSoft)
	err = converter.Convert()
	if err != nil {
		log.Fatal(err)
	}
}
