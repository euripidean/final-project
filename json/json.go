// Package json provides functions to create a JSON file from the data read from the spreadsheet (slice of maps)
package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// CreateJSON creates a JSON file from the data read from the spreadsheet
func CreateJSON(data []map[string]interface{}, fileName *string) {
	//if data is empty, throw an error and return
	if len(data) == 0 {
		error := fmt.Errorf("data cannot be empty")
		log.Println(error)
		return
	}
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Unable to convert data to JSON: %v", err)
	}

	file, err := os.Create(*fileName + ".json")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(jsonData))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// Ensure all operations have been applied to the underlying writer
	if err := writer.Flush(); err != nil {
		error := fmt.Errorf("failed to flush writer: %v", err)
		log.Println(error)
	} else {
		log.Println("JSON file created successfully")
	}
}
