package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func main() {
	jobs := readFile("technical-test-jobs.csv")
	professions := readFile("technical-test-professions.csv")
	printContractByCategory(jobs, professions)
}

func printContractByCategory(jobs [][]string, professions [][]string) {

	outputMap := make(map[string][]Value)
	var categories []string
	categories = append(categories, "TOTAL")

	for i := range jobs {

		// we can ignore the first row as this is the header
		if i == 0 {
			continue
		}

		contract := jobs[i][1]

		// if no contract for the job, we increment "Other" contract
		if contract == "" {
			contract = "Other"
		}
		professionId, _ := strconv.Atoi(jobs[i][0])

		category := getCategory(professions, professionId)
		categories = AppendIfMissing(categories, category)

		outputMap = incrementValue(outputMap, contract, category)
		outputMap = incrementValue(outputMap, "TOTAL", category)
		outputMap = incrementValue(outputMap, contract, "TOTAL")
		outputMap = incrementValue(outputMap, "TOTAL", "TOTAL")
	}

	printOutput(outputMap, categories)
}

func incrementValue(outputMap map[string][]Value, contract string, category string) map[string][]Value {

	var values = outputMap[contract]
	if values == nil {

		values = createNewValues(category)
		outputMap[contract] = values
		return outputMap
	}

	values = appendValue(values, category)
	outputMap[contract] = values

	return outputMap
}

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

func createNewValues(category string) []Value {
	var values []Value
	newValue := Value{category, 1}
	return append(values, newValue)
}

func printOutput(output map[string][]Value, categories []string) {
	var longestContract = longestString(mapKeys(output))
	var firstRow = getFirstRow(categories, longestContract)

	fmt.Println(getStringOfChar("-", len(firstRow)))
	fmt.Println(firstRow)
	fmt.Println(getStringOfChar("-", len(firstRow)))


	for contract := range output {
		fmt.Print("|" + contract + getStringOfChar(" ", len(longestContract) - len(contract)) + "|")
		printContractValues(output[contract], categories)
		fmt.Println(getStringOfChar("-", len(firstRow)))
	}

}

func mapKeys(output map[string][]Value) []string {
	keys := reflect.ValueOf(output).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}

	return strkeys
}

func printContractValues(values []Value, categories []string) {
	for _, category := range categories {
		printValueForCategory(values, category)
	}
	fmt.Print("\n")
}

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

func getFirstRow(categories []string, longestContract string) string {
	var row string
	row += "|" + getStringOfChar(" ", len(longestContract))
	for i := range categories {
		row += "| "
		row += categories[i]
		row += " "
	}
	row += "|"
	return row
}

func getStringOfChar(s string, length int) string{
	var spaces string
	for i := 0; i < length; i++ {
		spaces += s
	}
	return spaces
}

func longestString(strings []string) string {
	var longestString string
	for _, s := range strings {
		if len(s) > len(longestString) {
			longestString = s
		}
	}

	return longestString
}



// `getCategory` gets the job category for the given profession id
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

// `readFile` takes a filename and returns a two-dimensional list of strings
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
	count int
}