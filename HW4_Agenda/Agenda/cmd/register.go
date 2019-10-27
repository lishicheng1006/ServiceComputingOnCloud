package cmd

import (
	"dailyProject/Agenda/service"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  "Register a new account in the format of : Agenda register -uUsername -pPassword -eEmail -cPhone",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		if username == "" {
			log.Fatal(errors.New("Username is required"))
		}
		if password == "" {
			log.Fatal(errors.New("Password is required"))
		}
		if email == "" {
			log.Fatal(errors.New("Email is required"))
		}
		if phone == "" {
			log.Fatal(errors.New("Phone is required"))
		}

		if err := service.GetInstance().CreateUser(username, password, email, phone); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Register a user successfully")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	registerCmd.Flags().StringP("username", "u", "", "Non-repeatable username")
	registerCmd.Flags().StringP("password", "p", "", "Password")
	registerCmd.Flags().StringP("email", "e", "", "Email address")
	registerCmd.Flags().StringP("phone", "c", "", "Phone number")
}
