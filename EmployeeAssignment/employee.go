package main

import "fmt"

type employee struct {
	Id         string  `json:"id`
	Name       string  `json:"name`
	Age        int     `json:"age`
	Salary     float64 `json:"salary`
	Department string  `json:"department`
}

var employees []employee

func constructInitialEmployees() {
	employees = []employee{{
		Id:         "1001",
		Name:       "Harsh",
		Age:        20,
		Salary:     10000.5,
		Department: "Production",
	}, {
		Id:         "1002",
		Name:       "Rohan",
		Age:        22,
		Salary:     450000,
		Department: "IT",
	}, {
		Id:         "1003",
		Name:       "Aarohi",
		Age:        25,
		Salary:     103200,
		Department: "Sales",
	}, {
		Id:         "1004",
		Name:       "Girish",
		Age:        26,
		Salary:     120000,
		Department: "Testing",
	}, {
		Id:         "1005",
		Name:       "Kanchan",
		Age:        23,
		Salary:     250000,
		Department: "Marketing",
	},
	}
}

func printEmployees() {
	for _, emp := range employees {
		fmt.Println("Id:", emp.Id, ", Name:", emp.Name, ", Age:", emp.Age, ", Salary:", emp.Salary, ", Department:", emp.Department)
	}
}

func fetchMaxSalary() float64 {
	maxEmpSal := 0.0
	for _, emp := range employees {
		if emp.Salary > maxEmpSal {
			maxEmpSal = emp.Salary
		}
	}
	return maxEmpSal
}
