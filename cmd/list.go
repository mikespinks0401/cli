/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/mikespinks0401/cobra-todo/model"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks that need to be done",
	Long: `List all task in the database that are not done yet`,
	Run: func(cmd *cobra.Command, args []string) {
		done := cmd.Flag("done")
		isDone := done.Value
		fmt.Println("The value of Done",isDone.String())	
		list, err := task.GetList()
		if err != nil {
			fmt.Println(err.Error())
		}
		index := 1
		for _,val := range list{
			var done bool
			if val.Done == 1 {
				done = true
			} else {
				done = false
			}
			if isDone.String() == "true" && !done {
				continue
			}	
			taskTime,_ := time.Parse("2006-01-02 15:04:05",val.Created_at)
			formatted := taskTime.Format("2006 01/02")
			fmt.Printf("%d. %s created:%s\n", index, val.Task, formatted)
			index++

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().Bool("done", false, "Set this to recieve only the task that are done")
}
