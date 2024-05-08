// Package main is the entry point of the program. It reads the data from the Google Sheet, creates a JSON file, and visualizes the data.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/euripidean/final-project/json"
	"github.com/euripidean/final-project/sheets"
	"github.com/euripidean/final-project/visualization"

	"github.com/joho/godotenv"
)

// main executes the program to read data from a Google Sheet, create a JSON file, and visualize the data
func main() {
    // load the .env file
   if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // set API key
	apiKey := os.Getenv("API_KEY")

    // set spreadsheet ID and the output file name
    sid := flag.String("spreadsheetID", os.Getenv("SPREADSHEET_ID"), "Google Sheet ID")
    file := flag.String("file", "data", "JSON and HTML file to create")
    flag.Parse()

    // set read range
    readRange := os.Getenv("READ_RANGE")

    // get data from the sheet
    headers, data, err := sheets.GetSheetData(apiKey, *sid, readRange)
    if err != nil {
        log.Fatalf("Failed to get sheet data: %v", err)
    }

    // create JSON file
    json.CreateJSON(data)

    // visualize the data
    visualization.Visualize(*file, headers, "./template.tmpl")
}
