# Math-Skills
## Introduction

This project- math-skills is designed to calculate and print the average, median, variance, and standard deviation of a statistical population of data from a file. The data in the file is presented as a list of numbers, one per line.

This program reads the file specified in the command line, then calculates the average, median, variance and standard deviation of the data in the file then finaly prints out the results as an output.

## Features
- statistical calculations are performed using appropriate functions. 
* The data represents statistical population.
* The values are rounded integers
* only integers are allowed
- Error handling is implemented to handle invalid data or file reading errors.
* This program takes any file with numeric values as long as its considered in the scope
* It only considers integers above intfloat64

## Fuctions
- Average
The average is calculated by summing all the values and dividing by the total number of values.
- Median
The median is the middle value of the sorted list of numbers. If the list has an even number of values, the median is the average of the two middle values.
- Variance
The variance is the average of the squared differences from the mean.
- Standard Deviation
The standard deviation is the square root of the variance.

Output
After reading the file, the program will execute each of the calculations asked and print the results in the following manner:
Example Output

Average: 3
Median: 3
Variance: 1
Standard Deviation: 1

 ## Dependencies 
 * The progrm is written in Go language
 - The program relies on the Go standard library and potentially any additional libraries or packages imported within the Go program


## Usage

1. To run this program run this command.
```bash
git clone https://learn.zone01kisumu.ke/git/aokingo/math-skills.git
```
2. 
```bash
cd math-skills
```
3. 
```bash
code. 
```
4. 
```bash
go run . data.txt
```

To test the program run this command
```bash
go test -v
``` 
## Author   [aokingo](https://learn.zone01kisumu.ke/git/aokingo).

