package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

type Expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", "postgres://tgrpfnnv:GCDxLHXUGuHxWXtQzc14KwQS4jCSnUgK@tiny.db.elephantsql.com/tgrpfnnv")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.POST("/expenses", createUserHandler)
	// e.GET("/expenses", getUsersHandler)

	log.Fatal(e.Start(":2565"))

	// fmt.Println("Please use server.go for main file")
	// fmt.Println("start at port:", os.Getenv("PORT"))
}
