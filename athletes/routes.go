package athlete

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AthleteController struct {
	athleteService *AthleteService
}

func (c *AthleteController) retrieveAthletes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"users": "yeaahh boi"}`))
}

func (c *AthleteController) createAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Created"}`))

}

func RegisterUserRoutes(athleteService *AthleteService, router *mux.Router) {
	controller := &AthleteController{athleteService: athleteService}
	router.HandleFunc("/athletes", controller.createAthlete).Methods(http.MethodPost)
	router.HandleFunc("/athletes", controller.retrieveAthletes).Methods(http.MethodGet)
}
