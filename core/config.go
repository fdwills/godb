package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type db_config struct {
	Read        string
	Write       string
	MaxPoolSize int `yaml:"max_pool_size"`
	ShardId     int `yaml:"shard_id"`
}

type config struct {
	Mode     string
	Listen   string
	DBDriver string    `yaml:"db_driver"`
	DBMaster db_config `yaml:"db_master"`
}

var config_instance *config

func GetConfig() *config {
	if config_instance == nil {

		dat, err := ioutil.ReadFile("./config/production.yaml")
		if err != nil {
			log.Fatal("failed to read: %v", err)
			return nil
		}
		conf := config{}
		yaml.Unmarshal(dat, &conf)
		if err != nil {
			log.Fatal("failed to load config: %v", err)
			return nil
		}

		config_instance = &conf
	}
	return config_instance
}
