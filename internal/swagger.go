package internal

import (
	"fmt"
	"net/http"

	"github.com/pseudo-su/golang-service-template/internal/config"
	log "github.com/sirupsen/logrus"
)

type SwaggerRouteCfg interface {
	Env() string
	ServiceBasepath() string
}

var (
	SpecPath     = "openapi.json"
	UIPath       = "swagger.html"
	RedirectPath = "swagger-oauth2-redirect.html"
)

var EnvWhitelist = []string{
	"local",
	"dev",
	"staging",
}

func swaggerWhitelisted(env string) bool {
	for _, a := range EnvWhitelist {
		if a == env {
			return true
		}
	}
	return false
}

func GetSwaggerFn() GetBytesFn {
	return func(specURL, redirectURL string) ([]byte, error) {
		openapi, err := GetSwagger()
		if err != nil {
			return nil, err
		}
		b, err := openapi.MarshalJSON()
		if err != nil {
			return nil, err
		}
		return b, nil
	}
}

func OpenAPISpecRoute(cfg SwaggerRouteCfg) (route *config.Route) {
	routePath := fmt.Sprintf("/%s", SpecPath)

	specPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), SpecPath)
	redirectURLPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), RedirectPath)

	route = &config.Route{
		Path:    routePath,
		Method:  http.MethodGet,
		Handler: createServeBytesHandler(cfg.Env(), specPath, redirectURLPath, "application/json", GetSwaggerFn()),
	}
	return route
}

func SwaggerUIRoute(cfg SwaggerRouteCfg) (route *config.Route) {
	routePath := fmt.Sprintf("/%s", UIPath)

	specPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), SpecPath)
	redirectURLPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), RedirectPath)

	route = &config.Route{
		Path:    routePath,
		Method:  http.MethodGet,
		Handler: createServeBytesHandler(cfg.Env(), specPath, redirectURLPath, "text/html", GetSwaggerUIPage),
	}
	return route
}

func SwaggerUIRedirectRoute(cfg SwaggerRouteCfg) (route *config.Route) {
	routePath := fmt.Sprintf("/%s", RedirectPath)
	specPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), SpecPath)
	redirectURLPath := fmt.Sprintf("%s/%s", cfg.ServiceBasepath(), RedirectPath)

	route = &config.Route{
		Path:    routePath,
		Method:  http.MethodGet,
		Handler: createServeBytesHandler(cfg.Env(), specPath, redirectURLPath, "text/html", GetSwaggerUIRedirect),
	}
	return route
}

type GetBytesFn func(string, string) ([]byte, error)

func createServeBytesHandler(env, specPath, redirectURLPath, contentType string, fn GetBytesFn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !swaggerWhitelisted(env) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		swaggerUI, err := fn(specPath, redirectURLPath)
		if err != nil {
			log.Warn(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", contentType)
		_, err = w.Write(swaggerUI)

		if err != nil {
			log.Warn(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	})
}
