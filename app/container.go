package container

import (
	"CRUD_API/app/controllers"
	"CRUD_API/infra/conn"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	db := conn.ConnectDb()
	controllers.NewSystemController(e, db)
}
