package srv

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"log"
)

var Config *configFile

type configFile struct {
	Proxy struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"proxy"`
}

func (t *configFile) readFile(env string) []byte {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("./config/%s.yaml", env))
	if err != nil {
		log.Panicf("Config read error: %s", err.Error())
	}
	return yamlFile
}

func (t *configFile) Load() {
	file := t.readFile("config")
	err := yaml.Unmarshal(file, t)
	if err != nil {
		file = t.readFile("config.sample")
		err := yaml.Unmarshal(file, t)
		if err != nil {
			log.Panicf("Error parse yaml file: %s", err.Error())
		}
	}
}

func init() {
	Config = &configFile{}
	Config.Load()
}
