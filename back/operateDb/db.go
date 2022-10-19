package operateDb

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// データベースに入れるtable
type user struct {
	Id       int       `json:"id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	Auth0_id string    `json:"auth0_id"`
	Height   float64   `json:"height"`
	Birthday time.Time `json:"birthday"`
}
type user_meal struct {
	Id      int       `json:"id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	User_id int       `json:"user_id"`
	Name    string    `json:"name"`
	Calorie int       `json:"calorie"`
	At      time.Time `json:"at"`
}
type user_weight struct {
	Id      int       `json:"id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	User_id int       `json:"user_id"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"data"`
}
type character struct {
	User_id int       `json:"user_id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	Name    string    `json:"name"`
	Level   int       `json:"level"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"at"`
	Exp     int       `json:"exp"`
}

func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "dietapp"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	//データベースを開ける
	dbCon, err := gorm.Open(DBMS, CONNECT)
	//ローカル変数のdbCOnをグローバル変数のdbに入れる
	db = dbCon
	//作成するtableの構造体を引数に入れてテーブルを作る
	db.AutoMigrate(&user{}, &user_meal{}, &user_weight{}, &character{})
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
