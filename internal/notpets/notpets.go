package notpets

import (
	"encoding/json"
	"net/http"

	"github.com/pseudo-su/golang-service-template/internal/config"
)

//nolint:go-lint
type NotPetsRouteCfg interface{}

var notpets NotPets

func init() {
	notpets = NotPets{
		NotPet{
			Id:   1,
			Name: "Rex",
			Tag:  nil,
		},
		NotPet{
			Id:   2,
			Name: "Spot",
			Tag:  nil,
		},
		NotPet{
			Id:   3,
			Name: "Barry",
			Tag:  nil,
		},
	}
}

func ListNotPetsRoute(cfg NotPetsRouteCfg) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(&notpets)

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}),
	}
}
