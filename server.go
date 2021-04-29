package main

import (
	"errors"
	athlete "intervals/athletes"
	"intervals/infra"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "There is nothing here :)"}`))
}

func main() {
	port, portErr := getPort()
	athletesService, err := createAthletesService()
	if err != nil || portErr != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", notFound)
	athlete.RegisterUserRoutes(athletesService, router)
	log.Fatal(http.ListenAndServe(port, router))
}

func createAthletesService() (*athlete.AthleteService, error) {
	db, err := infra.CreatePool()

	if err != nil {
		return nil, err
	}

	repository := athlete.CreateRepository(db)
	return athlete.CreateService(repository), nil
}

func getPort() (string, error) {
	appPort := os.Getenv("PORT")
	if appPort != "" {
		return ":" + appPort, nil
	}
	return "", errors.New("port was not configured")
}
