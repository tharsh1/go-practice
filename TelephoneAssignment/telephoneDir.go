package main

import "fmt"

type telephoneRecord struct {
	contact int
	name    string
	address string
}

var telephoneDir []telephoneRecord

func addRecord(teleNum int, name string, address string) {
	telephoneDir = append(telephoneDir, telephoneRecord{
		contact: teleNum,
		name:    name,
		address: address,
	})
	fmt.Println("Telephone record added succeddfully")
}

func searchRecord(teleNum int) {
	found := false
	for _, teleRecord := range telephoneDir {
		if teleRecord.contact == teleNum {
			fmt.Println("Telephone Number:", teleRecord.contact, ", Name:", teleRecord.name, ", Address:", teleRecord.address)
			found = true
		}
	}
	if !found {
		fmt.Println("Telephone number not found in directory")
	}
}

func deleteRecord(teleNum int) {
	findings := []telephoneRecord{}
	findingsIdx := []int{}
	for idx, teleRecord := range telephoneDir {
		if teleRecord.contact == teleNum {
			findings = append(findings, teleRecord)
			findingsIdx = append(findingsIdx, idx)
		}
	}
	if len(findings) == 0 {
		fmt.Println("Telephone number not found in directory")
	} else if len(findings) == 1 {
		telephoneDir = append(telephoneDir[:findingsIdx[0]], telephoneDir[(findingsIdx[0]+1):]...)
	} else if len(findings) > 1 {
		var delIndex int
		for i := 0; i < len(findings); i++ {
			fmt.Println("Index:", findingsIdx[i], "Telephone Number:", findings[i].contact, ", Name:", findings[i].name, ", Address:", findings[i].address)
		}
		fmt.Println("Following are the matched findings, Enter the index of record to be deleted: ")
		fmt.Scanln(&delIndex)
		telephoneDir = append(telephoneDir[:delIndex], telephoneDir[(delIndex+1):]...)
	}
}

func displayRecord() {
	for _, teleRecord := range telephoneDir {
		fmt.Println("Telephone Number:", teleRecord.contact, ", Name:", teleRecord.name, ", Address:", teleRecord.address)
	}
}

func updateRecord(teleNum int) {
	findings := []telephoneRecord{}
	findingsIdx := []int{}
	for idx, teleRecord := range telephoneDir {
		if teleRecord.contact == teleNum {
			findings = append(findings, teleRecord)
			findingsIdx = append(findingsIdx, idx)
		}
	}
	if len(findings) == 0 {
		fmt.Println("Telephone number not found in directory")
	} else if len(findings) == 1 {
		fmt.Println("Enter the updated contact number: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].contact)
		fmt.Println("Enter the updated name: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].name)
		fmt.Println("Enter the updated address: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].address)
	} else if len(findings) > 1 {
		var updateIndex int
		for i := 0; i < len(findings); i++ {
			fmt.Println("Index:", findingsIdx[i], "Telephone Number:", findings[i].contact, ", Name:", findings[i].name, ", Address:", findings[i].address)
		}
		fmt.Println("Following are the matched findings, Enter the index of record to be updated: ")
		fmt.Scanln(&updateIndex)
		fmt.Println("Enter the updated contact number: ")
		fmt.Scanln(&telephoneDir[updateIndex].contact)
		fmt.Println("Enter the updated name: ")
		fmt.Scanln(&telephoneDir[updateIndex].name)
		fmt.Println("Enter the updated address: ")
		fmt.Scanln(&telephoneDir[updateIndex].address)
	}
}

func findDuplicates() {
	duplicates := map[int][]int{}
	for idx, teleRecord := range telephoneDir {
		duplicates[teleRecord.contact] = append(duplicates[teleRecord.contact], idx)
	}

	for _, indexs := range duplicates {
		if len(indexs) > 1 {
			for _, idx := range indexs {
				fmt.Println("Telephone Number:", telephoneDir[idx].contact, ", Name:", telephoneDir[idx].name, ", Address:", telephoneDir[idx].address)
			}
		}
	}
}
