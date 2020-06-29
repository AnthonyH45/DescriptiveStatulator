package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Printf("%q", strSplit)
	toRet := []int{1}
	return toRet, nil
}

func printStats(data []int) error {
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
