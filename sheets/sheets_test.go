// Description: Unit tests for the sheets package
package sheets

import (
	"reflect"
	"testing"
)

func TestGetSheetData(t *testing.T) {
	apiKey := "your-api-key"
	spreadsheetID := "your-spreadsheet-id"
	readRange := "Sheet1!A1:C"

	headers, data, err := GetSheetData(apiKey, spreadsheetID, readRange)
	if err != nil {
		t.Errorf("Error retrieving sheet data: %v", err)
	}

	// Assert the expected number of headers
	expectedHeaders := []string{"Header1", "Header2", "Header3"}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("Expected headers %v, but got %v", expectedHeaders, headers)
	}

	// Assert the expected number of data rows
	expectedRowCount := 2
	if len(data) != expectedRowCount {
		t.Errorf("Expected %d data rows, but got %d", expectedRowCount, len(data))
	}

	// Add more assertions for the data values if needed
}
