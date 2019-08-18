package internal

import (
	"net/http"

	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/internal/pets"
	log "github.com/sirupsen/logrus"
)

func requestSetupMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type ApplicationConfig interface {
	Env() string
	ServiceBasepath() string
	ServerPort() int
}

func Bootstrap(cfg ApplicationConfig) *config.Server {
	basepath := cfg.ServiceBasepath()
	server := config.NewServer().
		WithMiddleware(
			requestSetupMiddleware,
		).
		MountRoutes(
			basepath,
			OpenAPISpecRoute(cfg),
			SwaggerUIRoute(cfg),
			SwaggerUIRedirectRoute(cfg),
			pets.ListPetsRoute(cfg),
			pets.CreatePetRoute(cfg),
			pets.GetPetRoute(cfg),
		)
	return server
}
