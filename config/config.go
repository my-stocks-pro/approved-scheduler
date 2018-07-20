package config

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"fmt"
)

type TypeConfig struct {
	Tick         uint64
	Host         string
	HostAPI      string
	HostApproved string
	MethodGET    string
	MethodPOST   string
	RedisDB      string
}

func GetConfig() *TypeConfig {

	prod := os.Getenv("PROD")

	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/approved-scheduler-config.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	if prod == "1" {
		conf.MethodPOST = fmt.Sprintf(conf.MethodPOST, conf.HostApproved)
		conf.MethodGET = fmt.Sprintf(conf.MethodGET, conf.HostAPI)
	} else {
		conf.MethodPOST = fmt.Sprintf(conf.MethodPOST, conf.Host)
		conf.MethodGET = fmt.Sprintf(conf.MethodGET, conf.Host)
	}

	fmt.Println(conf)

	return conf
}