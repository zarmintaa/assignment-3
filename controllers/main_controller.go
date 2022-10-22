package controllers

import (
	"assignment-3/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResult(ctx *gin.Context) {
	data := helpers.GetDataJson()
	valueWater, statusWater := helpers.GetStatusWeather("water", data.Status.Water)
	valueWind, statusWind := helpers.GetStatusWeather("water", data.Status.Wind)

	helpers.IntervalFunction(helpers.UpdateDataJson, 15)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"statusWater": statusWater,
		"valueWater":  valueWater,
		"statusWind":  statusWind,
		"valueWind":   valueWind,
	})
}
