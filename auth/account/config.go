package account

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App   appConfig `yaml:"app"`
	Db    database  `yaml:"database"`
	Redis redis     `yaml:"redis"`
}
type appConfig struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Salt string `yaml:"salt"`
}

type database struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
}

type redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
}

func NewConfig(path string) *Config {
	conf := &Config{}
	yamlFile, err := ioutil.ReadFile(path + "/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("config", conf.Db.Host)
	return conf
}
