package main

import (
	"alta-store/configs"
	"alta-store/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	configs.InitDb()
	configs.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configs.HTTP_PORT)))
}
