package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/ken-harada/household-accounts/routes"
	"github.com/ken-harada/household-accounts/template"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "household-accounts",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "mariadb:3306",
		DBName:               "household-accounts",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	e := echo.New()
	e.Renderer = template.NewTemplate()
	routes.Route(e)

	e.Logger.Fatal(e.Start(":8000"))
}
