package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConn *gorm.DB

func ConnectDB (){
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectStr := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(connectStr))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &UserRole{}, &Role{}, &Permission{}, &RolePermission{}, &Order{}, &OrderDetails{}, &Product{}, &Supplier{})
	MysqlConn = db
}