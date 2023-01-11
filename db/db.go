package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

 func DBConn()(*sql.DB, error){
	db, err := sql.Open("sqlite3", "todoDB.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database %v", err)
	}	
	return db, nil
 }