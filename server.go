package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/lib/pq"
)

type Expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func createExpenseHandler(c echo.Context) error {
	e := Expense{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", e.Title, e.Amount, e.Note, pq.Array(e.Tags))
	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}

func getExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query expense statment:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	e := Expense{}
	err = row.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expense not found"})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expense:" + err.Error()})
	}
}

func updateExpenseHandler(c echo.Context) error {
	id := c.Param("id")

	e := Expense{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1 RETURNING id;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare statment update:" + err.Error()})
	}

	if err := stmt.QueryRow(id, e.Title, e.Amount, e.Note, pq.Array(e.Tags)).Scan(&e.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "error execute update:" + err.Error()})
	}

	return c.JSON(http.StatusOK, e)

}

func getExpensesHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses statment:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
	}

	expenses := []Expense{}

	for rows.Next() {
		e := Expense{}
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expense:" + err.Error()})
		}
		expenses = append(expenses, e)
	}

	return c.JSON(http.StatusOK, expenses)
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

	e.POST("/expenses", createExpenseHandler)
	e.GET("/expenses/:id", getExpenseHandler)
	e.PUT("/expenses/:id", updateExpenseHandler)
	e.GET("/expenses", getExpensesHandler)

	log.Fatal(e.Start(":2565"))

	// fmt.Println("Please use server.go for main file")
	// fmt.Println("start at port:", os.Getenv("PORT"))
}
