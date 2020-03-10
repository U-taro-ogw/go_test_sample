package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) initializeRoutes() {
	p := app.Router.PathPrefix("/v1").Subrouter()

	p.HandleFunc("/apis", app.getApisInfo).Methods("GET")
}

func (app *App) getApisInfo(w http.ResponseWriter, r *http.Request) {
	// port 5000と6000にrequest
	// responseを配列に詰める
	response := map[string]string{"key":"value"}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}