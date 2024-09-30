/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/mergestat/timediff"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

func ListCmd(cmd *cobra.Command, args []string) {
	file, err := os.Open("data/list.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tDate\tStatus")
	for _, row := range data {
		itemTime, _ := time.Parse(time.RFC3339, row[2])
		rData := fmt.Sprintf("%v\t%v\t%v\t%s", row[0], row[1], timediff.TimeDiff(itemTime), row[3])
		_, err := fmt.Fprintln(w, rData)
		if err != nil {
			return
		}
	}
	w.Flush()
	file.Close()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ListCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
