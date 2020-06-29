package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData() ([]int, error) {
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

	var toRet []int
	for _, i := range strSplit {
		toAdd, convErr := strconv.ParseInt(i, 10, 0)

		if convErr != nil {
			fmt.Printf("Error converting %q to 64 bit base 10 integer", i)
			return []int{}, convErr
		}
		toRet = append(toRet, int(toAdd))
	}

	return toRet, nil
}

func printStats(data []int) error {
	// sort data
	sort.Ints(data)
	// printMin(data)
	min := data[0]
	fmt.Printf("Min: %d\n", min)
	// printMax(data)
	max := data[len(data)-1]
	fmt.Printf("Max: %d\n", max)
	// printRange
	rang := max - min
	fmt.Printf("Range: %d\n", rang)
	// printSize(data)
	size := len(data)
	fmt.Printf("Size: %d\n", size)
	// printSum(data)
	sum := func(data []int) int {
		toRet := 0
		for _, i := range data {
			toRet += i
		}
		return toRet
	}
	s := sum(data)
	fmt.Printf("Sum: %d\n", s)

	return nil
}

func main() {
	data, err := getData()

	if err != nil {
		fmt.Println("Error getData()")
		os.Exit(1)
	}

	printStats(data)
}
