package commons

import (
	"crudapirest/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/testgo")
	if err!=nil{
		log.Fatal(err)
	}
	return db
}

func Migrate()  {
	db:= GetConnection()
	defer db.Close()
	log.Println("Migrando....")
	db.AutoMigrate(&models.Seller{}, &models.Buyer{})
}
