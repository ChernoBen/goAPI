package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	//especificando tipo e nome do banco de dados
	database, err := gorm.Open("sqlite3", "test.tb")
	if err != nil {
		panic("Connection fail!")
	}
	database.AutoMigrate(&Book{})
	database.AutoMigrate(&User{})
	//DB sera usado para obter acesso ao banco nos controllers
	DB = database
}
