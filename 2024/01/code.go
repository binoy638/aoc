package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readInput(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	var res []string

	if err != nil {
		return res, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil

}

func main() {
	startTime := time.Now()
	lines, err := readInput("input.txt")

	if err != nil {
		fmt.Printf("Error getting input: %v", err)
	}

	var list1, list2 []int

	for _, value := range lines {
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

	endTime := time.Now()

	totalTime := endTime.Nanosecond() - startTime.Nanosecond()

	fmt.Printf("Total time taken: %d", totalTime)
}
