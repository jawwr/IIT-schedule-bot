package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/services"
	"strconv"
)

type ScheduleController struct {
	RequestMapping string
	Service        services.ScheduleService
}

func (c ScheduleController) BindEndpoints(server *gin.Engine) {
	server.GET(c.RequestMapping+"/today", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.Service.GetToday())
	})
	server.GET(c.RequestMapping+"/week", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.Service.GetWeek())
	})
	server.GET(c.RequestMapping+"/week/:number", func(context *gin.Context) {
		weekNumberStr, _ := context.Params.Get("number")
		weekNumber, err := strconv.Atoi(weekNumberStr)
		if err != nil {
			log.Fatal(err)
		}
		context.JSON(http.StatusOK, c.Service.GetWeekByNumber(weekNumber))
	})
}
