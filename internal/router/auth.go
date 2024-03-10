package router

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"todo-list-server-with-go/internal/model"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req LoginRequest

	// リクエストボディからemailとpasswordを取得
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// ログイン処理を実行
	account, err := model.Login(req.Email, req.Password)
	if err != nil {
		// ログイン失敗時の処理
		return c.JSON(http.StatusUnauthorized, "Login failed")
	}

	// ログイン成功時の処理
	// custom claimを作成
	claims := &model.JwtCustomClaims{
		account.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	// JWTを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// JWTを署名
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to sign token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   signedToken,
		"account": account,
	})
}

func SignUp(c echo.Context) error {
	var req SignupRequest

	// リクエストボディからemailとpasswordを取得
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// サインアップ処理を実行
	account, err := model.Signup(req.Email, req.Password)
	if err != nil {
		// サインアップ失敗時の処理
		return c.JSON(http.StatusBadRequest, "Signup failed")
	}

	// サインアップ成功時の処理
	// custom claimを作成
	claims := &model.JwtCustomClaims{
		account.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	// JWTを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// JWTを署名
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to sign token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   signedToken,
		"account": account,
	})
}
