package cmd

import (
	container "CRUD_API/app"
	"CRUD_API/app/https/middlewares"
	"CRUD_API/infra/logger"
	"os"

	"github.com/labstack/echo"
)

func Start() {
	e := echo.New()
	if err := middlewares.Attach(e); err != nil {
		logger.Error("error occur when attaching middlewares", err)
		os.Exit(1)
	}
	container.Init(e)
	e.Logger.Fatal(e.Start(":8081"))
}
