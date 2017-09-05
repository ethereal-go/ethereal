package ethereal

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

type Config struct {
}

// Load configuration data set in application
func (c Config) LoadConfigFromApp() {
	_, currentPath, _, _ := runtime.Caller(0)
	viper.SetConfigName("app")
	viper.AddConfigPath(path.Dir(currentPath) + "/config")
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}
