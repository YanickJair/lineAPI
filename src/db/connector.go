package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/go-pg/pg/v9"
	"gopkg.in/yaml.v2"
)

var once sync.Once
var dataBase *pg.DB

//DBConfig config
type dBConfig struct {
	DbName   string `yaml:"dbName"`
	DbUser   string `yaml:"dbUser"`
	Password string `yaml:"password,omitempty"`
	Host     string `yaml:"host,omitempty"`
}

func (c *dBConfig) getConfig() *dBConfig {
	fmt.Println("Loading...")
	yamlFile, err := ioutil.ReadFile("config/db.yaml")
	if err != nil {
		log.Printf("file error #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal error #%v", err)
	}
	return c
}

//Connector - create a new connection to DB
func Connector() *pg.DB {
	once.Do(func() {
		var config dBConfig
		config.getConfig()
		dataBase = pg.Connect(&pg.Options{
			User:     config.DbUser,
			Database: config.DbName,
			Password: config.Password,
		})
	})
	return dataBase
}
