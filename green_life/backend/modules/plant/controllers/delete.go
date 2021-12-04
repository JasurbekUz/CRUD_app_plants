package controllers

import (
	"strconv"
	m "backend/modules/plant/modules"
	"github.com/gin-gonic/gin"
)

func Delete (ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	result, ok := m.DeletePlant(id)

	if ok {

		ctx.JSON(200, gin.H{

			"deleted": result,
		})

	} else {

		ctx.JSON(404, nil)
	}
}