package controllers

import (
	m "backend/modules/plant/modules"
	. "backend/modules/plant/models"
	"github.com/gin-gonic/gin"
)

func Post (ctx *gin.Context) {

	var newPlant PostPlant
	ctx.BindJSON(&newPlant)

	ctx.JSON(201, m.InsertPlant(newPlant))
}