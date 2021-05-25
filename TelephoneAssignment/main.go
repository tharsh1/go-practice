package main

import (
	"fmt"
)

type telephoneRecord struct {
	name    string
	address string
}

func main() {
	telephoneDir := map[int]telephoneRecord{}
	cntnu := 1
	var choice int
	var teleNum int
	var name string
	var address string
	for cntnu == 1 {
		fmt.Println("1. Add, 2. Delete, 3. Search, 4. Update, 5. Display\nEnter your coice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter telephone number: ")
			fmt.Scanln(&teleNum)
			fmt.Println("Enter assiciated name: ")
			fmt.Scanln(&name)
			fmt.Println("Enter assiciated address: ")
			fmt.Scanln(&address)
			telephoneDir[teleNum] = telephoneRecord{
				name:    name,
				address: address,
			}
			fmt.Println("Telephone record added succeddfully")
		case 2:
			fmt.Println("Enter telephone number which is to be deleted: ")
			fmt.Scanln(&teleNum)
			if _, found := telephoneDir[teleNum]; found {
				delete(telephoneDir, teleNum)
				fmt.Println("Telephone record deleted successfully")
			} else {
				fmt.Println("Telephone number not found in directory")
			}
		case 3:
			fmt.Println("Enter telephone number which is to be searched: ")
			fmt.Scanln(&teleNum)
			if tele, found := telephoneDir[teleNum]; found {
				fmt.Println("Telephone Number:", teleNum, ", Name:", tele.name, ", Address:", tele.address)
			} else {
				fmt.Println("Telephone number not found in directory")
			}
		case 4:
			fmt.Println("Enter telephone number whose details is to be updated: ")
			fmt.Scanln(&teleNum)
			if _, found := telephoneDir[teleNum]; found {
				fmt.Println("Enter the updated name: ")
				fmt.Scanln(&name)
				fmt.Println("Enter the updated address")
				fmt.Scanln(&address)
				telephoneDir[teleNum] = telephoneRecord{
					name:    name,
					address: address,
				}
				fmt.Println("Telephone record updates successfully")
			} else {
				fmt.Println("Telephone number not found in directory")
			}
		case 5:
			for telNum, telDetails := range telephoneDir {
				fmt.Println("Telephone Number:", telNum, ", Name:", telDetails.name, ", Address:", telDetails.address)
			}
		default:
			fmt.Println("Invalid input")
		}
		fmt.Println("Do you want to continue(Y=1/N=0): ")
		fmt.Scanln(&cntnu)
	}
}
