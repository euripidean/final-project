package json

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestCreateJSON(t *testing.T) {
	// Test data
	data := []map[string]interface{}{
		{"name": "John Doe", "age": 30},
		{"name": "Jane Smith", "age": 25},
	}

	// Create JSON
	CreateJSON(data)

	// Verify if the JSON file is created
	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		t.Errorf("JSON file was not created")
	}

	// Read the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the JSON data
	var jsonData []map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jsonData); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Verify the decoded JSON data
	expectedData := []map[string]interface{}{
		{"name": "John Doe", "age": 30},
		{"name": "Jane Smith", "age": 25},
	}
	if !reflect.DeepEqual(jsonData, expectedData) {
		t.Errorf("Decoded JSON data does not match expected data")
	}
}
