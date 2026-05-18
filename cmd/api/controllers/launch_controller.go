package controllers

import (
	docs "github.com/Raphsodyz/spacedevs-go/docs"
	models "github.com/Raphsodyz/spacedevs-go/models"
	"github.com/Raphsodyz/spacedevs-go/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type LaunchController struct {
	searchUC usecase.SearchLaunchUseCase
}

func NewLaunchController(searchUC usecase.SearchLaunchUseCase) *LaunchController {
	return &LaunchController{
		searchUC: searchUC,
	}
}

// SearchLaunch godoc
// @Summary Search launches
// @Description Method for fuzzy search of mission, location, pad, rocket and launch.
// @Tags launch
// @Accept json
// @Produce json
// @Param Mission query string false "Mission name"
// @Param Location query string false "Location name"
// @Param Pad query string false "Pad name"
// @Param Rocket query string false "Rocket name"
// @Param Launch query string false "Launch name"
// @Param Page query int false "Page number"
// @Success 200 {object} models.SearchLaunchResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /launch/search [get]
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

	if result == nil || result.NumberOfEntities == 0 {
		c.JSON(204, gin.H{"message": "No launches found"})
		return
	}

	c.JSON(200, result)
}

func RegisterRoutes(router *gin.Engine, controller *LaunchController) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	v1 := router.Group("/api/v1")
	launch := v1.Group("/launch")
	launch.GET("/search", controller.SearchLaunch)
}
