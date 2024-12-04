package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/binoy638/aoc/utils"
)

func isReportValid(reportStr string) bool {
	isValid := true
	// make a list using the str
	list := strings.Fields(reportStr)

	mode := -1 // 0 bad seq, 1 inc, 2 dec

	for i := 0; i < len(list)-1; i++ {

		curr := utils.ParseInt(list[i])

		next := utils.ParseInt(list[i+1])

		if i == 0 {
			if curr < next {
				mode = 1
			} else {
				mode = 2
			}
		} else {
			if mode == 1 && curr > next || mode == 2 && curr < next {
				isValid = false
				break
			}
		}

		diff := int(math.Abs(float64(curr) - float64(next)))

		// check the difference
		if diff < 1 || diff > 3 {
			isValid = false
			break
		}

		if i == len(list)-2 {
			mode = -1
		}
	}

	return isValid
}

func main() {
	startTime := time.Now()
	inputs, err := utils.ReadInputFromFile("input.txt")

	if err != nil {
		fmt.Printf("Error getting input: %v", err)
	}

	validReports := 0

	for _, reportStr := range inputs {
		isValid := isReportValid(reportStr)

		if isValid {
			validReports += 1
		}
	}

	println(validReports)

	totalTime := time.Since(startTime)

	fmt.Printf("Total time taken: %s \n", totalTime)

}
