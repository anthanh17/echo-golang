package main

import (
	"goEcho/db"
	"goEcho/handler"
	"runtime"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

	logErr("Test co lôi xay ra")

	e := echo.New()

	e.GET("/", handler.Welcome)

	e.Logger.Fatal(e.Start(":3000"))
}

func logErr(errMsg string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Fatal(errMsg)
}
