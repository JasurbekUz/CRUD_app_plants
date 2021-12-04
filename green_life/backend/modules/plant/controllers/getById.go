package controllers

import (
	"strconv"
	m "backend/modules/plant/modules"
	"github.com/gin-gonic/gin"
)

func GetById (ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	result, ok := m.SelectById(id)

	if ok {

		ctx.JSON(200, result)

	} else {

		ctx.JSON(404, nil)
	}
}