/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mikespinks0401/cobra-todo/db"
	task "github.com/mikespinks0401/cobra-todo/model"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := task.GetList()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := task.PrintList(); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Please select the corresponding number to delete the task")

		var selected int
		fmt.Scan(&selected)
		fmt.Printf("Are you sure you would like to delete:(yes/no) \n%d. %s\n",selected, list[selected - 1].Task)
		var confirm string
		fmt.Scan(&confirm)
		if confirm == "no"{
			fmt.Println("Good bye")
		}
		conn, err := db.DBConn()
		if err != nil{
			fmt.Println(err.Error())
		}
		defer conn.Close()
		stmt, err := conn.Prepare("DELETE FROM todo WHERE task=?")
		if err != nil{
			fmt.Println(err.Error())
		}
		res, err := stmt.Exec(list[selected - 1].Task)
		if err != nil{
			fmt.Println(err.Error())
		}
		numRows, _ := res.RowsAffected()
		if numRows != 0 {
			fmt.Println("Successfully deleted task")
		}
		task.PrintList()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().Bool("all", false, "Delete all task from todo tablse")
}

