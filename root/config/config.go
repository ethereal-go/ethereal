package config

import (
	"path/filepath"
	"github.com/spf13/viper"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	BasePath []string
	FileName string
}

// Load configuration data set in application
func (c Config) LoadConfigFromApp() {
	var err error
	c.FileName = "app.json"

	workPath := BasePathClient()
	c.BasePath = append(c.BasePath, filepath.Join(workPath, "config"), workPath)

	viper.SetConfigName("app")
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
func (c Config) addAllPathsConfig(paths []string) {
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