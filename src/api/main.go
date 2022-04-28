package main

import (
	config2 "InmoGo/src/api/config"
	"InmoGo/src/api/controllers"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/services"
)

func main() {
	db, err := config2.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	//InitRepository
	propietarioRepository := repositories.NewPropietarioRepository(db)

	//InitService
	propietarioService := services.NewPropietarioService(propietarioRepository)

	//InitController
	propietarioController := controllers.NewPropietarioController(propietarioService)

	server := config2.NewServer(propietarioController)

	server.Run()
}
