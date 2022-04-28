package config

import (
	"InmoGo/src/api/controllers"
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct {
	config                *Config
	propietarioController *controllers.PropietarioController
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/propietario", s.handlePost)
	mux.HandleFunc("/propietario", s.handleGet)

	return mux
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
		res = s.propietarioController.Save(r.Body)
	}

	bytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	if strings.Contains(r.URL.Path, "/propietario") {
		res = s.propietarioController.Get(1)
	}
	bytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, propController *controllers.PropietarioController) *Server {
	return &Server{
		config:                config,
		propietarioController: propController,
	}
}
