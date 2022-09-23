package main

import (
	"adm/user"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		log.Fatal(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password) //Создать строку подключения
	fmt.Println(dbUri)

	var err error
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()

	user.RegisterUserHandlers(router, db)

}
