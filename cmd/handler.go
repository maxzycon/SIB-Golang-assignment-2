package cmd

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxzycon/SIB-Golang-Assigment-2/pkg/dto"
	"gorm.io/gorm"
)

type Controller struct {
	service Service
}

func (controller *Controller) GetAll(c *gin.Context) {
	data, err := controller.service.GetAll()
	if err != nil {
		generateResponse(c, 404, gin.H{
			"message": "error get all order",
		})
	}

	generateResponse(c, 200, gin.H{
		"Data": data,
	})
}

func (controller *Controller) DeleteOrder(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		generateResponse(c, 404, gin.H{
			"result": "err convert params to int",
		})
		return
	}

	err = controller.service.DeleteOrder(uint(id))
	if err != nil && err == gorm.ErrRecordNotFound {
		generateResponse(c, 404, gin.H{
			"result": fmt.Sprintf("Order with id %s Not Found", param),
		})
		return
	}

	if err != nil {
		generateResponse(c, 500, gin.H{
			"result": "create model error",
		})
		return
	}

	generateResponse(c, 200, gin.H{
		"result": "success",
	})
}

func (controller Controller) CreateOrder(c *gin.Context) {
	dto := dto.PaloadOrder{}
	err := c.BindJSON(&dto)
	if err != nil {
		generateResponse(c, 400, gin.H{
			"message": "bad request",
		})
		return
	}
	err = controller.service.CreateOrder(&dto)
	if err != nil {
		generateResponse(c, 500, gin.H{
			"result": "create model error",
		})
		return
	}

	generateResponse(c, 200, gin.H{
		"result": "ok",
	})
}

func (controller Controller) UpdateOrder(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		generateResponse(c, 404, gin.H{
			"result": "err convert params to int",
		})
		return
	}

	dto := dto.PaloadOrder{}
	err = c.BindJSON(&dto)
	if err != nil {
		generateResponse(c, 400, gin.H{
			"message": "bad request",
		})
		return
	}

	data, err := controller.service.UpdateOrder(uint(id), &dto)
	if err != nil {
		generateResponse(c, 500, gin.H{
			"result": "create model error",
		})
		return
	}

	generateResponse(c, 200, gin.H{
		"result": data,
	})
}

func generateResponse(c *gin.Context, code int, resp interface{}) {
	c.JSON(code, resp)
}
