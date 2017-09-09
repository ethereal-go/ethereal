package ethereal


import (
	"github.com/spf13/cobra"
	"fmt"
)

var cmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "To populate your tables from the database",
	Long: ``,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		switch arg {
		case "up" :
			role := Role{Name: "User", DisplayName: "User", Description: "Simple user"}
			user := User{Email: "", Name: "", Password: "", Role: role}
			App.Db.Save(&user)

			fmt.Println("Success fill database! Good job!")
		default:
			fmt.Println("Argument '" + arg + "' is not defined. ")
		}
	},
}