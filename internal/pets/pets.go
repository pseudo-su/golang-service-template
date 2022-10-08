//nolint:golint,godox,gochecknoinits
package pets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/internal/persistence"
	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/sirupsen/logrus"
)

func newPetFromDB(dbPet *persistence.Pet) *pkg.Pet {
	return &pkg.Pet{
		Id:   dbPet.APIID,
		Name: dbPet.Name,
		Tag:  dbPet.Tag,
	}
}

func newPetListFromDB(dbList []persistence.Pet) []pkg.Pet {
	result := []pkg.Pet{}

	for _, val := range dbList {
		result = append(result, *newPetFromDB(&val))
	}

	return result
}

type PetsRouteContext interface{}

func ListPetsRoute(routeCtx PetsRouteContext, petsRepo persistence.PetsRepositoryInterface) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pets, err := petsRepo.ListPets(r.Context(), &persistence.PaginationValues{
				Limit:  10,
				Offset: 0,
			})
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := json.Marshal(newPetListFromDB(pets))
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		}),
	}
}

func CreatePetRoute(routeCtx PetsRouteContext, petsRepo persistence.PetsRepositoryInterface) *config.Route {
	return &config.Route{
		Path:   "/pets",
		Method: http.MethodPost,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pet, err := petsRepo.CreatePet(r.Context(), &persistence.PetValues{
				Name: "Rex",
				Tag:  nil,
			})
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := json.Marshal(newPetFromDB(pet))
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write(resp)
		}),
	}
}

func GetPetRoute(routeCtx PetsRouteContext, petsRepo persistence.PetsRepositoryInterface) *config.Route {
	return &config.Route{
		Path:   "/pets/{petId}",
		Method: http.MethodGet,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			petID, err := strconv.ParseInt(vars["petId"], 10, 64)
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			pet, err := petsRepo.GetPetByAPIID(r.Context(), petID)
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := json.Marshal(newPetFromDB(pet))
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		}),
	}
}
