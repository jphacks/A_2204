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
	Id       int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Auth0_id string    `json:"auth0_id" gorm:"unique"`
	Height   float64   `json:"height"`
	Birthday time.Time `json:"birthday"`
}
type User_meal struct {
	Id      int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id int       `json:"user_id"`
	User    User      `gorm:"foreignKey:User_id"`
	Name    string    `json:"name"`
	Calorie int       `json:"calorie"`
	At      time.Time `json:"at"`
}
type User_weight struct {
	Id      int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id int       `json:"user_id"`
	User    User      `gorm:"foreignKey:User_id"`
	Weight  float64   `json:"weight"`
	At      time.Time `json:"at"`
}
type Character struct {
	User_id int    `json:"user_id" gorm:"primaryKey"`
	User    User   `gorm:"foreignKey:User_id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Exp     int    `json:"exp"`
}

func Init() {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	//データベースを開ける
	dbCon, err := gorm.Open(DBMS, CONNECT)
	//ローカル変数のdbCOnをグローバル変数のdbに入れる
	db = dbCon
	//作成するtableの構造体を引数に入れてテーブルを作る
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User_meal{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&User_weight{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&Character{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

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
