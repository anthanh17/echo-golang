package main

import (
	"goEcho/db"
	"goEcho/handler"
	"goEcho/repository/repo_impl"
	"goEcho/router"
	"os"

	log "goEcho/log"

	"github.com/labstack/echo/v4"
)

// func first call when package loaded then call func main
func init() {
	// fmt.Println("Init package main")
	os.Setenv("APP_NAME", "github_annt_echo_golang")
	log.InitLogger(false)
}

func main() {
	// fmt.Println("main function")

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

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.Api{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
