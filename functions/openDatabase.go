package functions

import (
	"database/sql"
	"log"

	"../globalVars"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDatabase() {
	var err error
	globalVars.DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shortink")

	if err != nil {
		log.Fatal("Database is not open")
		return
	}

	err = globalVars.Links.FillFromDatabase(globalVars.DB)
	if err != nil {
		log.Fatal("Data not loaded")
		return
	}
}

