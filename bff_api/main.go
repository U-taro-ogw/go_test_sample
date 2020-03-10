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
	response := map[string]string{"key":"value"}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	app := App{}
	app.Initialize()
	app.Run(":8082")
}

//func GetMainEngine() {
//	return http.ListenAndServe(":8082", router)
//}
//
//type Error struct {
//	Message string `json:"message"`
//}
//
//func GetApisInfo() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var errorObj Error
//		client := &http.Client{Timeout: time.Duration(30) * time.Second}
//		header := http.Header{}
//		header.Set("Content-Type", "application/json")
//		//header.Set("Authorization", c.Request.Header.Get("Authorization"))
//
//		fetchUrl := "http://flask_api_one:5000/api_info"
//
//		req, err := http.NewRequest("GET", fetchUrl, nil)
//		if err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusInternalServerError)
//			json.NewEncoder(w).Encode(errorObj)
//			return
//		}
//		req.Header = header
//
//		res, err := client.Do(req)
//		if err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusInternalServerError)
//			json.NewEncoder(w).Encode(errorObj)
//			return
//		}
//		defer res.Body.Close()
//
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusInternalServerError)
//			json.NewEncoder(w).Encode(errorObj)
//			return
//		}
//
//		fmt.Println("-------hogehogehoge")
//		fmt.Println(body)
//
//
//		errorObj.Message = "Not Found"
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(errorObj)
//	}
//}



//func getVolDetails(volName string, obj interface{}) error {
//	addr := os.Getenv("SOME_ADDR")
//	if addr == "" {
//		err := errors.New("SOME_ADDR environment variable not set")
//		fmt.Println(err)
//		return err
//	}
//	url := addr + "/path/to/somepage/" + volName
//	client := &http.Client{
//		Timeout: timeout,
//	}
//	resp, err := client.Get(url)
//	if resp != nil {
//		if resp.StatusCode == 500 {
//			fmt.Printf("VSM %s not found\n", volName)
//			return err
//		} else if resp.StatusCode == 503 {
//			fmt.Println("server not reachable")
//			return err
//		}
//	} else {
//		fmt.Println("server not reachable")
//		return err
//	}
//
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	defer resp.Body.Close()
//
//	return json.NewDecoder(resp.Body).Decode(obj)
//}

// GetVolAnnotations gets annotations of volume
//func GetVolAnnotations(volName string) (*Annotations, error) {
//	var volume Volume
//	var annotations Annotations
//	err := getVolDetails(volName, &volume)
//	if err != nil || volume.Metadata.Annotations == nil {
//		if volume.Status.Reason == "pending" {
//			fmt.Println("VSM status Unknown to server")
//		}
//		return nil, err
//	}
//	// Skipped some part,not required
//}

//func main() {
//	fmt.Println("Hello BFF")
//	handler := ApiHandler{
//		ApiOneUrl: "http://flask_api_one:5000",
//		ApiTwoUrl: "http://flask_api_two:6000/api_info",
//	}
//	GetMainEngine(handler).Run(":8082")
//}
//
//func GetMainEngine() *gin.Engine {
//	r := gin.Default()
//	v1 := r.Group("v1")
//	{
//		v1.GET("/apis", GetApisInfo())
//		//v1.POST("/apis/all_wait", handler.AllWait)
//		//v1.POST("/apis/quick_response", handler.QuickResponse)
//	}
//	return r
//}
//
//type ApiHandler struct {
//	ApiOneUrl string
//	ApiTwoUrl string
//}

//func (h *ApiHandler) Apis(c *gin.Context) {
	//var apiResponse []map[string]interface{}

	// port 5000にfetch
	// port 6000にfetch
	// merge
	// http.Geth.ApiOneUrl

	// goroutinで2つのapiを叩く
	// 叩いた結果を混ぜて（配列）にして返す

	//client := &http.Client{Timeout: time.Duration(30) * time.Second}
	//header := http.Header{}
	//header.Set("Content-Type", "application/json")
	//header.Set("Authorization", c.Request.Header.Get("Authorization"))
	//
	//fetchUrl := string(h.ApiOneUrl) + "/api_info"
	//
	//req, err := http.NewRequest("GET", fetchUrl, nil)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//req.Header = header
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{"hoge": string(body)})
	//return
//}

//func (h *ApiHandler) GetApiOneInfo() map[string]interface{} {
//	m := make(map[string]interface{})
//	return m
//}
//
//func (h *ApiHandler) GetApiTwoInfo() map[string]interface{} {
//	m := make(map[string]interface{})
//	return m
//}
