package main

import (
	"log"
	"mekari/repositories"
	"mekari/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	port := os.Getenv("port")

	if port == "" {
		port = "8000" //default port
	}

	err := repositories.GetDB()
	if err != nil {
		log.Println(err)
	}

	r := routers.SetupRouter()
	_ = r.Run(":" + port)

}
