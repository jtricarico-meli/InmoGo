package app

import (
	"InmoGo/src/api/config"
	"InmoGo/src/api/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func StartApp() {
	//Config router
	configureRouter()
	mapUrlToPing()

	//Init DB and Close when ShutDown
	config.InitDB()
	defer config.CloseDB()

	//RUN APP
	Router.Run(":8080")

}

func configureRouter() {
	// Default meli router - includes newrelic, datadog, attributes filter, jsonp and pprof middlewares
	Router = gin.Default()
}

func mapUrlToPing() {
	// Add health check

	Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//Router.Use(HandIp())

	//Inmuebles
	Router.GET("/inmuebles", controllers.FindAllInmuebles)
	Router.GET("/inmuebles/:ID", controllers.FindInmueble)
	Router.PUT("/inmuebles/:ID", controllers.UpdateInmueble)
	Router.POST("/inmuebles", controllers.SaveInmueble)

	fmt.Println("Starting Server!")
}

func HandIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
