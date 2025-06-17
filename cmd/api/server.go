package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init(){
	var  err error
	db, err = sql.Open("postgres", "postgres://snet:password@postgres/sneteste?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil{
		panic(err)
	}

	fmt.Println("Database successfully")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5698"))
}