package controllers

import (
	m "backend/modules/category/modules"
	. "backend/modules/category/models"
	"github.com/gin-gonic/gin"
)

func Post (ctx *gin.Context) {

	var newCategory PostCategory

	ctx.BindJSON(&newCategory)

	ctx.JSON(201, m.InsertCategory(newCategory))
}