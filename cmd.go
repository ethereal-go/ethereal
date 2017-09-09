package ethereal

import (
"github.com/spf13/cobra"
"fmt"
"os"
"github.com/spf13/viper"
"github.com/mitchellh/go-homedir"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI your application",
	Long:`
╔═══╗ ╔════╗ ╔╗╔╗ ╔═══╗ ╔═══╗ ╔═══╗ ╔══╗ ╔╗
║╔══╝ ╚═╗╔═╝ ║║║║ ║╔══╝ ║╔═╗║ ║╔══╝ ║╔╗║ ║║
║╚══╗   ║║   ║╚╝║ ║╚══╗ ║╚═╝║ ║╚══╗ ║╚╝║ ║║
║╔══╝   ║║   ║╔╗║ ║╔══╝ ║╔╗╔╝ ║╔══╝ ║╔╗║ ║║
║╚══╗   ║║   ║║║║ ║╚══╗ ║║║║  ║╚══╗ ║║║║ ║╚═╗
╚═══╝   ╚╝   ╚╝╚╝ ╚═══╝ ╚╝╚╝  ╚═══╝ ╚╝╚╝ ╚══╝
	`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func CliExecute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seeCobraTest.yaml)")

	RootCmd.AddCommand(cmdLocale)
	RootCmd.AddCommand(cmdSeed)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".seeCobraTest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".seeCobraTest")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

