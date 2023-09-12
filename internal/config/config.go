package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	MySQL struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Database string `yaml:"database" env-required:"true"`
	} `yaml:"mysql" env-required:"true"`
	Listen struct {
		Address string `yaml:"address" env-required:"true"`
		Net     string `yaml:"net" env-required:"true"`
	} `yaml:"listen" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		log.Println(os.Getwd())
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			log.Println("config message:", err)
			help, err := cleanenv.GetDescription(instance, nil)
			log.Println("config message:", help)
			log.Fatal("config message:", err)
		}
	})
	return instance
}
