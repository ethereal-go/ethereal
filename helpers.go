package ethereal

import (
	"github.com/spf13/viper"
	"strings"
	"os"
)

// Here functions helpers
// ----------------------------------

/**
 / Get configuration value
 */
func config(name string, byDefault ...interface{}) interface{} {
	var temp string
	if temp = os.Getenv(name); temp == "" {
		if temp = viper.GetString(strings.ToLower(name)); temp == ""{
			viper.SetDefault(name, byDefault)
		}
	}
	return temp
}

