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

//nolint
type PetsRouteCfg interface{}

var pets = pkg.Pets{
	NewPet(1, "Rex", nil),
	NewPet(2, "Spot", nil),
	NewPet(3, "Barry", nil),
}

func NewPet(id int64, name string, tag *string) pkg.Pet {
	return pkg.Pet{
		Id: id,
		NewPet: pkg.NewPet{
			Name: name,
			Tag:  tag,
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
			_, _ = w.Write(resp)
		}),
	}
}

func CreatePetRoute(cfg PetsRouteCfg) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodPost,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pet := NewPet(int64(len(pets)+1), "Rex", nil)
			pets = append(pets, pet)
			resp, err := json.Marshal(&pet)
			if err != nil {
				// TODO: handle error
				log.Warn("marshalling error")
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		}),
	}
}

func GetPetRoute(cfg PetsRouteCfg) *config.Route {
	return &config.Route{
		Path:   "/pets/{petId}",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			petID, err := strconv.ParseInt(vars["petId"], 10, 64)
			if err != nil {
				// TODO: add validation
				log.Warn("validation error")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			var pet pkg.Pet
			for _, p := range pets {
				if p.Id == petID {
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
			_, _ = w.Write(resp)
		}),
	}
}
