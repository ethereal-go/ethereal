package ethereal

import (
	"github.com/spf13/cobra"
	"fmt"
)

var cmdDatabase = &cobra.Command{
	Use:   "database",
	Short: "Cli database",
	Long: ``,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		switch arg {
		case "migrate" :
			App.Db.AutoMigrate(tables()...)
			fmt.Println("Success migrate tables in database! Good job!")
		default:
			fmt.Println("Argument '" + arg + "' is not defined. ")
		}
	},
}