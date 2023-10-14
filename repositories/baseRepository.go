package repositories

import (
	"log"
	"mekari/models"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := gorm.Open("postgres", "host="+dbHost+" user="+username+" password="+password+" dbname="+dbName+" port="+dbPort+" sslmode=disable TimeZone=Asia/Jakarta")

	if err != nil {

		log.Println(err)
	}
	err = conn.DB().Ping()
	if err != nil {
		log.Println(err)
	}

	db = conn
	if os.Getenv("mode") == "development" {
		db.LogMode(true)
	}
	db.Debug().AutoMigrate(
		&models.Employee{})
}

func GetDB() *gorm.DB {
	modeSeeder, err := strconv.ParseBool(os.Getenv("MODE_SEEDER"))
	if err != nil {
		log.Println(err.Error())
	}
	//seedeer
	if modeSeeder == true {

	}
	return db
}
