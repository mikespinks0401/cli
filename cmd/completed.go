/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	task "github.com/mikespinks0401/cobra-todo/model"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List of task completed today",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskList, err := task.GetList()
		if err != nil {
			fmt.Println("there was an error retrieiving the task")
		}
		filterToday := cmd.Flag("today")
		todayDate := time.Now().Format("2006-01-02")
		for _,val := range taskList{
			if val.Done != 1{
				continue	
			}
			
			if !val.Updated_at.Valid {
				continue	
			}
			taskDate, _ := time.Parse("2006-01-02 15:04:05", val.Updated_at.String)
			taskYearDay := taskDate.Format("2006-01-02")
			if filterToday.Value.String() == "true" && taskYearDay == todayDate{
				fmt.Println("Task:", val.Task,"Completed Today")
				return
			}
			fmt.Println("Task:", val.Task,"Completed:",val.Updated_at.String)
		}

	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	completedCmd.Flags().Bool("today", false, "add flag to filter completed task by done today")
}

func isEqualStrings(s1, s2 string)bool{
	return s1 == s2
}