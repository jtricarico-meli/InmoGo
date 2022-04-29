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
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/propietario/", s.handlerMethod)

	mux.HandleFunc("/", s.handleGet)

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
	if strings.Contains(r.URL.Path, "/propietario") {
		all, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		var prop = models.Propietario{}
		err = json.Unmarshal(all, &prop)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(prop)
		res = s.propietario.Save(&prop)
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
	if strings.Contains(r.URL.Path, "/propietario/") {
		intID, _ := strconv.Atoi(id)
		res = s.propietario.Get(intID)
	} else {
		res = "pong"
	}
	bytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(propietario *services.PropietarioService) *Server {
	return &Server{
		propietario: propietario,
	}
}
