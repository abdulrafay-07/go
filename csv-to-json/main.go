package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Extract first row
	headers := records[0]

	var jsonData []map[string]string
	for _, row := range records[1:] {
		rowMap := make(map[string]string)
		for i, val := range row {
			rowMap[headers[i]] = val
		}
		jsonData = append(jsonData, rowMap)
	}

	// Convert to json
	jsonOutput, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonOutput))
}
