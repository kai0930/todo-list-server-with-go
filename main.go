package main

import (
	"github.com/labstack/echo/v4"
	"todo-list-server-with-go/internal/model"
	"todo-list-server-with-go/internal/router"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
