package main

import (
	. "backend/routes"
	d "backend/database"
	"github.com/gin-gonic/gin"
)

func main () {

	router := gin.Default()

	d.InitDatabase()

	InitRoutes(router)

	router.Run()
}