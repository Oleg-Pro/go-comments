package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	DatabaseUrlConfigKey = "DATABASE_URL"
	PortConfigKey        = "PORT"
)

var Conf *Config
var configKeys = [...]string{DatabaseUrlConfigKey, PortConfigKey}

type NewRelicConfig struct {
	AppName string
	Enabled bool
	License string
}

type Config struct {
	AppSecret    string
	DatabaseUrl  string
	Port         string
	NewRelicConf NewRelicConfig
}

func Load() {
	godotenv.Load()

	checkConfig()
	Conf = createConfig()
}

func checkConfig() {
	for _, key := range configKeys {
		exitIfEvnKeyIsAbsent(key)
	}
}

func createConfig() *Config {
	return &Config{
		DatabaseUrl: os.Getenv(DatabaseUrlConfigKey),
		Port:        os.Getenv(PortConfigKey),
	}
}

func exitIfEvnKeyIsAbsent(key string) {
	if value := os.Getenv(key); value != "" {
		log.Printf("%v %v", key, value)
	} else {
		log.Fatalln(fmt.Sprintf("No %s parameter", key))
	}
}
