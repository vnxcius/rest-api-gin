package database

import (
	"github.com/vnxcius/gin-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	err error
)

func Connection()  {
	dsn := "root:@tcp(localhost:3306)/api_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	// Criar uma migration de uma INSTÂNCIA do model ALUNO
	DB.AutoMigrate(&models.Aluno{})

}