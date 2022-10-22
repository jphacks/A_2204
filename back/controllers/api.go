package controllers

import (
	"dietApp/operateDb"
	"net/http"
	"strconv"

	"github.com/auth0/go-jwt-middleware/v2/validator"
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
	claims := c.Get("claims").(*validator.ValidatedClaims)
	auth0_id := claims.RegisteredClaims.Subject
	id := c.Param("id")
	var int_id int
	//stringからintにキャスト
	int_id, _ = strconv.Atoi(id)
	//構造体を読み込む
	u := new(operateDb.User_meal)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	db := operateDb.GetConnect()
	u.Id = int_id
	db.Where("name = ?", auth0_id).First(&u)
	//updata
	db.Model(&u).Update(&u)
	return c.JSON(http.StatusOK, u)
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
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// PUT /user/weights/:id
func PUT_user_weights_id(c echo.Context) error {
	claims := c.Get("claims").(*validator.ValidatedClaims)
	auth0_id := claims.RegisteredClaims.Subject
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
	db.Where("name = ?", auth0_id).First(&u)
	//updata
	db.Model(&u).Update(&u)
	return c.JSON(http.StatusOK, u)
}

// =========== /user/character ============

// GET /user/character
func GET_user_character(c echo.Context) error {
	return c.JSON(http.StatusOK, SampleJSON{"Coming soon"})
}

// PUT /user/character
func PUT_user_character(c echo.Context) error {
	//構造体を読み込む
	u := new(operateDb.Character)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	db := operateDb.GetConnect()
	//updata
	db.Model(&u).Update(&u)
	return c.JSON(http.StatusOK, u)
}
