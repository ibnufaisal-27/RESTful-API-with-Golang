package main

import (
	"log"
	"mekari/repositories"
	"mekari/routers"
	"os"

	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "mekari/docs"
)

// @title RESTFul API Service
// @version 1.0
// @description This is a repository of RESTFul API service.
// @contact.name Muhammad Ibn Faisal
// @contact.url https://github.com/ibnufaisal-27/mekari
// @contact.email mibnufaisal@gmail.com

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	_ = r.Run(":" + port)

}
