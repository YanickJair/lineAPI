package db

import (
	"io/ioutil"
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
	yamlFile, err := ioutil.ReadFile("config/db.yaml")
	if err != nil {
		panic("DB error: could not loud DB config.yaml file")
	}
	err = yaml.Unmarshal(yamlFile, c)

	if len(c.DbName) <= 0 {
		panic("no db name")
	}
	if err != nil {
		panic("file error: could not extract from db config file")
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
