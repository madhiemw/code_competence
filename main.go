package main

import (
	"codecompetence/config"
	routes "codecompetence/route"

	"github.com/labstack/echo"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	config.InitDB()
	// middleware.Logmiddleware(e)

	routes.Route(e, db)
	e.Logger.Fatal(e.Start(":2222"))
}
