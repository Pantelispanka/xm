package config

import "os"

var config Config

type Config struct {
	mongoUrl string
}

func (c *Config) config() {
	config.mongoUrl = os.Getenv("MONGO_URL")
}
