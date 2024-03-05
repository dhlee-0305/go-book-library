package util

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LoadFile(csvFile string) *list.List {
	// Open the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(headers)

	dataList := list.New()
	// Read the remaining records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // Reached end of file
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		dataList.PushBack(record)
	}

	/*
		for e := dataList.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	*/

	return dataList
}
