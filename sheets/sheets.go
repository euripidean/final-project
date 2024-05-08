// Package sheets provides functions to interact with Google Sheets API
package sheets

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// GetSheetData retrieves data from a Google Sheet
func GetSheetData(apiKey string, spreadsheetID string, readRange string) ([]string, []map[string]interface{}, error) {

	// Connection
	ctx := context.Background()

	// Connect to sheets client
	service, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to retrieve sheets client: %v", err)
	}

	// Get data from sheet
	resp, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	var headers []string
	var data []map[string]interface{}

	if len(resp.Values) > 0 {
    // Get the headers from the first row
    
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
return headers, data, nil
}
