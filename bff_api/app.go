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

	p.HandleFunc("/apis", app.apisInfo).Methods("GET")
}

func (app *App) apisInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("<<< apisInfo ")
	fetchURL := "http://flask_api_one:5000/api_info"
	code, body, err := GetResponse(fetchURL, r.Header.Get("Authorization"))
	if err != nil {
		respondWithJSON(w, code, err)
		return
	}
	respondWithJSON(w, http.StatusOK, string(body))
}

func GetResponse(url, jwtToken string) (int, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", jwtToken)
	req.Header = header

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil && code != http.StatusOK {
		return code, nil, fmt.Errorf("Server status error: %v", http.StatusText(code))
	}

	return code, body, nil
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
