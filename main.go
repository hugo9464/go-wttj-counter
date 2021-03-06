package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	jobs := ReadFile("technical-test-jobs.csv")
	professions := ReadFile("technical-test-professions.csv")
	outputString := GetContractByCategory(jobs, professions)
	fmt.Println(outputString)
}

// GetContractByCategory gets the contract by category table, based on given jobs and professions
// It returns the formatted contract by category table
func GetContractByCategory(jobs [][]string, professions [][]string) string {

	outputMap := make(map[string][]Value)
	var categories []string
	categories = append(categories, "TOTAL")
	var contracts []string
	contracts = append(contracts, "TOTAL")

	for i := range jobs {

		// we can ignore the first row as this is the header
		if i == 0 {
			continue
		}

		contract := jobs[i][1]
		contracts = AppendIfMissing(contracts, contract)

		// if no contract for the job, we increment "Other" contract
		if contract == "" {
			contract = "Other"
		}
		professionId, _ := strconv.Atoi(jobs[i][0])

		category := GetCategory(professions, professionId)
		categories = AppendIfMissing(categories, category)

		// Once we have the contract and the category, with increment it in the outputMap
		outputMap = IncrementValue(outputMap, contract, category)
		outputMap = IncrementValue(outputMap, "TOTAL", category)
		outputMap = IncrementValue(outputMap, contract, "TOTAL")
		outputMap = IncrementValue(outputMap, "TOTAL", "TOTAL")
	}

	return GetOutput(outputMap, categories, contracts)
}

// IncrementValue increments the value for the given contract and category
// it returns the updated outputMap
func IncrementValue(outputMap map[string][]Value, contract string, category string) map[string][]Value {

	var values = outputMap[contract]
	// if no values found for the contract, we create a new slice of values
	if values == nil {
		values = CreateNewValues(category)
		outputMap[contract] = values
		return outputMap
	}

	values = AppendValue(values, category)
	outputMap[contract] = values

	return outputMap
}

// AppendValue increments the count for the given category and append it in the slice of values
// it returns the update slice of values
func AppendValue(values []Value, category string) []Value {
	var value Value
	for i := range values {
		value = values[i]

		if value.category == category {
			value.count++
			values[i] = value
			return values
		}
	}
	value = Value{category, 1}
	return append(values, value)
}

// CreateNewValues creates a new slice of values for the given category with count=1
// it returns the new slice of values
func CreateNewValues(category string) []Value {
	var values []Value
	newValue := Value{category, 1}
	return append(values, newValue)
}

// GetOutput gets the given output map
// it respects the order for the given categories and contracts slices
// It returns the formatted output map as string
func GetOutput(output map[string][]Value, categories []string, contracts []string) string {

	var outputString string

	// we get the longest contract size to add needed spaces for formating
	var longestContractSize = LongestStringSize(contracts)
	var firstRow = GetFirstRow(categories, longestContractSize)

	outputString += GetStringOfChar("-", len(firstRow)) + "\n"
	outputString += firstRow + "\n"
	outputString += GetStringOfChar("-", len(firstRow)) + "\n"

	for _, contract := range contracts {
		outputString += "|" + contract + GetStringOfChar(" ", longestContractSize-len(contract)) + "|"
		outputString += GetContractValues(output[contract], categories)
		outputString += GetStringOfChar("-", len(firstRow)) + "\n"
	}

	return outputString
}

// GetContractValues gets Values as row based on the order of the given categories
// It returns the given values as string
func GetContractValues(values []Value, categories []string) string {
	var contractValues string
	for _, category := range categories {
		contractValues += GetValueForCategory(values, category)
	}
	return contractValues + "\n"
}

// GetValueForCategory gets the value for the given category
// formats the cell with need spaces before and after the value
// It returns the value as string
func GetValueForCategory(values []Value, category string) string {
	var valueForCategory string
	var count int
	for _, value := range values {
		if value.category == category {
			count = value.count
			continue
		}
	}
	var categoryLength = utf8.RuneCountInString(category) + 2
	var spacesToAdd = (categoryLength - CountDigits(count)) / 2

	valueForCategory += GetStringOfChar(" ", spacesToAdd) + strconv.Itoa(count)
	if (spacesToAdd*2 + CountDigits(count)) < categoryLength {
		spacesToAdd++
	}

	return valueForCategory + GetStringOfChar(" ", spacesToAdd) + "|"
}

// CountDigits counts the digits for the given number
func CountDigits(i int) int {
	var count int
	if i == 0 {
		return 1
	}
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

// GetFirstRow gets the formatted first row for the given categories
// it adds an empty cell at the beginning based on the given longestContractSize
// It returns the first row as string
func GetFirstRow(categories []string, longestContractSize int) string {
	var row string
	row += "|" + GetStringOfChar(" ", longestContractSize)
	for i := range categories {
		row += "| "
		row += categories[i]
		row += " "
	}
	row += "|"
	return row
}

// GetStringOfChar returns a string of the given char
func GetStringOfChar(s string, length int) string {
	var chars string
	for i := 0; i < length; i++ {
		chars += s
	}
	return chars
}

// LongestStringSize gets the longest string size between the given strings
func LongestStringSize(strings []string) int {
	var longestString string
	for _, s := range strings {
		if len(s) > len(longestString) {
			longestString = s
		}
	}

	return len(longestString)
}

// GetCategory gets the job category for the given profession id
func GetCategory(professions [][]string, id int) string {

	for i := range professions {

		// we can ignore the first row as this is the header
		if i == 0 {
			continue
		}

		professionid, _ := strconv.Atoi(professions[i][0])
		if professionid == id {
			return professions[i][2]
		}
	}

	return "Other"
}

// ReadFile takes a filename and returns a two-dimensional slice of strings
func ReadFile(name string) [][]string {

	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}

// AppendIfMissing appends a value to the given slice if missing
func AppendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

type Value struct {
	category string
	count    int
}
