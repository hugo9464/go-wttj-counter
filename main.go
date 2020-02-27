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
	jobs := readFile("technical-test-jobs.csv")
	professions := readFile("technical-test-professions.csv")
	printContractByCategory(jobs, professions)
}

// Prints the contract by category table, based on given jobs and professions
func printContractByCategory(jobs [][]string, professions [][]string) {

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
		contracts = appendIfMissing(contracts, contract)

		// if no contract for the job, we increment "Other" contract
		if contract == "" {
			contract = "Other"
		}
		professionId, _ := strconv.Atoi(jobs[i][0])

		category := getCategory(professions, professionId)
		categories = appendIfMissing(categories, category)

		// Once we have the contract and the category, with increment it in the outputMap
		outputMap = incrementValue(outputMap, contract, category)
		outputMap = incrementValue(outputMap, "TOTAL", category)
		outputMap = incrementValue(outputMap, contract, "TOTAL")
		outputMap = incrementValue(outputMap, "TOTAL", "TOTAL")
	}

	printOutput(outputMap, categories, contracts)
}

// incrementValue increments the value for the given contract and category
// it returns the updated outputMap
func incrementValue(outputMap map[string][]Value, contract string, category string) map[string][]Value {

	var values = outputMap[contract]
	// if no values found for the contract, we create a new slice of values
	if values == nil {
		values = createNewValues(category)
		outputMap[contract] = values
		return outputMap
	}

	values = appendValue(values, category)
	outputMap[contract] = values

	return outputMap
}

// appendValue increments the count for the given category and append it in the slice of values
// it returns the update slice of values
func appendValue(values []Value, category string) []Value {
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

// createNewValues creates a new slice of values for the given category
// it returns the new slice of values
func createNewValues(category string) []Value {
	var values []Value
	newValue := Value{category, 1}
	return append(values, newValue)
}

// printOutput prints the given output map
// it respects the order for the given categories and contracts slices
func printOutput(output map[string][]Value, categories []string, contracts []string) {

	// we get the longest contract size to add needed spaces for formating
	var longestContractSize = longestStringSize(contracts)
	var firstRow = getFirstRow(categories, longestContractSize)

	fmt.Println(getStringOfChar("-", len(firstRow)))
	fmt.Println(firstRow)
	fmt.Println(getStringOfChar("-", len(firstRow)))

	for _, contract := range contracts {
		fmt.Print("|" + contract + getStringOfChar(" ", longestContractSize - len(contract)) + "|")
		printContractValues(output[contract], categories)
		fmt.Println(getStringOfChar("-", len(firstRow)))
	}

}

// printContractValues prints Values as row base on the order of the given categories
func printContractValues(values []Value, categories []string) {
	for _, category := range categories {
		printValueForCategory(values, category)
	}
	fmt.Print("\n")
}

// printValueForCategory prints the value for the given category
// formats the cell with need spaces before and after the value
func printValueForCategory(values []Value, category string) {
	var count int
	for _, value := range values {
		if value.category == category {
			count = value.count
			continue
		}
	}
	var categoryLength = utf8.RuneCountInString(category)+2
	var spacesToAdd = (categoryLength - countDigits(count)) / 2

	fmt.Print(getStringOfChar(" ", spacesToAdd))
	fmt.Print(count)
	if (spacesToAdd*2 + countDigits(count)) < categoryLength {
		spacesToAdd++
	}
	fmt.Print(getStringOfChar(" ", spacesToAdd) + "|")
}

// countDigits counts the digits for the given number
func countDigits(i int) int {
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

// getFirstRow gets the formatted first row for the given categories
// it adds an empty cell at the beginning based on the given longestContractSize
func getFirstRow(categories []string, longestContractSize int) string {
	var row string
	row += "|" + getStringOfChar(" ", longestContractSize)
	for i := range categories {
		row += "| "
		row += categories[i]
		row += " "
	}
	row += "|"
	return row
}

// getStringOfChar returns a string of the given char
func getStringOfChar(s string, length int) string{
	var spaces string
	for i := 0; i < length; i++ {
		spaces += s
	}
	return spaces
}

// longestStringSize gets the longest string size between the given strings
func longestStringSize(strings []string) int {
	var longestString string
	for _, s := range strings {
		if len(s) > len(longestString) {
			longestString = s
		}
	}

	return len(longestString)
}

// getCategory gets the job category for the given profession id
func getCategory(professions [][]string, id int) string {

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

// readFile takes a filename and returns a two-dimensional slice of strings
func readFile(name string) [][]string {

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

// appendIfMissing appends a value to the given slice if missing
func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

type Value struct {
	category string
	count int
}