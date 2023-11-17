package create

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"sync"
	// "strings"
)

func GenerateRandomData(numRows, numColumns int) [][]string {
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(numRows)
	// Seed the random number generator
	rand.NewSource(42)

	// Create column headers
	columns := []string{"QAH", "FTR", "XCK", "AFN", "AFQ", "OFP", "VAU", "SIE", "YIC", "CWP", "USN", "ZJO", "VQW", "PSB", "FHC", "GCH", "QJJ", "FGY", "QPE", "SEJ"}
	// Create column headers with random uppercase letters
	// var columns []string
	// for i := 0; i < numColumns; i++ {
	// 	columns = append(columns, randColumn())
	// }

	// Generate random data
	var data [][]string
	for i := 0; i < numRows; i++ {
		// var row []string
		go func(numColumns int) {
			defer wg.Done()
			var row []string
			for j := 0; j < numColumns; j++ {
				// row = append(row, strconv.Itoa(rand.Intn(100)+1))
				row = append(row, strconv.FormatFloat(rand.Float64()*100, 'f', 14, 64))
			}
			m.Lock()
			data = append(data, row)
			m.Unlock()
		}(numColumns)
	}
	wg.Wait()

	data = append([][]string{columns}, data...) // Add column headers as the first row
	return data
}

// func randColumn() string {
// 	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	var result strings.Builder
// 	for i := 0; i < 3; i++ {
// 		result.WriteByte(letters[rand.Intn(len(letters))])
// 	}
// 	return result.String()
// }

func WriteToCSVByRow(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func WriteAllToCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		return err
	}

	return nil
}
