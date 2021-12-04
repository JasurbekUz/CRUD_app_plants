package controllers

import (
	m "backend/modules/plant/modules"
	"github.com/gin-gonic/gin"
)

func GetAll (ctx *gin.Context) {

	ctx.JSON(200, m.SelectAll())
}