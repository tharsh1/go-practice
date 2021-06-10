package main

import "fmt"

func main() {
	constructInitialEmployees()
	printEmployees()
	fmt.Println("Maximum salary is:", fetchMaxSalary())
	handleRequests()
}
