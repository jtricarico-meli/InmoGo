package main

import (
	config2 "InmoGo/src/api/config"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/services"
	"InmoGo/src/api/utils"
)

func main() {
	jwtMaker, _ := utils.NewJWTMaker("AAABBBBCCCDDDEEEFFFFAAABBBBCCCDDDEEEFFFFAAABBBBCCCDDDEEEFFFF")

	db, err := config2.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	//InitRepository
	propietarioRepository := repositories.NewPropietarioRepository(db)
	inmuebleRepository := repositories.NewInmuebleRepository(db)
	pagoRepository := repositories.NewPagoRepository(db)
	inquilinoRepository := repositories.NewInquilinoRepository(db)
	alquilerRepository := repositories.NewAlquilerRepository(db)

	//InitService
	propietarioService := services.NewPropietarioService(propietarioRepository, jwtMaker)
	inmuebleService := services.NewInmuebleService(inmuebleRepository)
	pagoService := services.NewPagoService(pagoRepository)
	inquilinoService := services.NewInquilinoService(inquilinoRepository)
	alquilerService := services.NewAlquilerService(alquilerRepository)

	server := config2.NewServer(propietarioService, inmuebleService, pagoService, inquilinoService, alquilerService)

	server.Run()
}
