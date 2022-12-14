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

type User_req struct {
	Id       string  `json:"id"`
	Auth0_id string  `json:"auth0_id"`
	Height   float64 `json:"height"`
	Birthday int64   `json:"birthday"`
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

type Character_res struct {
	User_id int    `json:"user_id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Exp     int    `json:"exp"`
}

// =========== /user ===========
// GET /user
func GET_user(c echo.Context) error {
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Where("Auth0_id = ?", claims.RegisteredClaims.Subject).First(&user)

	return c.JSON(http.StatusOK, user)
}

// POST /user
func POST_user(c echo.Context) error {
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User

	user.Auth0_id = claims.RegisteredClaims.Subject
	reqBody := new(User_req)
	if err := c.Bind(reqBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	user.Height = reqBody.Height
	user.Birthday = reqBody.Birthday

	if err := db.Create(&user).Error; err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	character := &operateDb.Character{
		User_id: user.Id,
		Name:    "",
		Level:   1,
		Exp:     0,
	}

	if err := db.Create(&character).Error; err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// PUT /user
func PUT_user(c echo.Context) error {
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Where("Auth0_id = ?", claims.RegisteredClaims.Subject).First(&user)

	reqBody := new(User_req)
	if err := c.Bind(reqBody); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user.Height = reqBody.Height
	user.Birthday = reqBody.Birthday

	if err := db.Table("users").Where("id = ?", user.Id).Updates(operateDb.User{Height: user.Height, Birthday: user.Birthday}).Error; err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqBody)
}

// =========== /user/meals ===========
// GET /user/meals
func GET_user_meals(c echo.Context) error {
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	db = db.Where("user_id = ?", user.Id)

	before, err := strconv.ParseInt(c.QueryParam("before"), 10, 64)
	if err == nil {
		db = db.Where("at <= ?", before)
	}
	after, err := strconv.ParseInt(c.QueryParam("after"), 10, 64)
	if err == nil {
		db = db.Where("? <= at", after)
	}

	name := c.QueryParam("name")
	if name != "" {
		db = db.Where("name = ?", name)
	}

	calorieMin, err := strconv.Atoi(c.QueryParam("calorie_min"))
	if err == nil {
		db = db.Where("? <= calorie", calorieMin)
	}
	calorieMax, err := strconv.Atoi(c.QueryParam("calorie_max"))
	if err == nil {
		db = db.Where("calorie <= ?", calorieMax)
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
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	//????????????????????????
	u := new(operateDb.User_meal)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	u.User_id = user.Id
	//DB???insert?????????
	if err := db.Create(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, u)
}

// DELETE /user/meals/:id
func DELETE_user_meals_id(c echo.Context) error {
	id := c.Param("id")
	var int_id int
	//string??????int???????????????
	int_id, _ = strconv.Atoi(id)
	//????????????????????????
	u := new(operateDb.User_meal)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	u.Id = int_id
	db := operateDb.GetConnect()
	// ????????????????????????????????????????????????
	db.First(&u)
	// ???????????????????????????????????????????????????????????????
	db.Delete(&u)
	return c.JSON(http.StatusOK, u)
}

// PUT /user/meals/:id
func PUT_user_meals_id(c echo.Context) error {
	claims := c.Get("claims").(*validator.ValidatedClaims)
	auth0_id := claims.RegisteredClaims.Subject
	id := c.Param("id")
	var int_id int
	//string??????int???????????????
	int_id, _ = strconv.Atoi(id)
	//????????????????????????
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
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	db = db.Where("user_id = ?", user.Id)

	before, err := strconv.ParseInt(c.QueryParam("before"), 10, 64)
	if err == nil {
		db = db.Where("at <= ?", before)
	}
	after, err := strconv.ParseInt(c.QueryParam("after"), 10, 64)
	if err == nil {
		db = db.Where("? <= at", after)
	}

	weightMin, err := strconv.ParseFloat(c.QueryParam("weight_min"), 64)
	if err == nil {
		db = db.Where("? <= weight", weightMin)
	}
	weightMax, err := strconv.ParseFloat(c.QueryParam("weight_max"), 64)
	if err == nil {
		db = db.Where("weight <= ?", weightMax)
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
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	//????????????????????????
	u := new(operateDb.User_weight)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	u.User_id = user.Id

	//DB???insert?????????
	if err := db.Create(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, u)
}

// DELETE /user/weights/:id
func DELETE_user_weights_id(c echo.Context) error {
	id := c.Param("id")
	var int_id int
	//string??????int???????????????
	int_id, _ = strconv.Atoi(id)
	//????????????????????????
	u := new(operateDb.User_weight)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	u.Id = int_id
	db := operateDb.GetConnect()
	// ????????????????????????????????????????????????
	db.First(&u)
	// ???????????????????????????????????????????????????????????????
	db.Delete(&u)
	return c.JSON(http.StatusOK, u)
}

// PUT /user/weights/:id
func PUT_user_weights_id(c echo.Context) error {
	claims := c.Get("claims").(*validator.ValidatedClaims)
	auth0_id := claims.RegisteredClaims.Subject
	id := c.Param("id")
	var int_id int
	//string??????int???????????????
	int_id, _ = strconv.Atoi(id)
	//????????????????????????
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
	db := operateDb.GetConnect()
	claims := c.Get("claims").(*validator.ValidatedClaims)
	var user operateDb.User
	db.Model(&operateDb.User{Auth0_id: claims.RegisteredClaims.Subject}).First(&user)

	db = db.Where("user_id = ?", user.Id)

	character := &Character_res{}
	db.Table("characters").First(character)
	return c.JSON(http.StatusOK, character)
}

// PUT /user/character
func PUT_user_character(c echo.Context) error {
	//????????????????????????
	u := new(operateDb.Character)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	db := operateDb.GetConnect()
	//updata
	db.Model(&u).Update(&u)
	return c.JSON(http.StatusOK, u)
}
