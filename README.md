# final-project | Simple Data Visualization Utility

<!-- Go Report Badge -->



<!-- Go Report Badge -->

## Description

This package allows you to visualize data from a Google Sheets spreadsheet in a web browser. You can pass the ID of the Google Sheets spreadsheet as a flag at runtime, and the program will read the data from the spreadsheet, marshal it into a JSON object, and save it to a file. The program will then read the JSON object and display the data in an HTML bar chart using D3.js.

## How to Install

1. Clone the repository
2. Run `go build` in the root directory

3. Run the executable using `go run main.go` with the following flags:
   - `-sid` (optional): The ID of the Google Sheets spreadsheet you want to visualize. If you don't pass it as a flag, the default value will be the value of the `SPREADSHEET_ID` environment variable.
   - `-f` (optional): The name of the file where the JSON data will be saved (default: `data.json`). This will also be the name of the HTML file that will be served.
4. Open a web browser and navigate to `http://localhost:3030/<file>.html` to see the visualizations

## Environment Variables

```
`API_KEY` - Your Google Cloud Console API key, which must have the Google Sheets API enabled
`SPREADSHEET_ID` - The ID of the Google Sheets spreadsheet you want to visualize. This is optional, as you can pass it as a flag at runtime. However, if you don't pass it as a flag, the default value will be the value of this environment variable.
`READ_RANGE` - The range of the Google Sheets spreadsheet you want to read. I recommend using the range `Sheet1!A1:Z1000` to read the entire sheet.
```

## How to Use

1. Create a Google Cloud Console project and enable the Google Sheets API
2. Create an API key and set it as the value of the `API_KEY` environment variable
3. Create a Google Sheets spreadsheet and share it with the email address associated with the API key
4. Run the executable with the following flags:
   - `-sid` (required): The ID of the Google Sheets spreadsheet you want to visualize
   - `-file` (optional): The name of the file where the JSON data will be saved (default: `data.json`). This will also be the name of the HTML file that will be served.
5. Open a web browser and navigate to `http://localhost:5500/<file>.html` to see the visualizations

### Screenshot of output HTML
![Screenshot 2024-05-07 at 10 39 23 PM](https://github.com/euripidean/final-project/assets/33559193/ba9b8c76-aba8-4a4d-8d71-aa535efbecdd)
