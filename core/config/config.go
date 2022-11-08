package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)


type IConfig interface {
}


type Config struct {
	Environment   string
	Server        Server
	Database      Database
	Kafka         Kafka
	Redis         Redis
	
}

type Server struct {
	Port       uint64
	PathPrefix string
}

type Database struct {
	Host     string
	Port     uint64
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Kafka struct {
	Brokers          string
	Topic            string
	ClientID         string
	ConsumerGroup    string
	InitialOffset    string
	SecurityProtocol string
	Mechanism        string
	SaslUser         string
	SaslPassword     string
}

type Redis struct {
	Host string
	Pw   string
}

func NewConfig() *Config {

	var c Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config/")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("config file not found")
		} else {
			panic(fmt.Errorf("CONFIG:fatal error config file: %s ", err))
		}
	}
	viper.AutomaticEnv()
	viper.Unmarshal(&c)

	return &c
}
