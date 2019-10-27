package cmd

import (
	"dailyProject/Agenda/service"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with a registered account",
	Long:  "Login with a registered account in the format of : Agenda login -uUsername -pPassword",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username == "" {
			log.Fatal(errors.New("Username is required"))
		}
		if password == "" {
			log.Fatal(errors.New("Password is required"))
		}

		if err := service.GetInstance().UserLogin(username, password); err != nil {
			log.Fatal(err)
		}

		fmt.Println("You have successfully logged in")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
}
