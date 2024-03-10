package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTodo(c echo.Context) error {
	// "GET TODO"という文字列を返す
	return c.String(http.StatusOK, "GET TODO")
}
