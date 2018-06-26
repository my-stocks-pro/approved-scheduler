package config

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type TypeConfig struct {
	Tick    uint64
	URLGET  string
	URLPOST string
	RadisDB string
}

func GetConfig() *TypeConfig {

	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/approved-scheduler-config.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	return conf
}
