// Package json provides functions to create a JSON file from the data read from the spreadsheet (slice of maps)
package json

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

// TestCreateJSON tests the CreateJSON function
func TestCreateJSON(t *testing.T) {
	var fileName = "data"
	type args struct {
		data     []map[string]interface{}
		fileName *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "JSON is created when data is supplied",
			args: args{
				data: []map[string]interface{}{
					{"key1": "value1", "key2": "value2"},
					{"key1": "value3", "key2": "value4"},
				},
				fileName: &fileName,
			},
			wantErr: false,
		},
		{
			name: "JSON is not created when data is not supplied",
			args: args{
				data:     []map[string]interface{}{},
				fileName: &fileName,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call CreateJSON
			CreateJSON(tt.args.data, tt.args.fileName)

			// Check if file was created
			_, err := os.Stat("data.json")
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {

				// Read file content
				content, err := os.ReadFile("data.json")
				if err != nil {
					t.Fatalf("Failed to read file: %v", err)
				}

				// Unmarshal content
				var gotData []map[string]interface{}
				err = json.Unmarshal(content, &gotData)
				if err != nil {
					t.Fatalf("Failed to unmarshal content: %v", err)
				}

				// Compare unmarshalled data with original data
				if !reflect.DeepEqual(tt.args.data, gotData) {
					t.Errorf("Data mismatch: got %v, want %v", gotData, tt.args.data)
				}

				// Clean up
				os.Remove("data.json")
			}
		})
	}
}

// BenchmarkCreateJSON benchmarks the CreateJSON function
func BenchmarkCreateJSON(b *testing.B) {
	// Create data
	data := []map[string]interface{}{
		{"key1": "value1", "key2": "value2"},
		{"key1": "value3", "key2": "value4"},
	}
	fileName := "data"

	for i := 0; i < b.N; i++ {
		CreateJSON(data, &fileName)
	}

	// Clean up
	os.Remove("data.json")
}
