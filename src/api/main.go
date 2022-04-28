package main

import (
	config2 "InmoGo/src/api/config"
	"InmoGo/src/api/controllers"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/services"
)

func main() {
	config := config2.NewConfig()

	db, err := config2.ConnectDatabase(config)
	if err != nil {
		panic(err)
	}

	//InitRepository
	propietarioRepository := repositories.NewPropietarioRepository(db)

	//InitService
	propietarioService := services.NewPropietarioService(config, propietarioRepository)

	//InitController
	propietarioController := controllers.NewPropietarioController(propietarioService)

	server := config2.NewServer(config, propietarioController)

	server.Run()
}
