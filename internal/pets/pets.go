package pets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/pkg"
	log "github.com/sirupsen/logrus"
)

//nolint:go-lint
type PetsRouteContext interface{}

var pets pkg.Pets

func init() {
	pets = pkg.Pets{
		pkg.Pet{
			Id:   1,
			Name: "Rex",
			Tag:  nil,
		},
		pkg.Pet{
			Id:   2,
			Name: "Spot",
			Tag:  nil,
		},
		pkg.Pet{
			Id:   3,
			Name: "Barry",
			Tag:  nil,
		},
	}
}

func ListPetsRoute(routeCtx PetsRouteContext) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(&pets)

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}),
	}
}

func CreatePetRoute(routeCtx PetsRouteContext) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodPost,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pet := pkg.Pet{
				Id:   int64(len(pets) + 1),
				Name: "Rex",
				Tag:  nil,
			}
			pets = append(pets, pet)
			resp, err := json.Marshal(&pet)
			if err != nil {
				// TODO: handle error
				log.Warn("marshalling error")
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}),
	}
}

func GetPetRoute(routeCtx PetsRouteContext) *config.Route {
	return &config.Route{
		Path:   "/pets/{petId}",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			petId, err := strconv.ParseInt(vars["petId"], 10, 64)
			if err != nil {
				// TODO: add validation
				log.Warn("validation error")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			var pet pkg.Pet
			for _, p := range pets {
				if p.Id == petId {
					pet = p
				}
			}
			resp, err := json.Marshal(&pet)
			if err != nil {
				// TODO: handle error
				log.Warn("marshalling error")
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}),
	}
}
