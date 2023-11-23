package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzycon/SIB-Golang-Assigment-2/config"
)

func Init() {
	// ----- init mariadb
	db := config.InitMariaDb()
	// ----- ini handler, service
	serviceInit := Service{
		db: db,
	}
	controllerInit := Controller{
		service: serviceInit,
	}

	r := gin.Default()
	r.GET("/orders", controllerInit.GetAll)
	r.DELETE("/orders/:id", controllerInit.DeleteOrder)
	r.POST("/orders", controllerInit.CreateOrder)
	r.PATCH("/orders/:id", controllerInit.UpdateOrder)

	r.Run(":8005") // listen and serve on 0.0.0.0:8085
}
