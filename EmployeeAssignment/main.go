package main

import "fmt"

type employee struct {
	empid  int
	name   string
	age    int
	salary float64
	dept   string
}

var employees []employee

func main() {
	employees = []employee{{
		empid:  1001,
		name:   "Harsh",
		age:    20,
		salary: 10000.5,
		dept:   "Production",
	}, {
		empid:  1002,
		name:   "Rohan",
		age:    22,
		salary: 450000,
		dept:   "IT",
	}, {
		empid:  1003,
		name:   "Aarohi",
		age:    25,
		salary: 103200,
		dept:   "Sales",
	}, {
		empid:  1004,
		name:   "Girish",
		age:    26,
		salary: 120000,
		dept:   "Testing",
	}, {
		empid:  1005,
		name:   "Kanchan",
		age:    23,
		salary: 250000,
		dept:   "Marketing",
	},
	}

	printEmployees()

	fmt.Println("Maximum salary is:", fetchMaxSalary())
}

func printEmployees() {
	for _, emp := range employees {
		fmt.Println("ID:", emp.empid, ", Name:", emp.name, ", Age:", emp.age, ", Salary:", emp.salary, ", Department:", emp.dept)
	}
}

func fetchMaxSalary() float64 {
	maxEmpSal := 0.0
	for _, emp := range employees {
		if emp.salary > maxEmpSal {
			maxEmpSal = emp.salary
		}
	}
	return maxEmpSal
}
