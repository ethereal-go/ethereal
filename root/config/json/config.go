package json

import (
	"strings"
	"path/filepath"
	"github.com/spf13/viper"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"github.com/ethereal-go/ethereal/root/config"
)
const (
	FileName = "app.json"
	DirConf  = "config"
)

type ConfigurationJson struct {
	BasePath          []string
	FileName          string
	ExtensionFileName string
}

func NewConfig() config.Configurable {
	return &ConfigurationJson{}
}

// Load configuration data set in application
func (c *ConfigurationJson) Load() {
	var err error
	s := strings.Split(FileName, ".")
	c.FileName, c.ExtensionFileName = s[0], s[1]

	workPath := BasePathClient()
	c.BasePath = append(c.BasePath, filepath.Join(workPath, DirConf), workPath)

	viper.SetConfigName(c.FileName)
	c.addAllPathsConfig(c.BasePath)

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
/ Set all paths possible in application
*/
func (c ConfigurationJson) addAllPathsConfig(paths []string) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
}

func BasePathClient() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return workPath
}

