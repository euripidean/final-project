package visualization

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// Page represents a page to be visualized
type Page struct {
    TextFilePath string
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
func Visualize(TextFilePath string, TextFileName string, Headers []string) {
    
    page := Page{
        TextFilePath: TextFilePath,
        HTMLPagePath: fmt.Sprintf("%s.html", TextFileName),
        Content: "",
    }

    fileContents, err := os.ReadFile(page.TextFilePath)
    if err != nil {
        log.Fatalf("Failed to read file: %v", TextFileName)
    }

    page.Content = string(fileContents)

    t := template.Must(template.ParseFiles("template.tmpl"))

    newFile, err := os.Create(page.HTMLPagePath)
    if err != nil {
        log.Fatalf("Failed to create HTML file: %v", err)
    }

    data := TemplateData{
        Title:       "D3 Visualization",
        DataFile:   TextFilePath,
        Headers:    Headers,
        SvgWidth:    800,
        SvgHeight:   400,
        MarginTop:   20,
        MarginRight: 20,
        MarginBottom:30,
        MarginLeft:  40,
        DomainKey:   Headers[0],
    }

    t.Execute(newFile, data)
    log.Println("HTML file created successfully")
}
