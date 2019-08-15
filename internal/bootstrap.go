package internal

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorldRoute() *Route {
	return &Route{
		Path:   "/hello",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		}),
	}
}

func simpleMiddlewareFn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func Bootstrap() {
	server := NewServer().
		WithMiddleware(
			simpleMiddlewareFn,
		).
		WithRoutes(
			HelloWorldRoute(),
		)

	server.Start("", 80)
}
