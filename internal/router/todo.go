package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-list-server-with-go/internal/model"
)

func GetTodo(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	// クエリパラメータからgroupIdを取得
	groupId, err := uuid.Parse(c.QueryParam("groupId"))
	todos, err := model.GetTodos(id, groupId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get todos")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"todos": todos,
	})
}

func PostTodo(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	var req model.TodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	todo, err := model.CreateTodo(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create todo")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"todo": todo,
	})
}

func PutTodo(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	var req model.TodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	todoId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid todo id")
	}

	todo, err := model.PutTodo(id, todoId, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update todo")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"todo": todo,
	})
}

func DeleteTodo(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	todoId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid todo id")
	}
	err = model.DeleteTodo(id, todoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete todo")
	}
	return c.JSON(http.StatusOK, "Deleted todo")
}
