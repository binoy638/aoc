package main

import (
	"fmt"
	"math"
	"sort"
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

	var list1, list2 []int

	for _, value := range inputs {
		l := strings.Fields(value)
		firstNum, err := strconv.Atoi(l[0])
		if err != nil {
			panic("error while converting string to number")
		}
		list1 = append(list1, firstNum)

		secondNum, err := strconv.Atoi(l[1])
		if err != nil {
			panic("error while converting string to number")
		}

		list2 = append(list2, secondNum)

	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i := range len(list1) {
		sum += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	println(sum)

	totalTime := time.Since(startTime)

	fmt.Printf("Total time taken: %s \n", totalTime)
}
