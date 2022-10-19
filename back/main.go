package main

import (
	"dietApp/controllers"
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

	// /api/user/meals
	api.GET("/user/meals", controllers.GET_user_meals)
	api.GET("/user/meals/:id", controllers.GET_user_meals_id)
	api.POST("/user/meals", controllers.POST_user_meals)
	api.DELETE("/user/meals/:id", controllers.DELETE_user_meals_id)
	api.PUT("/user/meals/:id", controllers.PUT_user_meals_id)

	// /api/user/weights
	api.GET("/user/weights", controllers.GET_user_weights)
	api.GET("/user/weights/:id", controllers.GET_user_weights_id)
	api.POST("/user/weights", controllers.POST_user_weights)
	api.DELETE("/user/weights/:id", controllers.DELETE_user_weights_id)
	api.PUT("/user/weights/:id", controllers.PUT_user_weights_id)

	// /api/user/character
	api.GET("/user/character", controllers.GET_user_character)
	api.PUT("/user/character", controllers.PUT_user_character)

	// ===================================================

	// Hello Worldを返すエンドポイントの作成
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
