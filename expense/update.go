package expense

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
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
