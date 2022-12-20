package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://tgrpfnnv:GCDxLHXUGuHxWXtQzc14KwQS4jCSnUgK@tiny.db.elephantsql.com/tgrpfnnv"
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	log.Println("Connect to database success")

	createTb := `CREATE TABLE IF NOT EXISTS expenses ( 
		id SERIAL PRIMARY KEY, 
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")
}
