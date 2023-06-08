package mainServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/controllers"
)

type MainServer struct {
	ginServer   *gin.Engine
	Port        int
	controllers []controllers.RestController
}

func New(port int) *MainServer {
	return &MainServer{
		ginServer: gin.Default(),
		Port:      port,
		controllers: []controllers.RestController{
			&controllers.ScheduleController{
				RequestMapping: "/schedule",
			},
		},
	}
}

func (s MainServer) bindEndpoints() {
	for _, controller := range s.controllers {
		controller.BindEndpoints(s.ginServer)
	}
}

func (s MainServer) Start() {
	s.bindEndpoints()
	if err := s.ginServer.Run(fmt.Sprintf(":%d", s.Port)); err != nil {
		log.Fatal(err)
	}
}