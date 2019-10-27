package cmd

import (
	"dailyProject/Agenda/service"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "Delete the currently logged in account",
	Long:  "Delete the currently logged in account",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.GetInstance().DeleteUser(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("The user has been successfully deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

	// Here you will define your flags and configuration settings.
}
