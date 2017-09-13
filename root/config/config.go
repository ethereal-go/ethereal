package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Configurable interface {
	Load()
}

/**
/ Get configuration value
*/

func GetCnf(name string, byDefault ...interface{}) interface{} {
	var temp string
	if temp = os.Getenv(name); temp == "" {
		if temp = viper.GetString(strings.ToLower(name)); temp == "" {
			viper.SetDefault(name, byDefault)
		}
	}
	return temp

}
