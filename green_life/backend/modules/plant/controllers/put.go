package controllers

import (
	"fmt"
	"strconv"
	m "backend/modules/plant/modules"
	. "backend/modules/plant/models"
	"github.com/gin-gonic/gin"
)

func Put (ctx *gin.Context) {

	var plant PostPlant

	ctx.BindJSON(&plant)

	id, _ := strconv.Atoi(ctx.Param("id"))

	result, ok := m.UpdatePlant(id, plant)

	fmt.Println("/////////////////", plant)

	if ok {

		ctx.JSON(200, result)

	} else {

		ctx.JSON(404, nil)
	}
}