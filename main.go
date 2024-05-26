package main

import (
	"github.com/gin-contrib/cors"
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

	// Configuração do middleware CORS, estava ocorrendo erro nas requests.
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	router.POST("/ads", controllers.Create)
	router.GET("/ads", controllers.Get)
	router.GET("/ads/:id", controllers.GetById)
	router.PUT("/ads/:id", controllers.Update)
	router.DELETE("/ads/:id", controllers.Delete)

	router.Run("0.0.0.0:9000") // <- alterado pois com localhost não funciona
}
