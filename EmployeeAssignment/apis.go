package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	employeeRouter := mux.NewRouter()
	employeeRouter.HandleFunc("/employees", getEmployees).Methods("GET")
	employeeRouter.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	employeeRouter.HandleFunc("/employee", addEmployee).Methods("POST")
	employeeRouter.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":3000", employeeRouter))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /employees")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Printf("GET: /employee/%v\n", params["id"])
	for _, emp := range employees {
		if params["id"] == emp.Id {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST: /employee")
	w.Header().Set("Content-Type", "application/json")
	var emp employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	employees = append(employees, emp)
	json.NewEncoder(w).Encode(emp)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Printf("DELETE: /employee/%v\n", params["id"])
	var delIndex int
	for idx, emp := range employees {
		if emp.Id == params["id"] {
			delIndex = idx
			break
		}
	}
	employees = append(employees[:delIndex], employees[(delIndex+1):]...)
	json.NewEncoder(w).Encode(employees)
}
