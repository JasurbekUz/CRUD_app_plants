package controllers

import (
	m "backend/modules/category/modules"
	"github.com/gin-gonic/gin"
)

func GetAll (ctx *gin.Context) {

	ctx.JSON(200, m.SelectAll())
}