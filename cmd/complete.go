package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var id int64

func CompleteCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 && id == -1 {
		fmt.Println("REQUIRED FLAG: complete [id] or --id\tChecks a new todo item")
		return
	}
	if id == -1 {
		parsedInt, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("INVALID ID: complete [id] or --id\tChecks a new todo item")
			return
		}
		id = parsedInt
	}

	file, err := os.Open("data/list.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()
	file.Close()

	rowEdited := false
	for ind, _ := range rows {
		parsedInt, _ := strconv.ParseInt(rows[ind][0], 10, 63)
		if parsedInt == id {
			rowEdited = true
			if rows[ind][3] != "pending" {
				fmt.Println("IT IS ALREADY COMPLETE!")
				return
			}
			rows[ind][3] = "done"
		}
	}
	if !rowEdited {
		fmt.Println("No TODO item found.")
		return
	}

	file, err = os.Create("data/list.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(rows)
	fmt.Printf("Edited todo with ID %v\n", id)
	file.Close()

}

var completeCmd = &cobra.Command{
	Use:   "complete [id] or --id",
	Short: "checks the todo item",
	Args:  cobra.MaximumNArgs(1),
	Run:   CompleteCmd,
}

func init() {
	addCmd.Flags().Int64Var(&id, "id", -1, "ID of the todo")
	rootCmd.AddCommand(completeCmd)
}
