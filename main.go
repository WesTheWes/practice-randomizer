package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func randomScale(w http.ResponseWriter, r *http.Request) {
	scaleResponse := generateRandomScale()
	ret, _ := json.Marshal(scaleResponse)
	fmt.Fprintf(w, "%s", ret)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/random_scale", randomScale).Methods("GET")
	router.Handle("/", http.FileServer(http.Dir(".")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
