// Package sheets provides functions to interact with Google Sheets API
package sheets

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// SheetsService provides functions to interact with Google Sheets API
type SheetsService struct {
	service *sheets.Service
	apiKey  string
}

// NewSheetsService creates a new SheetsService
func NewSheetsService(apiKey string) (*SheetsService, error) {
	ctx := context.Background()
	service, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve sheets client: %v", err)
	}
	return &SheetsService{service: service, apiKey: apiKey}, nil
}

// GetSheetData retrieves data from a Google Sheet
func (s *SheetsService) GetSheetData(spreadsheetID string, readRange string) ([]string, []map[string]interface{}, error) {
	// Get data from sheet
	resp, err := s.service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
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
