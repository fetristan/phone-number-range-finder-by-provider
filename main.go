package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Structure for a phone number and its provider.
type PhoneEntry struct {
	Number   int64
	Provider string
}

// Structure to represent a range of numbers with the count of numbers.
type Range struct {
	Start    int64
	End      int64
	Provider string
	Count    int64 // Number of numbers in the range
}

// Function to read phone numbers and providers from a CSV file.
func readPhoneData(filePath string) ([]PhoneEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Use semicolon as the field delimiter
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var phoneData []PhoneEntry
	for _, record := range records {
		number, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return nil, err
		}
		phoneData = append(phoneData, PhoneEntry{Number: number, Provider: record[1]})
	}
	return phoneData, nil
}

// Function to find ranges of numbers.
func findRanges(phoneData []PhoneEntry) []Range {
	sort.Slice(phoneData, func(i, j int) bool {
		if phoneData[i].Provider == phoneData[j].Provider {
			return phoneData[i].Number < phoneData[j].Number
		}
		return phoneData[i].Provider < phoneData[j].Provider
	})

	var ranges []Range
	if len(phoneData) == 0 {
		return ranges
	}

	for i := 0; i < len(phoneData); {
		start := phoneData[i]
		end := start

		for i++; i < len(phoneData) && phoneData[i].Provider == start.Provider && phoneData[i].Number == end.Number+1; i++ {
			end = phoneData[i]
		}

		count := end.Number - start.Number + 1
		ranges = append(ranges, Range{Start: start.Number, End: end.Number, Provider: start.Provider, Count: count})
	}

	return ranges
}

// Function to write ranges to a CSV file.
func writeRangesToCSV(ranges []Range, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV file header
	header := []string{"Start", "End", "Provider", "Range Length"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write the data
	for _, r := range ranges {
		record := []string{
			fmt.Sprintf("%013d", r.Start),
			fmt.Sprintf("%013d", r.End),
			r.Provider,
			strconv.FormatInt(r.Count, 10),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	filePath := "./did.csv" // Ensure the file path is correct
	phoneData, err := readPhoneData(filePath)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	ranges := findRanges(phoneData)
	exportFilePath := "./ranges_export.csv" // Path for the new CSV file
	if err := writeRangesToCSV(ranges, exportFilePath); err != nil {
		fmt.Println("Error writing CSV file:", err)
		return
	}

	fmt.Println("CSV file generated successfully:", exportFilePath)
}
