package main

import (
	"fmt"
	"mathskills/function"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	// A condition to ensure the program only takes 2 arguments
	if len(args) != 2 {
		fmt.Println("usage : go run . <datafile.txt>")
		return
	}

	filename := args[1]

	filecontent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	// checking if the file containing data is empty
	if len(filecontent) == 0 {
		fmt.Println("empty file")
		return
	}

	contentsplit := strings.Split(string(filecontent), "\n")

	values := []float64{}
	// checking if there are empty spaces/ unfilled data among the data set, if there is they are ignored
	for _, ch := range contentsplit {
		ch = strings.TrimSpace(ch)
		if ch == "" {
			continue
		}
		// converting each character(data) into float64
		dataset, err := strconv.ParseFloat(ch, 64)
		if err != nil {
			fmt.Println("Error in conversion", err)
			return
		}
		values = append(values, dataset)
	}

	result1 := functions.Average(values)
	result2 := functions.Median(values)
	result3 := functions.Variance(values)
	result4 := functions.StandardDev(values)

	fmt.Printf("Average : %.0f\n", result1)
	fmt.Printf("Median : %.0f\n", result2)
	fmt.Printf("Variance : %.0f\n", result3)
	fmt.Printf("Standard Deviation : %.0f\n", result4)
}
