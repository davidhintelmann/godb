package main

import (
	"fmt"
	"time"

	"github.com/davidhintelmann/godb/create"
)

func main() {
	// Two variables below will generate
	// 1.00 GB of data in csv format
	// numRows := 3_000_000
	// numColumns := 20
	numRows := 3_000_000
	numColumns := 20
	start_data := time.Now()
	data := create.GenerateRandomData(numRows, numColumns)
	end_data := time.Now()
	duration_data := end_data.Sub(start_data)
	fmt.Printf("Duration of create.GenerateRandomData(): %v\n", duration_data)

	start_write := time.Now()
	csvFilename := "sample.csv"
	if err := create.WriteToCSVByRow(csvFilename, data); err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}
	end_write := time.Now()
	duration_write := end_write.Sub(start_write)
	fmt.Printf("Duration of create.WriteToCSVByRow(): %v\n", duration_write)

	fmt.Printf("Data written to %s\n", csvFilename)
}

// Measure-Command {Start-Process -FilePath "go" -WorkingDirectory "C:\Users\david\go\src\github.com\davidhintelmann\godb" -ArgumentList "run main.go" -NoNewWindow -wait}
// TotalSeconds: 18.0417504

// Measure-Command {start-process .\main.exe -NoNewWindow -wait}
// TotalSeconds: 18.0623932
