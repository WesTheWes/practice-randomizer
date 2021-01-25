package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func randomScale(w http.ResponseWriter, r *http.Request) {
	scaleResponse := generateRandomScale()
	ret, _ := json.Marshal(scaleResponse)
	fmt.Fprintf(w, "%s", ret)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/random_scale", randomScale).Methods("GET")
	router.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":3000", router)
}
