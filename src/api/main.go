package main

import (
	config2 "InmoGo/src/api/config"
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
	inmuebleRepository := repositories.NewInmuebleRepository(db)

	//InitService
	propietarioService := services.NewPropietarioService(propietarioRepository)
	inmuebleService := services.NewInmuebleService(inmuebleRepository)

	server := config2.NewServer(propietarioService, inmuebleService)

	server.Run()
}
