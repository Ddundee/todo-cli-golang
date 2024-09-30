package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func ClearCmd(cmd *cobra.Command, args []string) {
	os.Create("data/list.csv")
	fmt.Println("Cleared all items")
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears all items",
	Run:   ClearCmd,
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
