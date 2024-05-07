// Package main is the entry point of the program. It reads the data from the Google Sheet, creates a JSON file, and visualizes the data.
package main

import (
	"flag"
	"log"
	"os"

	"final-project/json"
	"final-project/sheets"
	"final-project/visualization"

	"github.com/joho/godotenv"
)

func main() {
    // load the .env file
   if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // set API key
	apiKey := os.Getenv("API_KEY")

    // set spreadsheet ID
    sid := flag.String("spreadsheetID", os.Getenv("SPREADSHEET_ID"), "Google Sheet ID")
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
    visualization.Visualize("data.json", "data", headers)
}
