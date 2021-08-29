package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Aws      Aws
	Database Database
	Router   Router
}

type Aws struct {
	DefaultRegion string `required:"true" split_words:"true"`
}

type Database struct {
	Username      string `required:"true"`
	Password      string `required:"true"`
	Host          string `required:"true"`
	Port          string `required:"true"`
	Database      string `required:"true"`
	AutoMigration bool   `default:"false" split_words:"true"`
}

type Router struct {
	Port         uint     `default:"8080"`
	AllowOrigins []string `required:"true" split_words:"true"`
}

func Get() (Config, error) {
	var config Config
	var err error

	err = envconfig.Process("aws", &config.Aws)
	fmt.Printf("aws config: %v", config.Aws)
	if err != nil {
		return config, err
	}

	err = envconfig.Process("database", &config.Database)
	fmt.Printf("database config: %v", config.Database)
	if err != nil {
		return config, err
	}

	err = envconfig.Process("router", &config.Router)
	fmt.Printf("router config: %v", config.Router)
	if err != nil {
		return config, err
	}

	return config, nil
}
