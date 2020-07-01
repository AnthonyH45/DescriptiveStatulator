package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData() ([]float64, error) {
	userData := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the sample data as comma separated values or whitespace separated values: ")
	str, readErr := userData.ReadString('\n')

	if readErr != nil {
		fmt.Println("Error reading user input")
		return nil, readErr
	}

	fmt.Println("You entered", str)
	str = strings.Replace(str, "\n", "", -1)

	var strSplit []string
	if strings.Contains(str, ",") {
		strSplit = strings.Split(str, ",")
	} else {
		strSplit = strings.Split(str, " ")
	}

	var toRet []float64
	for _, i := range strSplit {
		toAdd, convErr := strconv.ParseFloat(i, 64)

		if convErr != nil {
			fmt.Printf("Error converting %q to 64 bit base 10 integer", i)
			return []float64{}, convErr
		}
		toRet = append(toRet, toAdd)
	}

	return toRet, nil
}

func printStats(data []float64) error {
	// sort data
	sort.Float64s(data)
	fmt.Printf("%f\n", data)

	// printMin(data)
	min := data[0]
	fmt.Printf("Min: %f\n", min)

	// printMax(data)
	max := data[len(data)-1]
	fmt.Printf("Max: %f\n", max)

	// printRange
	rang := max - min
	fmt.Printf("Range: %f\n", rang)

	// printSize(data)
	size := len(data)
	fmt.Printf("Size: %d\n", size)

	/*
		// printSum(data)
		sum := func(data []float64) float64 {
			toRet := 0
			for _, i := range data {
				toRet += i
			}
			return toRet
		}
		s := sum(data)
		fmt.Printf("Sum: %d\n", s)

		// printMean(data)
		mean := float64(s) / float64(size)
		fmt.Printf("Mean: %f\n", mean)

		// printMedian(data)
		var median float64
		if len(data)%2 == 1 {
			median = float64(data[int(size/2.0)])
		} else {
			median = float64((data[int(size/2.0)-1] + data[int(size/2.0)])) / 2.0
		}
		fmt.Printf("Median: %f\n", median)

		// printMode(data)
		freq := func(data []int) int {
			freq := make([]int, max+1)
			for i := 0; i < len(data); i++ {
				freq[data[i]]++
			}

			max := freq[0]
			for i := 0; i < len(freq); i++ {
				if freq[i] > max {
					max = i
				}
			}
			return max
		}
		mode := freq(data)
		fmt.Printf("Mode: %d\n", mode)

		// printSD(data)
		var xDiff float64 = 0
		for _, i := range data {
			xDiff += math.Pow(float64(float64(i)-mean), float64(2))
		}
		variance := (xDiff / float64(size-1))
		SD := math.Sqrt(variance)
		fmt.Printf("Standard Deviatino: %f\n", SD)

		// printVariance(data)
		fmt.Printf("Variance: %f\n", variance)

		if size < 4 {
			fmt.Println("No quartile information if less than 4 data points")
			return nil
		}

		// printQ1(data)
		var Q1UL float64 = (0.25 * (float64(size) + float64(1)))
		Q1 := (math.Ceil(Q1UL) + math.Floor(Q1UL)) / float64(2)
		fmt.Printf("Q1: %f\n", Q1)

		// printQ2(data)
		Q2 := median
		fmt.Printf("Q2: %f\n", Q2)

		// printQ3(data)
		var Q3UL float64 = (0.75 * (float64(size) + float64(1)))
		Q3 := (math.Ceil(Q3UL) + math.Floor(Q3UL)) / float64(2)
		fmt.Printf("Q3: %f\n", Q3)

		// printIQR(data)
		IQR := Q3 - Q1
		fmt.Printf("IQR: %f\n", IQR)
	*/
	return nil
}

func main() {
	data, err := getData()

	if err != nil {
		fmt.Printf("Error getData(), %s", err.Error())
		os.Exit(1)
	}

	//fmt.Printf("You entered: %f\n", data)

	printStats(data)
}
