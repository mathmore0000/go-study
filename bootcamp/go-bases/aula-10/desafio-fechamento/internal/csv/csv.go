package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read input file %s: %v", filePath, err))
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("Unable to parse file as CSV for %s: %v", filePath, err))
	}

	return records
}
