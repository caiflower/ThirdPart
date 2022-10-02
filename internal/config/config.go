package config

import (
	"flag"
	"github.com/jinzhu/configor"
	"os"
)

var Path string
var EnableENV bool
var Config ThirdPartConfig

const (
	EMAIL_SERVER_ADDRESS = "EMAIL_SERVER_ADDRESS"
	EMAIL_USERNAME       = "EMAIL_USERNAME"
	EMAIL_PASSWORD       = "EMAIL_PASSWORD"
)

type (
	ThirdPartConfig struct {
		Email
		Server
	}

	Email struct {
		Address  string
		UserName string
		Password string
	}

	Server struct {
		Port string `default:"8080"`
	}
)

func init() {
	flag.BoolVar(&EnableENV, "e", false, "--e true")
	flag.StringVar(&Path, "config", "config/config.yml", "--config xxx")
}

func Init() (e error) {
	e = configor.Load(&Config, Path)
	if EnableENV {
		Config.Address = os.Getenv(EMAIL_SERVER_ADDRESS)
		Config.UserName = os.Getenv(EMAIL_USERNAME)
		Config.Password = os.Getenv(EMAIL_PASSWORD)
	}
	return
}
