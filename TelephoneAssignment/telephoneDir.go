package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type telephoneRecord struct {
	Contact int    `json:"Contact"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

var telephoneDir []telephoneRecord

func addRecord(teleNum int, Name string, Address string) {
	//Accept details and add to slice
	telephoneDir = append(telephoneDir, telephoneRecord{
		Contact: teleNum,
		Name:    Name,
		Address: Address,
	})
	fmt.Println("Telephone record added succeddfully")
	createCSV()
}

func searchRecord(teleNum int) {
	//Linearly search in slice for the Contacts which match teleNum
	found := false
	for _, teleRecord := range telephoneDir {
		if teleRecord.Contact == teleNum {
			fmt.Println("Telephone Number:", teleRecord.Contact, ", Name:", teleRecord.Name, ", Address:", teleRecord.Address)
			found = true
		}
	}
	if !found {
		fmt.Println("Telephone number not found in directory")
	}
}

func deleteRecord(teleNum int) {
	//To store all the records having Contact as teleNum
	findings := []telephoneRecord{}
	//To store indices of the found duplicate records
	findingsIdx := []int{}
	for idx, teleRecord := range telephoneDir {
		if teleRecord.Contact == teleNum {
			findings = append(findings, teleRecord)
			findingsIdx = append(findingsIdx, idx)
		}
	}
	if len(findings) == 0 {
		fmt.Println("Telephone number not found in directory")
	} else if len(findings) == 1 {
		telephoneDir = append(telephoneDir[:findingsIdx[0]], telephoneDir[(findingsIdx[0]+1):]...)
	} else if len(findings) > 1 {
		//Ask for which specific record should be deleted
		var delIndex int
		for i := 0; i < len(findings); i++ {
			fmt.Println("Index:", findingsIdx[i], "Telephone Number:", findings[i].Contact, ", Name:", findings[i].Name, ", Address:", findings[i].Address)
		}
		fmt.Println("Following are the matched findings, Enter the index of record to be deleted: ")
		fmt.Scanln(&delIndex)
		telephoneDir = append(telephoneDir[:delIndex], telephoneDir[(delIndex+1):]...)
	}
	createCSV()
}

func displayRecord() {
	for _, teleRecord := range telephoneDir {
		fmt.Println("Telephone Number:", teleRecord.Contact, ", Name:", teleRecord.Name, ", Address:", teleRecord.Address)
	}
}

func updateRecord(teleNum int) {
	//To store all the records having Contact as teleNum
	findings := []telephoneRecord{}
	//To store indices of the found duplicate records
	findingsIdx := []int{}
	for idx, teleRecord := range telephoneDir {
		if teleRecord.Contact == teleNum {
			findings = append(findings, teleRecord)
			findingsIdx = append(findingsIdx, idx)
		}
	}
	if len(findings) == 0 {
		fmt.Println("Telephone number not found in directory")
	} else if len(findings) == 1 {
		fmt.Println("Enter the updated Contact number: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].Contact)
		fmt.Println("Enter the updated Name: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].Name)
		fmt.Println("Enter the updated Address: ")
		fmt.Scanln(&telephoneDir[findingsIdx[0]].Address)
	} else if len(findings) > 1 {
		//Ask for which specific record should be updated
		var updateIndex int
		for i := 0; i < len(findings); i++ {
			fmt.Println("Index:", findingsIdx[i], "Telephone Number:", findings[i].Contact, ", Name:", findings[i].Name, ", Address:", findings[i].Address)
		}
		fmt.Println("Following are the matched findings, Enter the index of record to be updated: ")
		fmt.Scanln(&updateIndex)
		fmt.Println("Enter the updated Contact number: ")
		fmt.Scanln(&telephoneDir[updateIndex].Contact)
		fmt.Println("Enter the updated Name: ")
		fmt.Scanln(&telephoneDir[updateIndex].Name)
		fmt.Println("Enter the updated Address: ")
		fmt.Scanln(&telephoneDir[updateIndex].Address)
	}
	createCSV()
}

func findDuplicates() {
	//To store Contact no. as key and associated duplicate record's indices as value
	duplicates := map[int][]int{}
	for idx, teleRecord := range telephoneDir {
		duplicates[teleRecord.Contact] = append(duplicates[teleRecord.Contact], idx)
	}
	//Print all the items having multiple indices in the value
	for _, indexs := range duplicates {
		if len(indexs) > 1 {
			for _, idx := range indexs {
				fmt.Println("Telephone Number:", telephoneDir[idx].Contact, ", Name:", telephoneDir[idx].Name, ", Address:", telephoneDir[idx].Address)
			}
		}
	}
}

func createCSV() {
	//Write entire telephoneDir slice to CSV
	csvFile, err := os.Create("./TelephoneDirectory.csv")
	if err != nil {
		log.Fatalln("Error occured while creating file:", err)
	}

	writer := csv.NewWriter(csvFile)
	err = writer.Write([]string{"Contact", "Name", "Address"})
	if err != nil {
		log.Fatalln("Error occured while writing to a file:", err)
	}
	writer.Flush()
	err = writer.Error()
	if err != nil {
		log.Fatalln("Error occured while writing to the file:", err)
	}

	for _, teleRecord := range telephoneDir {
		err = writer.Write([]string{strconv.Itoa(teleRecord.Contact), teleRecord.Name, teleRecord.Address})
		if err != nil {
			log.Fatalln("Error occured while writing to a file:", err)
		}
		writer.Flush()
		err = writer.Error()
		if err != nil {
			log.Fatalln("Error occured while writing to the file:", err)
		}
	}
}

func loadCSV() {
	//Read entire CSV and load all the data into telephoneDir
	telephoneDir = nil
	csvFile, err := os.Open("./TelephoneDirectory.csv")
	if err != nil {
		log.Fatalln("Error occured while opening file:", err)
	}

	reader := csv.NewReader(csvFile)

	_, err = reader.Read()
	if err != nil {
		log.Fatalln("Error occured while reading file:", err)
	}

	for true {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error occured while reading:", err)
		}
		ContactInt, _ := strconv.Atoi(record[0])
		telephoneDir = append(telephoneDir, telephoneRecord{
			Contact: ContactInt,
			Name:    record[1],
			Address: record[2],
		})
	}
}
