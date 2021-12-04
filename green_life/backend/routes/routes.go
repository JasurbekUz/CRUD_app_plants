package routes

import (
	m "backend/middlewares"
	ctg "backend/modules/category/controllers"
	pl "backend/modules/plant/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes (r *gin.Engine) {

	r.Use(m.Cors())

	// CATEGORIES
	categories := r.Group("/categories")

	categories.POST("/new", ctg.Post)
	categories.GET("/", ctg.GetAll)
	categories.GET("/:id", ctg.GetById)
	categories.PUT("/edit/:id", ctg.Put)
	categories.DELETE("/rm/:id", ctg.Delete)
	
	// PLANTS
	plants := r.Group("/plants")
	
	plants.POST("/new", pl.Post)
	plants.GET("/", pl.GetAll)
	plants.GET("/:id", pl.GetById)
	plants.PUT("/edit/:id", pl.Put)
	plants.DELETE("/rm/:id", pl.Delete)
	
}