package ethereal

import (
	"fmt"
	"github.com/agoalofalife/ethereal/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"log"
)

type Config struct {
	basePath string
	fileName string
}

// Load configuration data set in application
func (c Config) LoadConfigFromApp() {
	var err error
	c.fileName = "app.json"

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	c.basePath = filepath.Join(workPath, "config")
	appConfigPath := filepath.Join(workPath, "config", "app.json")

	if !utils.FileExists(path.Dir(appConfigPath)) {
		panic("Not Found config file.")
	}

	viper.SetConfigName("app")
	viper.AddConfigPath(c.basePath)

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("kFatal error config file: %s \n", err))
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
