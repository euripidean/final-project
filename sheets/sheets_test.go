// Package sheets provides functions to interact with Google Sheets API
package sheets

import (
	"reflect"
	"testing"

	"google.golang.org/api/sheets/v4"
)

type MockSheetsService struct {
	response *sheets.ValueRange
}

func (m *MockSheetsService) SpreadsheetsValuesGet(spreadsheetId string, range_ string) (*sheets.ValueRange, error) {
	return m.response, nil
}

// Test helper function to compare two slices of maps as Go does not allow direct comparison of slices of maps
func compareSliceOfMaps(a, b []map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, aa := range a {
		found := false
		for _, bb := range b {
			if reflect.DeepEqual(aa, bb) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestSheetsService_GetSheetData(t *testing.T) {

	type args struct {
		spreadsheetID string
		readRange     string
	}
	tests := []struct {
		name    string
		service SheetsServiceInterface
		args    args
		want    []string
		want1   []map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test GetSheetData returns the data from the sheet",
			service: &MockSheetsService{
				response: &sheets.ValueRange{
					Values: [][]interface{}{
						{"Month", "Number Sold"},
						{"January", 25},
						{"February", 30},
					},
				},
			},
			args: args{
				spreadsheetID: "spreadsheetID",
				readRange:     "readRange",
			},
			want:    []string{"Month", "Number Sold"},
			want1:   []map[string]interface{}{{"Month": "January", "Number Sold": 25}, {"Month": "February", "Number Sold": 30}},
			wantErr: false,
		},
		// Ran out of time to write more tests as this was a pretty complicated set up!
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SheetsService{
				service: tt.service,
			}
			got, got1, err := s.GetSheetData(tt.args.spreadsheetID, tt.args.readRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("SheetsService.GetSheetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SheetsService.GetSheetData() got = %v, want %v", got, tt.want)
			}
			if tt.want != nil && !compareSliceOfMaps(got1, tt.want1) {
				t.Errorf("SheetsService.GetSheetData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
