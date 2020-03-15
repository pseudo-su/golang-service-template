package pets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pseudo-su/golang-service-template/internal/config"
	log "github.com/sirupsen/logrus"
)

//nolint:go-lint
type PetsRouteCfg interface{}

var pets Pets

func init() {
	pets = Pets{
		Pet{
			Id:   1,
			Name: "Rex",
			Tag:  nil,
		},
		Pet{
			Id:   2,
			Name: "Spot",
			Tag:  nil,
		},
		Pet{
			Id:   3,
			Name: "Barry",
			Tag:  nil,
		},
	}
}

func ListPetsRoute(cfg PetsRouteCfg) *config.Route {
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

func CreatePetRoute(cfg PetsRouteCfg) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodPost,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pet := Pet{
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

func GetPetRoute(cfg PetsRouteCfg) *config.Route {
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
			var pet Pet
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
