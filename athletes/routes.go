package athlete

import (
	"encoding/json"
	"log"
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
	var athlete Athlete
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&athlete)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "the provided athlete is incorrect"}`))
	}

	createdAthlete, err := c.athleteService.createAthlete(athlete)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error creating the athlete"}`))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAthlete)

}

func RegisterUserRoutes(athleteService *AthleteService, router *mux.Router) {
	controller := &AthleteController{athleteService: athleteService}
	router.HandleFunc("/athletes", controller.createAthlete).Methods(http.MethodPost)
	router.HandleFunc("/athletes", controller.retrieveAthletes).Methods(http.MethodGet)
}
