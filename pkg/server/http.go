package server

import (
	"net/http"
	"strconv"
	"time"
	"fmt"

	log "github.com/AKovalevich/trafficcop/pkg/log/logrus"
)

func (server *Server) configureMainHttpServer() {
	s := http.NewServeMux()

	for _, entrypoint := range server.mainConfiguration.EntryPoints {
		log.Do.Info("Add entrypoint: " + entrypoint.String())

		entrypoint.Init()
		routesList := entrypoint.RoutesList()
		for _, route := range routesList {
			s.HandleFunc(route.Path, route.Handler)
		}
	}
	server.mainHttpServer = &http.Server{
		Handler:      s,
		Addr: server.mainConfiguration.ServerHost + ":" + strconv.Itoa(server.mainConfiguration.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (server *Server) configureWebUiHttpServer() {
	s := http.NewServeMux()
	s.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!\n")
	})
	server.webUiHttpServer = &http.Server{
		Handler:      s,
		Addr: server.mainConfiguration.WebUIHost + ":" + strconv.Itoa(server.mainConfiguration.WebUIPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (server *Server) runMainServer() {
	log.Do.Infof("Starting main trafficcop server on: %s", server.mainHttpServer.Addr)
	log.Do.Warn(server.mainHttpServer.ListenAndServe())
}

func (server *Server) runWebUiServer() {
	log.Do.Infof("Starting Web UI on: %s", server.webUiHttpServer.Addr)
	log.Do.Warn(server.webUiHttpServer.ListenAndServe())
}
