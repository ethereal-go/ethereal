package ethereal

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"github.com/agoalofalife/ethereal/utils"
)

type Config struct {
	basePath string
	fileName string
}

// Load configuration data set in application
func (c Config) LoadConfigFromApp() {
	var err error
	_, currentPath, _, _ := runtime.Caller(0)
	_, pathUser, _, _ := runtime.Caller(1)

	c.fileName = "app.json"
	c.basePath = path.Dir(currentPath) + "/config/"

	var cpan string
	if cpan, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "config", "app.json")

	fmt.Println(utils.FileExists(appConfigPath))
	fmt.Println(cpan)

	//fmt.Println(utils.FileExists(path.Dir(pathUser) + "/" + c.fileName))
	fmt.Println(pathUser)
	//fmt.Println(path.Dir(pathUser) + "/" + c.fileName)

	viper.SetConfigName("app")
	viper.AddConfigPath(c.basePath)

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
