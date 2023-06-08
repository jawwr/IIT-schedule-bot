package controllers

import "github.com/gin-gonic/gin"

type RestController interface {
	BindEndpoints(engine *gin.Engine)
}
