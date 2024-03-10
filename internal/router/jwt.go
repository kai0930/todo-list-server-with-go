package router

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"todo-list-server-with-go/internal/model"
)

func GetIdFromJWT(c echo.Context) (uuid.UUID, error) {
	// JWTを取得
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	id := claims.ID
	return id, nil
}
