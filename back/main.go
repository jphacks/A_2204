package main

import (
	"dietApp/myMiddleware"
	"log"
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// .envファイルを読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	e := echo.New()

	api := e.Group("/api") // Auth0の認証が必要なエンドポイント

	// CORSの設定
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{os.Getenv("FRONT_URL")},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType},
	}))

	// Auth0の認証ミドルウェア
	api.Use(myMiddleware.Auth0)

	// =========== 以下にapiのルーティングを記述 ==========
	api.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims)
	})
	// ===================================================

	// Hello Worldを返すエンドポイントの作成
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
