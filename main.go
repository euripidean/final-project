package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content	  string
}


func main() {
    // load the .env file
   if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // set API key
	apiKey := os.Getenv("API_KEY")

    // context for connection
    ctx := context.Background()

    // connect to sheets client
    service, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // The ID of the spreadsheet you want to access. TODO: stretch goal to maybe get multiple spreadsheets from a folder?
    spreadsheetID := os.Getenv("SPREADSHEET_ID")

    // This way of implementing requires a range for the Get method. Meep moop.
    readRange := "Sheet1!A1:Z"

    fmt.Println("This is just here to keep fmt as an import when I save.")

   

    resp, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    var data []map[string]interface{}

if len(resp.Values) > 0 {
    // Get the headers from the first row
    var headers []string
    for _, header := range resp.Values[0] {
        headers = append(headers, header.(string))
    }

    // Get the data from the remaining rows
    for _, row := range resp.Values[1:] {
        rowData := make(map[string]interface{})
        for i, cell := range row {
            if i < len(headers) {
                rowData[headers[i]] = cell
            }
        }
        data = append(data, rowData)
    }
}

// Convert the data to JSON
jsonData, err := json.Marshal(data)
if err != nil {
    log.Fatalf("Unable to convert data to JSON: %v", err)
}

file, err := os.Create("data.json")
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
    log.Fatalf("Failed to flush writer: %v", err)
}

visualize("data.json", "data")

}

func visualize(TextFilePath string, TextFileName string) {

    page := Page{
        TextFilePath: TextFilePath,
        HTMLPagePath: fmt.Sprintf("%s.html", TextFilePath),
        Content: "",
    }

    fileContents, err := os.ReadFile(page.TextFilePath)
    if err != nil {
        log.Fatalf("Failed to read file: %v", TextFileName)
    }

    chartData := string(fileContents)

    page.Content = chartData

    t := template.Must(template.New("template.tmpl").ParseFiles("template.templ"))

    newFile, err := os.Create(page.HTMLPagePath)
    if err != nil {
        log.Fatalf("Failed to create HTML file: %v", err)
    }

    t.Execute(newFile, page)

}
