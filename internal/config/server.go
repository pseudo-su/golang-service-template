package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	http.Server
	router      *mux.Router
	basePath    string
	ShutdownReq chan bool
	done        chan bool
}

func NewServer() *Server {
	return &Server{
		Server:      http.Server{},
		router:      mux.NewRouter(),
		basePath:    "/",
		ShutdownReq: make(chan bool),
		done:        make(chan bool),
	}
}

func (s *Server) WithBasePath(basePath string) *Server {
	s.basePath = basePath
	return s
}

func (s *Server) WithMiddleware(mwf ...mux.MiddlewareFunc) *Server {
	s.router.Use(mwf...)
	return s
}

func (s *Server) WithRoutes(routes ...*Route) *Server {
	return s.MountRoutes(s.basePath, routes...)
}

func (s *Server) MountRoutes(path string, routes ...*Route) *Server {
	sub := s.router.PathPrefix(path).Subrouter()
	for _, route := range routes {
		sub.Handle(route.Path, route.Handler).Methods(route.Method)
	}
	return s
}

func DescribeRoutes(router *mux.Router) {
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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
}
func (s *Server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Infof("Shutdown request (signal: %v)", sig)
	case sig := <-s.ShutdownReq:
		log.Infof("Shutdown request (/shutdown %v)", sig)
	}

	log.Infof("Stoping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Infof("Shutdown request error: %v", err)
	}
}

//Start starts the server on the defined port
func (s *Server) Start(port int) {
	// Set server values
	s.Handler = s.router
	s.Addr = fmt.Sprintf(":%v", port)

	log.Infof("Starting server %v", s.Addr)
	DescribeRoutes(s.router)

	finish := make(chan error)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.WithError(err).Error("server closed")
		}
		finish <- err
	}()

	//wait shutdown
	s.WaitShutdown()

	<-finish
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
