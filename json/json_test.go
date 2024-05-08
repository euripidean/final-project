// Package json provides functions to create a JSON file from the data read from the spreadsheet (slice of maps)
package json

import "testing"

// TestCreateJSON tests the CreateJSON function
func TestCreateJSON(t *testing.T) {
	type args struct {
		data []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateJSON(tt.args.data)
		})
	}
}
