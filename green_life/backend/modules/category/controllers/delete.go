package controllers

import (
	"strconv"
	m "backend/modules/category/modules"
	"github.com/gin-gonic/gin"
)

func Delete (ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	result, ok := m.DeleteCtg(id)

	if ok {

		ctx.JSON(200, result)

	} else {

		ctx.JSON(404, nil)
	}
}