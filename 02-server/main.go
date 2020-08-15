package main

import (
	"log"
    "os"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Message struct {
	Message string `json:message`
}

type Data struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Comments []Message `json:"comments"`
}
var data = Data{}

func SaveLog() {
	file, err := os.OpenFile("catatan.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	e, err := json.Marshal(data)
	if err != nil {
        fmt.Println(err)
        return
	}
	
	log.SetOutput(file)
	log.Print(string(e))
}

func PostMessage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&data)
	SaveLog()
	json.NewEncoder(w).Encode(&data)
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/", PostMessage).Methods("POST")

	http.ListenAndServe(":8000", router)
}