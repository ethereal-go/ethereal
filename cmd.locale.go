package ethereal


import (
	"github.com/spf13/cobra"
	"fmt"
)

var cmdLocale = &cobra.Command{
	Use:   "locale",
	Short: "Localization management",
	Long: ``,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		switch arg {
		case "fill" :
			I18nGraphQL().Fill()
			fmt.Println("Success fill locale in database! Good job!")
		default:
			fmt.Println("Argument '" + arg + "' is not defined. ")
		}
	},
}