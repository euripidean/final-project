# Final Project Proposal

For my Final Go Project, I want to find an easy way to visualize simple spreadsheet data in a web browser. I want to be able to upload a CSV file and have the data displayed in a bar chart.

## Project Description

Oftentimes you want to share a simple visualization of data with someone who doesn't need full access to the spreadsheet, just needs a quick overview (or is likely to break your forumlae if you give them access to the spreadsheet). This project will allow you to upload a CSV file and have the data displayed in a bar chart in a web browser.

## Features

- Pass through a Google Sheets `Spreadsheet ID` at runtime
- The program will then read the data from the Google Sheet
- It will then Marshal the data into a JSON object and save it to a file
- The program will then read the JSON object and display the data in an HTML bar chart using D3.js

## Technologies

- Go
- Google Sheets API
- D3.js

## Limitations

- The program will only work with Google Sheets and (currently) only with the specified `Spreadsheet ID`
- The visualization will only be a simple bar chart
- Spreadsheet data needs to be in a specific format (e.g. x columns, with the first being the labels and the others being the values)

## Learning Goals

- Learn how to use the Google Sheets API
- Implement a simple chart visualization using D3.js
- Learn how to marshal data into JSON
- Learn how to read and write files in Go
- Learn how to serve a web page from a Go server
- Use the `flag` package to pass arguments to the program

## Stretch Goals

Potential future development ideas include:

- Allow the user to select the type of chart they want to display
- Allow the user to customize the chart's appearance
- Use the Google Drive API to walk through all spreadsheets in a folder and perform the same operation on each one
