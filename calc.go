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
	fmt.Printf("Min: %d\n", data[0])
	// printMax(data)
	fmt.Printf("Max: %d\n", data[len(data)-1])

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
