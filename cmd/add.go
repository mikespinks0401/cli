/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mikespinks0401/cobra-todo/model"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the DB",
	Long: `Add a task to the DB, the task should be printed when successfuly inserted`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("add must be called with a string")
		}
		taskToAdd := ""
		for _,val := range args{
			if taskToAdd == ""{
				taskToAdd = taskToAdd + val
			} else {
			taskToAdd = taskToAdd + " " + val
			}
		}

		if err := task.AddTask(taskToAdd); err != nil {
			panic("could not add task")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
