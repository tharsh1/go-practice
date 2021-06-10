package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/telephoneRecords", getTelephoneRecords)
	log.Fatalln(http.ListenAndServe(":3000", nil))
}

func getTelephoneRecords(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(telephoneDir)
}
