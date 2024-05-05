package main

import (
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
    spreadsheetID := os.Getenv("SPREADSHEET_ID")

    // set read range
    readRange := os.Getenv("READ_RANGE")

    // get data from the sheet
    data, err := sheets.GetSheetData(apiKey, spreadsheetID, readRange)
    if err != nil {
        log.Fatalf("Failed to get sheet data: %v", err)
    }

    // create JSON file
    json.CreateJSON(data)

    // visualize the data
    visualization.Visualize("data.json", "data")
}
