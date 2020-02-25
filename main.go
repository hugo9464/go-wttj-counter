package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Key struct {
	Contract, Category string
}

func main() {
	jobs := readFile("technical-test-jobs.csv")
	professions := readFile("technical-test-professions.csv")
	output := getOutput(jobs, professions)
	fmt.Print(output)
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
		jobid, _ := strconv.Atoi(jobs[i][0])
		category := getCategory(professions, jobid)

		output[Key{"TOTAL", "TOTAL"}]++
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

	//TODO: trouver comment gérer ce cas
	return "Profession ID not found"
}

// `readFile` takes a filename and returns a two-dimensional list of spreadsheet cells.
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
