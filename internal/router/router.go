package router

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"todo-list-server-with-go/internal/model"
)

func SetRouter(e *echo.Echo) error {
	// ミドルウェアの設定
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	})) // ログの出力
	e.Use(middleware.Recover()) // エラーが発生した場合に500エラーを返す
	e.Use(middleware.CORS())    // CORSに対応する

	// ルーティング

	// jwt認証のないルート
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to TODO App!")
	})
	publicGroup := e.Group("/auth")
	{
		publicGroup.POST("/signup", SignUp)
		publicGroup.POST("/login", Login)
	}

	// jwt認証のあるルート
	privateGroup := e.Group("/app")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(model.JwtCustomClaims)
			},
			SigningKey: []byte("secret"),
		}
		privateGroup.Use(echojwt.WithConfig(config))
		privateGroup.GET("/todo", GetTodo)
		privateGroup.POST("/todo", PostTodo)
		privateGroup.PUT("/todo/:id", PutTodo)
		privateGroup.DELETE("/todo/:id", DeleteTodo)
		privateGroup.GET("/group", GetGroup)
		privateGroup.POST("/group", PostGroup)
		privateGroup.PUT("/group/:id", PutGroup)
		privateGroup.DELETE("/group/:id", DeleteGroup)
	}

	err := e.Start(":8000")
	return err
}
