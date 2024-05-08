// Package sheets provides functions to interact with Google Sheets API
package sheets

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsServiceInterface interface {
	SpreadsheetsValuesGet(spreadsheetId string, range_ string) (*sheets.ValueRange, error)
}
// SheetsService provides functions to interact with Google Sheets API
type SheetsService struct {
	service SheetsServiceInterface
	apiKey  string
}

type SheetsServiceWrapper struct {
	service *sheets.Service
}

func (s *SheetsServiceWrapper) SpreadsheetsValuesGet(spreadsheetId string, range_ string) (*sheets.ValueRange, error) {
    return s.service.Spreadsheets.Values.Get(spreadsheetId, range_).Do()
}

// NewSheetsService creates a new SheetsService
func NewSheetsService(apiKey string) (*SheetsService, error) {
	ctx := context.Background()
	service, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve sheets client: %v", err)
	}
	return &SheetsService{service: &SheetsServiceWrapper{service: service}, apiKey: apiKey}, nil
}

// GetSheetData retrieves data from a Google Sheet
func (s *SheetsService) GetSheetData(spreadsheetID string, readRange string) ([]string, []map[string]interface{}, error) {
	if s.service == nil {
		return nil, nil, fmt.Errorf("sheets service is not initialized")
	}
	// Get data from sheet
	resp, err := s.service.SpreadsheetsValuesGet(spreadsheetID, readRange)
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
