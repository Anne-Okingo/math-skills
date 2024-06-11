package main

import (
	"bufio"
	"fmt"
	"math"
	"mathskills/function"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	// Ensuring the program only takes two arguments
	if len(args) != 2 {
		fmt.Println("usage : go run . <datafile.txt>")
		os.Exit(0)
	}

	dataFile := args[1]
	Suffix := ".txt"
	// checking for onlt .txt file as dataFile
	if !strings.HasSuffix(dataFile, Suffix) {
		fmt.Println("Error datafile should be <.txt>\n\nusage : go run . <datafile.txt>")
		os.Exit(0)
	} else {
		// Open the file for reading
		file, err := os.Open("data.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		// reading the opened file
		scanner := bufio.NewScanner(file)

		value := []float64{}
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			data, err := strconv.ParseFloat(line, 64)
			if err != nil {
				fmt.Println("Error in conversion", err)
				os.Exit(0)
			}
			value = append(value, data)
		}

		if len(value) == 0 {
			fmt.Println("Empty File")
			return
		}

		result1 := functions.Average(value)
		result2 := functions.Median(value)
		result3 := functions.Variance(value)
		result4 := functions.StandardDev(value)

		fmt.Printf("Average : %.0f\n", math.Round(result1))
		fmt.Printf("Median : %.0f\n", math.Round(result2))
		fmt.Printf("Variance : %.0f\n", math.Round(result3))
		fmt.Printf("Standard Deviation : %.0f\n", math.Round(result4))

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
