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

// Visualize creates an HTML page to visualize the data
func Visualize(TextFilePath string, TextFileName string) {
	
	page := Page{
		TextFilePath: TextFilePath,
		HTMLPagePath: fmt.Sprintf("%s.html", TextFileName),
		Content: "",
	}

	fileContents, err := os.ReadFile(page.TextFilePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", TextFileName)
	}

	chartData := string(fileContents)

	page.Content = chartData

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		log.Fatalf("Failed to create HTML file: %v", err)
	}

	t.Execute(newFile, page)

}
