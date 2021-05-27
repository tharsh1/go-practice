package main

import (
	"fmt"
)

func main() {
	cntnu := 1
	var choice int
	var teleNum int
	var name string
	var address string
	for cntnu == 1 {
		fmt.Println("1. Add, 2. Delete, 3. Search, 4. Update, 5. Display, 6. Find Duplicates, 7. Create CSV\nEnter your coice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter telephone number: ")
			fmt.Scanln(&teleNum)
			fmt.Println("Enter assiciated name: ")
			fmt.Scanln(&name)
			fmt.Println("Enter assiciated address: ")
			fmt.Scanln(&address)
			addRecord(teleNum, name, address)
		case 2:
			fmt.Println("Enter telephone number which is to be deleted: ")
			fmt.Scanln(&teleNum)
			deleteRecord(teleNum)
		case 3:
			fmt.Println("Enter telephone number which is to be searched: ")
			fmt.Scanln(&teleNum)
			searchRecord(teleNum)
		case 4:
			fmt.Println("Enter telephone number whose details is to be updated: ")
			fmt.Scanln(&teleNum)
			updateRecord(teleNum)
		case 5:
			displayRecord()
		case 6:
			findDuplicates()
		case 7:
			createCSV()
		default:
			fmt.Println("Invalid input")
		}
		fmt.Println("Do you want to continue(Y=1/N=0): ")
		fmt.Scanln(&cntnu)
	}
}
