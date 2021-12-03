package controllers

import (
	"strconv"
	m "backend/modules/category/modules"
	. "backend/modules/category/models"
	"github.com/gin-gonic/gin"
)

func Put (ctx *gin.Context) {

	var category Category

	ctx.BindJSON(&category)

	id, _ := strconv.Atoi(ctx.Param("id"))

	result, ok := m.UpdateCategory(id,category)

	if ok {

		ctx.JSON(200, result)

	} else {

		ctx.JSON(404, nil)
	}
}