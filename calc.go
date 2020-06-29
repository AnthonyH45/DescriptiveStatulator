package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() ([]int64, error) {
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

	var toRet []int64
	for _, i := range strSplit {
		toAdd, convErr := strconv.ParseInt(i, 10, 64)

		if convErr != nil {
			fmt.Printf("Error converting %q to 64 bit base 10 integer", i)
			return []int64{}, convErr
		}
		toRet = append(toRet, toAdd)
	}

	return toRet, nil
}

func printStats(data []int64) error {
	fmt.Printf("Statistics for %q", data)
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
