package main

import (
	athlete "intervals/athletes"
	"intervals/infra"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "There is nothing here :)"}`))
}

func main() {

	athletesService, err := createAthletesService()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", notFound)
	athlete.RegisterUserRoutes(athletesService, router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createAthletesService() (*athlete.AthleteService, error) {
	db, err := infra.CreatePool()

	if err != nil {
		return nil, err
	}

	repository := athlete.CreateRepository(db)
	return athlete.CreateService(repository), nil
}
