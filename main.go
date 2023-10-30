package main

import (
	"goEcho/db"
	"goEcho/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect database
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "annt",
		Password: "Abc@123456789",
		DbName:   "golang",
	}
	sql.Connect()
	// Call when main exit
	defer sql.Close()

	e := echo.New()

	e.GET("/", handler.Welcome)

	e.Logger.Fatal(e.Start(":3000"))
}
