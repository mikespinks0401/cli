package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mikespinks0401/cobra-todo/db"
)

type Task struct{
	Task string
	Created_at string
	Done int
	Updated_at sql.NullString
}

func GetList()([]Task, error){
	conn, err := db.DBConn()
	if err != nil{
		panic(err)
	}	
	defer conn.Close()
	rows, err := conn.Query("SELECT Task, Created_at, Done, Updated_at FROM todo")
	if err != nil {
		return nil, rows.Err()
	}
	var tasks[]Task
	for rows.Next(){
		var task Task
		if err := rows.Scan(&task.Task, &task.Created_at, &task.Done, &task.Updated_at ); err != nil{
			fmt.Println(err.Error())
		}
		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no task exist in todo")
	}
	return tasks, nil
}

func AddTask(task string)error{
	if task == ""{
		return fmt.Errorf("must provide a task to add")
	}
	conn, err := db.DBConn()
	if err != nil {
		return err
	}	
	defer conn.Close()
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	q := `INSERT INTO todo(task, created_at) VALUES(?,?)`
	stmt,err := conn.Prepare(q)
	if err != nil{
		fmt.Println(err.Error())
	}
	res, err:= stmt.Exec(task, dateTime)
	if err != nil {
		fmt.Println(err)
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("there was an error retrieving your task")
	}
	row := conn.QueryRow("SELECT (task) FROM todo WHERE id = ?", id)
	var lastTask string
	row.Scan(&lastTask)
	fmt.Println("Successfully Inserted Task:", lastTask)
	return nil
}

func PrintList()error{
	list, err := GetList()
	if err != nil{
		return fmt.Errorf("error:%s", err.Error())
	}
	for i, val := range list{
		fmt.Printf("%d. %s\n", i + 1, val.Task)
	}
	return nil
}