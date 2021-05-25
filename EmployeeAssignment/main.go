package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
	dept   string
}

func main() {
	employees := map[int]employee{
		1001: {
			name:   "Harsh",
			age:    20,
			salary: 10000.5,
			dept:   "Production",
		},
		1002: {
			name:   "Rohan",
			age:    22,
			salary: 450000,
			dept:   "IT",
		},
		1003: {
			name:   "Aarohi",
			age:    25,
			salary: 103200,
			dept:   "Sales",
		},
		1004: {
			name:   "Girish",
			age:    26,
			salary: 120000,
			dept:   "Testing",
		},
		1005: {
			name:   "Kanchan",
			age:    23,
			salary: 250000,
			dept:   "Marketing",
		},
	}

	employees[1006] = employee{
		name:   "Akshata",
		age:    21,
		salary: 234501,
		dept:   "Medical",
	}

	for eid, emp := range employees {
		fmt.Println("ID:", eid, ", Name:", emp.name, ", Age:", emp.age, ", Salary:", emp.salary, ", Department:", emp.dept)
	}

	maxEmp := employees[1001]
	for _, emp := range employees {
		if emp.salary > maxEmp.salary {
			maxEmp = emp
		}
	}
	fmt.Println(maxEmp.name, "has the maximum salary of", maxEmp.salary)
}
