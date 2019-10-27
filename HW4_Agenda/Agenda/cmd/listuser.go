package cmd

import (
	"dailyProject/Agenda/service"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listuserCmd represents the listuser command
var listuserCmd = &cobra.Command{
	Use:   "listuser",
	Short: "List all users",
	Long:  "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.GetInstance().ListAllUsers(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("The user list has been shown")
	},
}

func init() {
	rootCmd.AddCommand(listuserCmd)

	// Here you will define your flags and configuration settings.
}
