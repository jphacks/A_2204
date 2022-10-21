package controllers

import (
	"dietApp/operateDb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SampleJSON struct {
	Message string `JSON:"message"`
}

// =========== /user/meals ===========
// GET /user/meals
func GET_user_meals(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// GET /user/meals/:id
func GET_user_meals_id(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// POST /user/meals
func POST_user_meals(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// DELETE /user/meals/:id
func DELETE_user_meals_id(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// PUT /user/meals/:id
func PUT_user_meals_id(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// =========== /user/weights ============
// GET /user/weights
func GET_user_weights(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// GET /user/weights/:id
func GET_user_weights_id(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// POST /user/weights
func POST_user_weights(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// DELETE /user/weights/:id
func DELETE_user_weights_id(c echo.Context) error {
	id := c.Param("id")
	var int_id int
	//stringからintにキャスト
	int_id, _ = strconv.Atoi(id)
	//構造体を読み込む
	u := new(operateDb.User_weight)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	u.Id = int_id
	db := operateDb.GetConnect()
	// まずは削除したいレコードの情報を埋める
	db.First(&u)
	// 完全にレコードを特定できる状態で削除を行う
	db.Delete(&u)
	return c.JSON(http.StatusOK, u)
}

// PUT /user/weights/:id
func PUT_user_weights_id(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// =========== /user/character ============

// GET /user/character
func GET_user_character(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// PUT /user/character
func PUT_user_character(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}
