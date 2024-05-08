// Package visualization provides functions to visualize data in HTML format
package visualization

import (
	"fmt"
	"html/template"
	"os"
)

// Page represents a page to be visualized
type Page struct {
    JSONFilePath string
    HTMLPagePath string
    Content      string
}

// TemplateData holds the data to be passed to the template
type TemplateData struct {
    Title       string
    DataFile    string
    Headers     []string
    SvgWidth    int
    SvgHeight   int
    MarginTop   int
    MarginRight int
    MarginBottom int
    MarginLeft  int
    DataKey     string
    DomainKey   string
}

// Visualize creates an HTML page to visualize the data
func Visualize(FileName string, Headers []string, TemplatePath string) error {

    if len(Headers) <= 1 {
        return fmt.Errorf("there need to be at least two headers to visualize the data")
    }
    
    page := Page{
        JSONFilePath: fmt.Sprintf("%s.json",FileName),
        HTMLPagePath: fmt.Sprintf("%s.html", FileName),
        Content: "",
    }

    fileContents, err := os.ReadFile(page.JSONFilePath)
    if err != nil {
        error := fmt.Errorf("failed to read JSON file: %v", err)
        fmt.Println(error)
    }

    page.Content = string(fileContents)

    t, err := template.ParseFiles(TemplatePath)
    if err != nil {
        error := fmt.Errorf("failed to parse template: %v", err)
        fmt.Println(error)
    }

    newFile, err := os.Create(page.HTMLPagePath)
    if err != nil {
        error := fmt.Errorf("failed to create HTML file: %v", err)
        fmt.Println(error)
    }

    data := TemplateData{
        Title:       "D3 Visualization",
        DataFile:   page.JSONFilePath,
        Headers:    Headers,
        SvgWidth:    800,
        SvgHeight:   400,
        MarginTop:   20,
        MarginRight: 20,
        MarginBottom:30,
        MarginLeft:  40,
        DomainKey:   Headers[0],
    }

    err = t.Execute(newFile, data)
    if err != nil {
        error := fmt.Errorf("failed to execute template: %v", err)
        fmt.Println(error)
    } else {
        fmt.Println("Template executed successfully")
    }
    
    return nil
}
