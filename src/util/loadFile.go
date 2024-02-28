package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"models"
	"os"
	"strconv"
)

func LoadFile(csvFile string) {
	// Open the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(headers)

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

		book := models.Book{}
		bookId, _ := strconv.ParseInt(record[0], 10, 64)
		book.SetBook(bookId, record[1], record[2], record[3], record[4])
		book.Print()
	}

	// Print the hash map content
	/*
		for key, value := range data {
			fmt.Println("Key:", key, "Value:", value)
		}
	*/
}
