package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScheduleController struct {
	RequestMapping string
}

func (c ScheduleController) BindEndpoints(server *gin.Engine) {
	server.GET(c.RequestMapping+"/today", func(context *gin.Context) {
		context.JSON(http.StatusOK, c.getToday())
	})
}

func (c ScheduleController) getToday() string {
	return "get today request"
}
