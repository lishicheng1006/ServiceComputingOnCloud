package cmd

import (
	"dailyProject/Agenda/service"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of current account",
	Long:  "Log out of current account",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.GetInstance().UserLogout(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("You have successfully logged out")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.
}
