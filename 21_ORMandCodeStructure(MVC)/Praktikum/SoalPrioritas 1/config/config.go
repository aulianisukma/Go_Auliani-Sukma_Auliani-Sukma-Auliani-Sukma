package cpnfig

import (
	"fmt"
	"praktikum/models"

	"static-api-crud-echo/go/pkg/mod/github.com/jinzhu/gorm@v1.9.16"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {

	DB_Username := "root"
	DB_Password := "102jatihurip"
	DB_Port := "3306"
	DB_Host := "localhost"
	DB_Name := ""

	connectionString := fmt.Sprint("%s:%s@tcp(%s:%s)/%s?charset+utf8&parseTime=True&loc+Local",
		Db_Username,
		DB_Password,
		DB_Host,
		DB_Port,
		DB_Name,
)

		var err error
		DB, err = gorm.Open(mysql.open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.info),
		})
		if err != nil {
		panic(err)
		}

{

func InitialMigration() {
	DB.AutoMigrate(&models.Users{}, &models.Books{}, &models.Blogs{})	
}
}