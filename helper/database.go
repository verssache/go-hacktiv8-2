package helper

import (
	"log"

	"github.com/verssache/go-hacktiv8-2/config"
	"github.com/verssache/go-hacktiv8-2/orders"
	"github.com/verssache/go-hacktiv8-2/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := cfg.Database.DBUser + ":" + cfg.Database.DBPassword + "@tcp(" + cfg.Database.DBHost + ":" + cfg.Database.DBPort + ")/" + cfg.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&orders.Order{}, &orders.Item{}, &users.User{})
	return db
}
