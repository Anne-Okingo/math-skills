# Math-Skills
## Introduction

This project- math-skills is designed to calculate and print the average, median, variance, and standard deviation of a statistical population of data from a file. The data in the file is presented as a list of numbers, one per line.

## Usage
To run the program, use the following command:

bash
go run main.go <data_file>

Replace <data_file> with the path to the file containing the data.

Fuctions
Average
The average is calculated by summing all the values and dividing by the total number of values.
Median
The median is the middle value of the sorted list of numbers. If the list has an even number of values, the median is the average of the two middle values.
Variance
The variance is the average of the squared differences from the mean.
Standard Deviation
The standard deviation is the square root of the variance.
Output
The program prints the results of each calculation in the following format:

text
Average: <average_value>
Median: <median_value>
Variance: <variance_value>
Standard Deviation: <standard_deviation_value>

Example Output

text
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65

Language
This project is written in Go.

