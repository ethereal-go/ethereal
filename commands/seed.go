package commands
//
//import (
//	"github.com/spf13/cobra"
//	"fmt"
//	"github.com/ethereal-go/ethereal"
//)
//
//var CmdSeed = &cobra.Command{
//	Use:   "seed",
//	Short: "To populate your tables from the database",
//	Long: ``,
//	Args: cobra.MaximumNArgs(1),
//	Run: func(cmd *cobra.Command, args []string) {
//		arg := args[0]
//		switch arg {
//		case "up" :
//			role := ethereal.Role{Name: "User", DisplayName: "User", Description: "Simple user"}
//			user := ethereal.User{Email: "", Name: "", Password: "", Role: role}
//			ethereal.App.Db.Save(&user)
//
//			fmt.Println("Success fill database! Good job!")
//		default:
//			fmt.Println("Argument '" + arg + "' is not defined. ")
//		}
//	},
//}