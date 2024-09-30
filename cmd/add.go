package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var name string

func AddCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 && name == "" {
		fmt.Println("REQUIRED FLAG: add [name] or --name\tAdds a new todo item")
		return
	}
	if name == "" {
		name = args[0]
	}

	file, err := os.Open("data/list.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	file, err = os.OpenFile("data/list.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	row := []string{fmt.Sprint(len(rows)), name, time.Now().Format(time.RFC3339), "pending"}
	writer.Write(row)
	fmt.Printf("Added \"%s\" to todo", name)
	writer.Flush()
	file.Close()
}

var addCmd = &cobra.Command{
	Use:   "add [name] or --name",
	Short: "Adds a new todo item",
	Args:  cobra.MaximumNArgs(1),
	Run:   AddCmd,
}

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the todo")
	rootCmd.AddCommand(addCmd)
}
