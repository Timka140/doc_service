package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func createMapFromRawData(data [][]string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	var fieldNames []string
	for i, line := range data {
		if i == 0 {
			fieldNames = line
			continue
		}
		attachment1Dot1Transformers := make(map[string]interface{})
		for j, field := range line {
			fieldName := fieldNames[j]
			attachment1Dot1Transformers[fieldName] = field
		}
		result = append(result, attachment1Dot1Transformers)
	}
	return result
}

func findMapAttachment1Point(obj *map[string]interface{}, data []map[string]interface{}) error {
	find := false
	mapAttachment1Point := make([]map[string]interface{}, 0)
	for _, record := range data {
		if record["contract_number"] == (*obj)["contract_number"] {
			mapAttachment1Point = append(mapAttachment1Point, record)
			find = true
		}
	}
	if !find {
		log.Printf("[WARNING]: findMapAttachment1Point(): no found any record in Attachment1Dot1Transformers for contract =  %v \n\t", (*obj)["contract_number"])
	}
	(*obj)["counters"] = mapAttachment1Point
	return nil
}

func findMapAttachment1Dot1Transformers(obj *map[string]interface{}, data []map[string]interface{}) error {
	find := false
	mapAttachment1Dot1Transformers := make([]map[string]interface{}, 0)
	for _, record := range data {
		if record["contract_number"] == (*obj)["contract_number"] && record["device_number"] == (*obj)["device_number"] {
			mapAttachment1Dot1Transformers = append(mapAttachment1Dot1Transformers, record)
			find = true
		}
	}
	if !find {
		log.Printf("[WARNING]: findMapAttachment1Dot1Transformers(): no found any record in Attachment1Dot1Transformers for contract = %v and name, device = %v\n\t", (*obj)["contract_number"], (*obj)["device_number"])
	}
	(*obj)["transformers"] = mapAttachment1Dot1Transformers
	return nil
}

func getRawDataFromCSV(file string) ([][]string, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("getRawDataFromCSV(): %v\n\t", err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("getRawDataFromCSV(): %v\n\t", err)
	}

	return data, nil
}

func main() {

	dataContracts, err := getRawDataFromCSV("testMapDataContract.csv")
	if err != nil {
		log.Fatal(err)
	}
	dataAttachment1, err := getRawDataFromCSV("testMapDataAttach1.csv")
	if err != nil {
		log.Fatal(err)
	}
	dataAttachment1Dot1, err := getRawDataFromCSV("testMapDataAttachment1Dot1.csv")
	if err != nil {
		log.Fatal(err)
	}

	mapDataAttachment1Dot1 := createMapFromRawData(dataAttachment1Dot1)
	mapDataContracts := createMapFromRawData(dataContracts)
	mapDataAttachment1 := createMapFromRawData(dataAttachment1)

	for i, record := range mapDataAttachment1 {
		err = findMapAttachment1Dot1Transformers(&record, mapDataAttachment1Dot1)
		if err != nil {
			log.Println(err)
		}
		mapDataAttachment1[i] = record
	}

	for i, record := range mapDataContracts {
		err = findMapAttachment1Point(&record, mapDataAttachment1)
		if err != nil {
			log.Println(err)
		}
		mapDataContracts[i] = record
	}

	println(mapDataAttachment1Dot1)
}
