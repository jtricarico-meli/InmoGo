package controllers

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/services"
	"InmoGo/src/api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

var (
	inmueble = services.Inmueble{}
)

func SaveInmueble(c *gin.Context) {
	request := utils.GetRequest(c, models.Inmueble{})
	if err := mapstructure.Decode(request.Body, &inmueble); err != nil {
		fmt.Println(err)
	}
	if err := inmueble.Save(inmueble); err != nil {
		fmt.Println(err)
	}
}

func FindInmueble(c *gin.Context) {
	request := utils.GetRequest(c, models.Inmueble{})
	if err := inmueble.Save(request.Body.(services.Inmueble)); err != nil {
		fmt.Println(err)
	}
}

func FindAllInmuebles(c *gin.Context) {

}

func UpdateInmueble(c *gin.Context) {

}
