package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func parseCsv(filename string) ([]Problem) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully Opened CSV file")

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var problems []Problem

	for _, line := range csvLines {
		problem := Problem{
				question: line[0],
				answer: line[1],
		}
		problems = append(problems, problem)
	}

	csvFile.Close()
	return problems
}