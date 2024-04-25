package controllers

import (
	"compel/helpers"
	"compel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInstrument(router *gin.Engine) {
	routerMenu := router.Group("/api/instruments/")

	routerMenu.GET("/get-all/", GetInstruments)
	routerMenu.GET("/get-random/", GetRandomInstrument)
}

func GetInstruments(c *gin.Context) {
	var instruments []models.CatInstruments
	err := models.DB.Find(&instruments).Error
	if err != nil {
		c.JSON(helpers.ResponseGeneral(c, "Error to read data", http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, instruments)
}

func GetRandomInstrument(c *gin.Context) {
	var instrument models.CatInstruments
	err := models.DB.Take(&instrument).Error
	if err != nil {
		c.JSON(helpers.ResponseBadRequest(c, "Error to get data"))
		return
	}
	c.JSON(http.StatusOK, instrument)
}

func AddInstruments(c *gin.Context) {
	var body []string
	c.Bind(&body)
	err := models.AddInstruments(body)
}
