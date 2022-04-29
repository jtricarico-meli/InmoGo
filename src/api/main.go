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

	//InitService
	propietarioService := services.NewPropietarioService(propietarioRepository)

	server := config2.NewServer(propietarioService)

	server.Run()
}
