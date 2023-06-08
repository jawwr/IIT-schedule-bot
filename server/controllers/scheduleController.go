package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/services"
)

type ScheduleController struct {
	RequestMapping string
	Service        services.ScheduleService
}

func (c ScheduleController) BindEndpoints(server *gin.Engine) {
	server.GET(c.RequestMapping+"/today", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.Service.GetToday())
	})
}
