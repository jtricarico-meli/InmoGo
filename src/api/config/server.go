package config

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/services"
	"InmoGo/src/api/utils"
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

var userMail string

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handlerMethod)

	return mux
}

func shouldBeAuthenticate(r *http.Request) bool {
	return !strings.Contains(r.URL.Path, "login") &&
		!(strings.Contains(r.URL.Path, "propietario") && r.Method == "POST") &&
		!(r.Method == "OPTIONS")
}

func (s *Server) handlerMethod(w http.ResponseWriter, r *http.Request) {
	if shouldBeAuthenticate(r) {
		mail, err := utils.Authenticate(r, s.propietario.JWT)
		if err != nil {
			panic(err)
		}
		userMail = mail
	}
	switch r.Method {
	case "POST":
		s.handlePost(w, r)
	case "GET":
		s.handleGet(w, r)
	case "OPTIONS":
		s.handleOptions(w, r)
	default:
		s.setResponse(utils.InmoError{
			Code:    405,
			Message: "method not supported",
		},
			http.StatusMethodNotAllowed,
			w,
		)
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
		s.setResponse(err, http.StatusInternalServerError, w)
	}
	if strings.Contains(r.URL.Path, "/login") {

		var prop = models.Propietario{}
		err = json.Unmarshal(all, &prop)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res, err = s.propietario.Login(prop.Mail, prop.Password)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	}

	if strings.Contains(r.URL.Path, "/propietario") {

		var prop = models.Propietario{}
		err = json.Unmarshal(all, &prop)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res, err = s.propietario.Save(&prop)
	}

	if strings.Contains(r.URL.Path, "/inmueble") {

		var inmueble = models.Inmueble{}
		err = json.Unmarshal(all, &inmueble)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res = s.inmueble.Save(&inmueble)
	}

	if strings.Contains(r.URL.Path, "/pago") {

		var pago = models.Pago{}
		err = json.Unmarshal(all, &pago)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res = s.pago.Save(&pago)
	}

	if strings.Contains(r.URL.Path, "/alquiler") {

		var alquiler = models.Alquiler{}
		err = json.Unmarshal(all, &alquiler)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res = s.alquiler.Save(&alquiler)
	}

	if strings.Contains(r.URL.Path, "/inquilino") {

		var inquilino = models.Inquilino{}
		err = json.Unmarshal(all, &inquilino)
		if err != nil {
			fmt.Println(err)
			s.setResponse(err, http.StatusBadRequest, w)
		}
		res = s.inquilino.Save(&inquilino)
	}

	if err != nil {
		s.setErrorResponse(err, w)
	}

	if res != nil {
		s.setResponse(res, http.StatusCreated, w)
	}
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var err error
	i := strings.LastIndex(r.URL.Path, "/")
	id := r.URL.Path[i+1:]

	//PROPIETARIO
	if strings.Contains(r.URL.Path, "/propietario/") {
		res = s.propietario.Get(userMail)
	}

	//INMUEBLE
	if strings.Contains(r.URL.Path, "/inmueble/all") {
		intID, _ := strconv.Atoi(id)
		res, err = s.inmueble.GetAll(intID)
	} else if strings.Contains(r.URL.Path, "/inmueble/") {
		intID, _ := strconv.Atoi(id)
		res, err = s.inmueble.Get(intID)
	}

	//PAGO
	if strings.Contains(r.URL.Path, "/pago/all") {
		intID, _ := strconv.Atoi(id)
		res = s.pago.GetAll(intID)
	} else if strings.Contains(r.URL.Path, "/pago/") {
		intID, _ := strconv.Atoi(id)
		res, err = s.pago.Get(intID)
	}

	//ALQUILER
	if strings.Contains(r.URL.Path, "/alquiler/all") {
		intID, _ := strconv.Atoi(id)
		res, err = s.alquiler.GetAllByInmueble(intID)
	} else if strings.Contains(r.URL.Path, "/alquiler/") {
		intID, _ := strconv.Atoi(id)
		res, err = s.alquiler.Get(intID)
	}

	//INQUILINO
	if strings.Contains(r.URL.Path, "/inquilino/") {
		intID, _ := strconv.Atoi(id)
		res, err = s.inquilino.Get(intID)
	}

	if err != nil {
		s.setErrorResponse(err, w)
	} else {
		if res != nil {
			s.setResponse(res, http.StatusOK, w)
		}
	}
}

func (s *Server) handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) setErrorResponse(err error, w http.ResponseWriter) {
	var inmoError *utils.InmoError
	errJson := json.Unmarshal([]byte(err.Error()), &inmoError)
	if errJson != nil {
		s.setResponse(err, http.StatusInternalServerError, w)
		return
	}
	s.setResponse(err, inmoError.Code, w)
	return
}

func (s *Server) setResponse(v any, status int, w http.ResponseWriter) {
	bytes, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	w.WriteHeader(status)
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
