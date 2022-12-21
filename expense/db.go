package expense

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://tgrpfnnv:GCDxLHXUGuHxWXtQzc14KwQS4jCSnUgK@tiny.db.elephantsql.com/tgrpfnnv")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

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
}
