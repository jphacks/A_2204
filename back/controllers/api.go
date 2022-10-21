package controllers

import (
	"dietApp/operateDb"
	"net/http"
	"strconv"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

type SampleJSON struct {
	Message string `JSON:"message"`
}

type User_meal_res struct {
	Id      int       `json:"id"`
	User_id int       `json:"user_id"`
	Name    string    `json:"name"`
	Calorie int       `json:"calorie"`
	At      time.Time `json:"at"`
}

type User_weight_res struct {
	Id      int       `json:"id"`
	User_id int       `json:"user_id"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"at"`
}

// =========== /user/meals ===========
// GET /user/meals
func GET_user_meals(c echo.Context) error {
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	db = db.Where("user_id = ?", user.Id)

	before, err := time.Parse(time.RFC3339Nano, c.QueryParam("before"))
	if err == nil {
		db = db.Where("at < ?", before)
	}
	after, err := time.Parse(time.RFC3339Nano, c.QueryParam("after"))
	if err == nil {
		db = db.Where("? < at", after)
	}

	name := c.QueryParam("name")
	if name != "" {
		db = db.Where("name = ?", name)
	}

	calorieMin, err := strconv.Atoi(c.QueryParam("calorie_min"))
	if err == nil {
		db = db.Where("? < calorie", calorieMin)
	}
	calorieMax, err := strconv.Atoi(c.QueryParam("calorie_max"))
	if err == nil {
		db = db.Where("calorie < ?", calorieMax)
	}

	userMeals := []User_meal_res{}
	db.Table("user_meals").Find(&userMeals)

	return c.JSON(http.StatusOK, userMeals)
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
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	db = db.Where("user_id = ?", user.Id)

	before, err := time.Parse(time.RFC3339Nano, c.QueryParam("before"))
	if err == nil {
		db = db.Where("at < ?", before)
	}
	after, err := time.Parse(time.RFC3339Nano, c.QueryParam("after"))
	if err == nil {
		db = db.Where("? < at", after)
	}

	weightMin, err := strconv.ParseFloat(c.QueryParam("weight_min"), 64)
	if err == nil {
		db = db.Where("? < weight", weightMin)
	}
	weightMax, err := strconv.ParseFloat(c.QueryParam("weight_max"), 64)
	if err == nil {
		db = db.Where("calorie < ?", weightMax)
	}

	userWeights := []User_weight_res{}
	db.Table("user_weights").Find(&userWeights)

	return c.JSON(http.StatusOK, userWeights)
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
