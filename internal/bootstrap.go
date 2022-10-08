package internal

import (
	"database/sql"
	"net/http"

	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/internal/persistence"
	"github.com/pseudo-su/golang-service-template/internal/pets"
	log "github.com/sirupsen/logrus"
)

func requestSetupMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type ApplicationContext interface {
	Env() string
	ServiceBasepath() string
	ServerPort() int
	SqlDB() *sql.DB
}

func Bootstrap(appCtx ApplicationContext) *config.Server {
	basepath := appCtx.ServiceBasepath()
	petsRepo := persistence.NewPetsRepository(appCtx.SqlDB())
	server := config.NewServer().
		WithMiddleware(
			requestSetupMiddleware,
		).
		MountRoutes(
			basepath,
			OpenAPISpecRoute(appCtx),
			SwaggerUIRoute(appCtx),
			SwaggerUIRedirectRoute(appCtx),
			pets.ListPetsRoute(appCtx, petsRepo),
			pets.CreatePetRoute(appCtx, petsRepo),
			pets.GetPetRoute(appCtx, petsRepo),
		)
	return server
}
