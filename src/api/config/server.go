package config

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	propietario *services.PropietarioService
	inmueble    *services.InmuebleService
	alquiler    *services.AlquilerService
	inquilino   *services.InquilinoService
	pago        *services.PagoService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handlerMethod)

	return mux
}

func (s *Server) handlerMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.handlePost(w, r)
	case "GET":
		s.handleGet(w, r)
	default:
		panic(fmt.Sprintf("Not Found Handler for method: %s", r.Method))
	}
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(r.URL.Path, "/propietario") {

		var prop = models.Propietario{}
		err = json.Unmarshal(all, &prop)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(prop)
		res = s.propietario.Save(&prop)
	}

	if strings.Contains(r.URL.Path, "/inmueble") {

		var inmueble = models.Inmueble{}
		err = json.Unmarshal(all, &inmueble)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(inmueble)
		res = s.inmueble.Save(&inmueble)
	}

	if strings.Contains(r.URL.Path, "/pago") {

		var pago = models.Pago{}
		err = json.Unmarshal(all, &pago)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pago)
		res = s.pago.Save(&pago)
	}

	if strings.Contains(r.URL.Path, "/alquiler") {

		var alquiler = models.Alquiler{}
		err = json.Unmarshal(all, &alquiler)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(alquiler)
		res = s.alquiler.Save(&alquiler)
	}

	if strings.Contains(r.URL.Path, "/inquilino") {

		var inquilino = models.Inquilino{}
		err = json.Unmarshal(all, &inquilino)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(inquilino)
		res = s.inquilino.Save(&inquilino)
	}

	bytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	i := strings.LastIndex(r.URL.Path, "/")
	id := r.URL.Path[i+1:]

	//PROPIETARIO
	if strings.Contains(r.URL.Path, "/propietario/") {
		intID, _ := strconv.Atoi(id)
		res = s.propietario.Get(intID)
	}

	//INMUEBLE
	if strings.Contains(r.URL.Path, "/inmueble/all") {
		intID, _ := strconv.Atoi(id)
		res = s.inmueble.GetAll(intID)
	}
	if strings.Contains(r.URL.Path, "/inmueble/") {
		intID, _ := strconv.Atoi(id)
		res = s.inmueble.Get(intID)
	}

	//PAGO
	if strings.Contains(r.URL.Path, "/pago/all") {
		intID, _ := strconv.Atoi(id)
		res = s.pago.GetAll(intID)
	}
	if strings.Contains(r.URL.Path, "/pago/") {
		intID, _ := strconv.Atoi(id)
		res = s.pago.Get(intID)
	}

	//ALQUILER
	if strings.Contains(r.URL.Path, "/alquiler/all") {
		intID, _ := strconv.Atoi(id)
		res = s.alquiler.GetAllByInmueble(intID)
	}
	if strings.Contains(r.URL.Path, "/alquiler/") {
		intID, _ := strconv.Atoi(id)
		res = s.alquiler.Get(intID)
	}

	//INQUILINO
	if strings.Contains(r.URL.Path, "/inquilino/") {
		intID, _ := strconv.Atoi(id)
		res = s.inquilino.Get(intID)
	} else {
		res = "pong"
	}
	bytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(
	propietario *services.PropietarioService,
	inmueble *services.InmuebleService,
	pago *services.PagoService,
	inquilino *services.InquilinoService,
	alquiler *services.AlquilerService) *Server {
	return &Server{
		propietario: propietario,
		inmueble:    inmueble,
		pago:        pago,
		inquilino:   inquilino,
		alquiler:    alquiler,
	}
}
