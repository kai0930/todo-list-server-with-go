package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-list-server-with-go/internal/model"
)

func GetGroup(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	groups, err := model.GetGroups(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get groups")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"groups": groups,
	})
}

func PostGroup(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	var req model.GroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	group, err := model.CreateGroup(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create group")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"group": group,
	})
}

func PutGroup(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	var req model.GroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	groupId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid group id")
	}

	group, err := model.PutGroup(id, groupId, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update group")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"group": group,
	})
}

func DeleteGroup(c echo.Context) error {
	id, err := GetIdFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	groupId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid group id")
	}
	err = model.DeleteGroup(id, groupId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete group")
	}
	return c.JSON(http.StatusOK, "Group deleted")
}
