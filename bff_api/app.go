package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

	client := &http.Client{}
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", r.Header.Get("Authorization"))
	fetchURL := "http://flask_api_one:5000/api_info"

	req, _ := http.NewRequest("GET", fetchURL, nil)
	req.Header = header

	res, _ := client.Do(req)
	//if err != nil {
	//	respondWithJSON(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("======================================")
	fmt.Println(string(body))
	//response := map[string]string{"key":"value"}
	respondWithJSON(w, http.StatusOK, string(body))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}