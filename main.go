package main

import (
	"net/http"
	"technical_test/src/config"
	"technical_test/src/controller"

	_ "technical_test/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tittle technical_test
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api
func main() {
  config.ConnectionDB()
  router := gin.Default()
  router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  api:= router.Group("/api")
  {
  api.POST("/employees",controller.CreateEmployeeController)
  api.GET("/employees", controller.GetAllController)
  api.GET("/employees/:id", controller.GetByIdController)
  api.PUT("/employees/:id", controller.UpdateController)
  api.DELETE("/employees/:id", controller.DeletedController)
  }
  router.Run("0.0.0.0:8080")
}