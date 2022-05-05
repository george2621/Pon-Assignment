package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gorilla/mux"
)

type Message struct {
	Msg string `json:"message"`
}

var called uint64
var msg = Message{Msg: "Hello world"}

func showMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(msg.Msg); err != nil {
		log.Print(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(called); err != nil {
		log.Print(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
	
	atomic.AddUint64(&called, 1)
	log.Print("Show message")

}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Print(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(msg.Msg); err != nil {
		log.Print(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
	log.Print("Message has been updated")

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", showMessage).Methods("GET")
	r.HandleFunc("/", updateMessage).Methods("PUT")

	f, _ := os.Create("golang-server.log")
	defer f.Close()
	log.SetOutput(f)

	err := http.ListenAndServeTLS(":"+port, "localhost.crt", "localhost.key", r)
	if err != nil {
		log.Print(err)
	}else {
		log.Print("Server Started on port: "+port)
	}
}
