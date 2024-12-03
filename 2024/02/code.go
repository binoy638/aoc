package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/binoy638/aoc/utils"
)

func main() {
	startTime := time.Now()
	inputs, err := utils.ReadInputFromFile("input.txt")

	if err != nil {
		fmt.Printf("Error getting input: %v", err)
	}

	var list1 []int
	// create a map to store all unique values from list 2
	list2map := make(map[int]int)

	for _, value := range inputs {
		parts := strings.Fields(value)

		num1, err := strconv.Atoi(parts[0])

		if err != nil {
			panic("error while converting string to number")
		}

		// append num 1 to the list

		list1 = append(list1, num1)

		num2, err := strconv.Atoi(parts[1])

		if err != nil {
			panic("error while converting string to number")
		}

		// insert num2 in the map

		// check if the value exists in the map

		count, exists := list2map[num2]

		if !exists {
			// insert the value in the map with count 1
			list2map[num2] = 1
		}

		list2map[num2] = count + 1

	}

	sum := 0
	// loop list1 and get the sum

	for _, num := range list1 {
		// check if the num exists in the list2map
		count := list2map[num]

		sum += num * count

	}

	println(sum)

	totalTime := time.Since(startTime)

	fmt.Printf("Total time taken: %s \n", totalTime)

}
