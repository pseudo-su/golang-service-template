package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	router   *mux.Router
	basePath string
}

func NewServer() *Server {
	return &Server{
		router:   mux.NewRouter(),
		basePath: "/",
	}
}

func (server *Server) WithBasePath(basePath string) *Server {
	server.basePath = basePath
	return server
}

func (server *Server) WithMiddleware(mwf ...mux.MiddlewareFunc) *Server {
	server.router.Use(mwf...)
	return server
}

func (server *Server) WithRoutes(routes ...*Route) *Server {
	return server.MountRoutes(server.basePath, routes...)
}

func (server *Server) MountRoutes(path string, routes ...*Route) *Server {
	sub := server.router.PathPrefix(path).Subrouter()
	for _, route := range routes {
		sub.Handle(route.Path, route.Handler).Methods(route.Method)
	}
	return server
}

//Start starts the server on the defined port
func (server *Server) Start(port int) {
	_ = server.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		path, err := route.URLPath()
		if err != nil {
			return nil
		}
		for _, method := range methods {
			log.WithFields(log.Fields{
				"method": method,
				"path":   path,
			}).Info("path registered")
		}
		return nil
	})

	panic(
		http.ListenAndServe(
			fmt.Sprintf(":%v", port),
			handlers.RecoveryHandler()(server.router),
		),
	)
}

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

// .Handle(“/path”, a(b(c(http.HandlerFunc(yourHandler))
func (route *Route) WithMiddleware(m ...mux.MiddlewareFunc) *Route {
	hWithMiddleware := route.Handler
	for i := len(m) - 1; i >= 0; i-- {
		hWithMiddleware = m[i](hWithMiddleware)
	}

	route.Handler = hWithMiddleware
	return route
}
