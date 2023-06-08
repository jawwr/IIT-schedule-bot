package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/services"
	"strconv"
)

type ScheduleController struct {
	requestMapping string
	Service        services.ScheduleService
}

func (c ScheduleController) BindEndpoints(server *gin.Engine) {
	server.GET(c.requestMapping+"/today", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.Service.GetToday())
	})
	server.GET(c.requestMapping+"/week", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.Service.GetWeek())
	})
	server.GET(c.requestMapping+"/week/:number", func(context *gin.Context) {
		weekNumberStr, _ := context.Params.Get("number")
		weekNumber, err := strconv.Atoi(weekNumberStr)
		if err != nil {
			log.Fatal(err)
		}
		context.JSON(http.StatusOK, c.Service.GetWeekByNumber(weekNumber))
	})
}

func NewScheduleController(mapping string, service services.ScheduleService) *ScheduleController {
	return &ScheduleController{
		requestMapping: mapping,
		Service:        service,
	}
}
