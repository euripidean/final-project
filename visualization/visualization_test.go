package visualization

import (
	"os"
	"testing"
)

func TestVisualize(t *testing.T) {
    type args struct {
        FileName string
        Headers  []string
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            name: "Valid file name and headers",
            args: args{
                FileName: "testfile",
                Headers:  []string{"header1", "header2"},
            },
            wantErr: false,
        },
        {
            name: "Invalid file name",
            args: args{
                FileName: "/invalid/path",
                Headers:  []string{"header1", "header2"},
            },
            wantErr: true,
        },
        {
            name: "Empty headers",
            args: args{
                FileName: "testfile",
                Headers:  []string{},
            },
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create testfile.json
            if tt.args.FileName == "testfile" {
                err := os.WriteFile(tt.args.FileName+".json", []byte(`[{"header1":"value1","header2":"value2"}]`), 0644)
                if err != nil {
                    t.Fatal(err)
                }
                // Delete testfile.json after the test
                defer os.Remove(tt.args.FileName + ".json")
            }

            defer func() {
                if err := recover(); err != nil {
                    if tt.wantErr == false {
                        t.Errorf("Visualize() panic = %v, wantErr %v", err, tt.wantErr)
                    }
                }
            }()
            Visualize(tt.args.FileName, tt.args.Headers, "../template.tmpl")
            _, err := os.Stat(tt.args.FileName + ".html")
            if (err != nil) != tt.wantErr {
                t.Errorf("Visualize() error = %v, wantErr %v", err, tt.wantErr)
            }
            // Delete the created HTML file after the test
            if err == nil {
                os.Remove(tt.args.FileName + ".html")
            }
        })
    }
}
