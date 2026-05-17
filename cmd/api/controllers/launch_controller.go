package controllers

import (
	models "github.com/Raphsodyz/spacedevs-go/models"
	"github.com/Raphsodyz/spacedevs-go/usecase"
	"github.com/gin-gonic/gin"
)

type LaunchController struct {
	searchUC usecase.SearchLaunchUseCase
}

func NewLaunchController(searchUC usecase.SearchLaunchUseCase) *LaunchController {
	return &LaunchController{
		searchUC: searchUC,
	}
}

func (lc *LaunchController) SearchLaunch(c *gin.Context) {
	var req models.SearchLaunchRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := lc.searchUC.SearchByRequest(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func RegisterRoutes(router *gin.Engine, controller *LaunchController) {
	v1 := router.Group("/api/v1")
	launch := v1.Group("/launch")
	launch.GET("/search", controller.SearchLaunch)
}
