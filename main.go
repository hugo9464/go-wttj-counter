package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	jobs := readFile("technical-test-jobs.csv")
	professions := readFile("technical-test-professions.csv")
	output := getOutput(jobs, professions)
	//fmt.Print(output)
	printOutput(output)
}

func printOutput(output map[Key]int) {
	var keys []Key
	for k := range output {
		//fmt.Println(k)
		keys = append(keys, k)
	}

	var categories []string

	for i := range keys {
		contract := keys[i].Contract
		categories = AppendIfMissing(categories, keys[i].Category)
		fmt.Println(contract)
	}

	printFirstRow(categories)
}

func printFirstRow(categories []string) {
	for i := range categories {
		fmt.Print("| ")
		fmt.Print(categories[i])
		fmt.Print(" ")
	}
	fmt.Print("|")
}

// `getOutput` gets the map keys contract by category output to print
func getOutput(jobs [][]string, professions [][]string) map[Key]int {

	output := make(map[Key]int)

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
		professionId, err := strconv.Atoi(jobs[i][0])

		//if no professionId for the job, we increment "other" category
		if err != nil {
			output[Key{contract, "Other"}]++
			continue
		}

		category := getCategory(professions, professionId)

		output[Key{"TOTAL", "TOTAL"}]++
		output[Key{"TOTAL", category}]++
		output[Key{contract, "TOTAL"}]++
		output[Key{contract, category}]++
	}

	return output
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

type Key struct {
	Contract, Category string
}