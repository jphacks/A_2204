package operateDb

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// データベースに入れるtable
type User struct {
	Id       int       `json:"id"`
	Auth0_id string    `json:"auth0_id"`
	Height   float64   `json:"height"`
	Birthday time.Time `json:"birthday"`
}
type User_meal struct {
	Id      int       `json:"id"`
	User_id int       `json:"user_id"`
	Name    string    `json:"name"`
	Calorie int       `json:"calorie"`
	At      time.Time `json:"at"`
}
type User_weight struct {
	Id      int       `json:"id"`
	User_id int       `json:"user_id"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"data"`
}
type Character struct {
	User_id int       `json:"user_id"`
	Name    string    `json:"name"`
	Level   int       `json:"level"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"at"`
	Exp     int       `json:"exp"`
}

func Init() {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	//データベースを開ける
	dbCon, err := gorm.Open(DBMS, CONNECT)
	//ローカル変数のdbCOnをグローバル変数のdbに入れる
	db = dbCon
	//作成するtableの構造体を引数に入れてテーブルを作る
	db.AutoMigrate(&User{}, &User_meal{}, &User_weight{}, &Character{})
	if err != nil {
		panic(err.Error())
	}
}

func GetConnect() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
