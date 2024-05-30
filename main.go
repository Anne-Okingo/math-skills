// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"

// 	//"golang.org/x/text/number"
// 	functions "math-skills/function"
// )

// func main() {
// 	// fmt.Println(functions.Average([]int{1, 2, 3, 4}))
// 	args := os.Args

// 	if len(args) != 2 {
// 		fmt.Println("usage: go run <main.go><data.txt>")
// 		os.Exit(0)
// 	}
// 	filedata := os.Args[1]

// 	fileinfo, _ := os.Stat(filedata)
// 	if fileinfo.Size() == 0 {
// 		fmt.Println("Empty file")
// 		os.Exit(1)
// 	}

// 	file, err := os.ReadFile(filedata)
// 	if err != nil {
// 		fmt.Println("error reading file", err)
// 	}
// 	var values []float64
// 	filecontent := strings.Split(string(file), "\n")
// 	// if filecontent == nil {
// 	// 	return
// 	// }
// 	for _, num := range filecontent {
// 		if  num == "" {
// 			continue
// 		}
// 		 number, err := strconv.ParseFloat(num, 64)
// 		if err != nil {
// 			fmt.Println("error in conversion", err)
// 			return
// 		}
// 		values = append(values, number)
// 	}

// 	output := functions.HandleAverage(values)
// 	output1 := functions.Median(values)
// 	output2 := functions.HandleVariance(values)
// 	output3 := functions.StandardDev(values)

//		fmt.Printf("Average: %.0f\n ", output)
//		fmt.Printf("Median: %.0f\n", output1)
//		fmt.Printf("Variance: %.0f\n", output2)
//		fmt.Printf("Standard Deviation: %.0f\n ", output3)
//	}
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"mathskills/function"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("usage : go run . <datafile.txt>")
	}

	filename := args[1]

	// filestatus, _ := os.Stat(filename)
	// if filestatus.Size() == 0 {
	// 	fmt.Println("Error : empty file")
	// 	os.Exit(1)
	// }

	filecontent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	if len(filecontent) == 0{
		fmt.Println("empty file")
		return
	}

	contentsplit := strings.Split(string(filecontent), "\n")

	values := []float64{}
	for _, ch := range contentsplit {
		if ch == "" {
			continue
		}
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

	fmt.Printf("Average : %.0f\n",result1)
	fmt.Printf("Median : %.0f\n",result2)
	fmt.Printf("Variance : %.0f\n",result3)
	fmt.Printf("Standard Deviation : %.0f\n",result4)
}
