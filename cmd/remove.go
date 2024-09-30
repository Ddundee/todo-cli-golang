package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var removeId int64

func RemoveCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 && removeId == -1 {
		fmt.Println("REQUIRED FLAG: remove [id] or --id\tRemoves a new todo item")
		return
	}
	if removeId == -1 {
		parsedInt, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("INVALID ID: remove [id] or --id\tRemoves a new todo item")
			return
		}
		removeId = parsedInt

	}

	file, err := os.Open("data/list.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	var newRows [][]string
	rows, _ := reader.ReadAll()
	for ind, _ := range rows {
		parsedInt, _ := strconv.ParseInt(rows[ind][0], 10, 64)
		if parsedInt != removeId {
			newRows = append(newRows, rows[ind])
		}
	}

	file, err = os.Create("data/list.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(newRows)
	fmt.Printf("Removed todo with ID %v\n", removeId)
	file.Close()

}

var removeCmd = &cobra.Command{
	Use:   "remove [completeId] or --completeId",
	Short: "removes the todo item",
	Args:  cobra.MaximumNArgs(1),
	Run:   RemoveCmd,
}

func init() {
	removeCmd.Flags().Int64Var(&removeId, "id", -1, "ID of the todo")
	rootCmd.AddCommand(removeCmd)
}
