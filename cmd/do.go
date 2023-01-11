/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/mikespinks0401/cobra-todo/db"
	task "github.com/mikespinks0401/cobra-todo/model"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Change a command to done by selecting the number corresponding to the task",
	Long: `Use the command line to change a task to done by selecting the corresponding number`,
	Run: func(cmd *cobra.Command, args []string) {
		list,err := task.GetList()
		if err != nil{
			fmt.Println("error fetching the list")
		}
		var unfinishedList []task.Task

		for _, val := range list{
			if val.Done == 0 {
				unfinishedList = append(unfinishedList, val)
			}
		}	
		fmt.Println("Please select the corresponding number of the task you would like to complete")
		for i, val := range unfinishedList{
			fmt.Printf("%d. %s\n", i+1, val.Task)
		}	
		var selected int	
		fmt.Scan(&selected)
		for selected < 1 || selected > len(unfinishedList){
			fmt.Println("Must enter a list number that exist")	
			fmt.Scan(&selected)
		}

		taskToUpdate := unfinishedList[selected - 1]
		conn, err := db.DBConn()
		if err != nil{
			fmt.Println("there was an error connecting to the db:", err.Error())
			return
		}
		stmt, err := conn.Prepare("UPDATE todo SET done=1, updated_at=? WHERE task=?")
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		todayDate := time.Now().Format("2006-01-02 15:04:05")
		res, err := stmt.Exec(todayDate, taskToUpdate.Task)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1{
			fmt.Println("Successfully updated the task")
		}
	
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
