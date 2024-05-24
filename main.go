package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaodrts/Drts.Ads.GO.API/config"
	"github.com/joaodrts/Drts.Ads.GO.API/controllers"
)

func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/ads", controllers.Create)
	router.GET("/ads", controllers.Get)
	router.GET("/ads/:id", controllers.GetById)
	router.PUT("/ads/:id", controllers.Update)
	router.DELETE("/ads/:id", controllers.Delete)

	router.Run("localhost:9000")
}
