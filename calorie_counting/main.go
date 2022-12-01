package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, readErr := os.OpenFile("./input", os.O_RDONLY, os.ModePerm)
	if readErr != nil {
		fmt.Printf("error reading file %s", readErr)
		return
	}
	defer input.Close()

	calorieCountList, err := listTotalCalories(input)
	if err != nil {
		fmt.Printf("error getting greatest calorie count %s", err)
	}

	sort.Ints(calorieCountList)

	greatestCalorieCount := calorieCountList[len(calorieCountList)-1]

	topThreeCalorieCountList := calorieCountList[len(calorieCountList)-3:]
	sumOfTopThreeCounts := sumOfSlice(topThreeCalorieCountList)

	fmt.Printf("Greatest calorie count %d\n", greatestCalorieCount)
	fmt.Printf("Sum of top three %d\n", sumOfTopThreeCounts)
}

func listTotalCalories(reader io.Reader) ([]int, error) {
	var totalCalorieCount int
	var totalCaloriesList []int

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		// if it's not an empty line add value to total calorie count
		if len(line) > 0 {
			calories, parseErr := strconv.Atoi(line)
			if parseErr != nil {
				return totalCaloriesList, fmt.Errorf("error parsing line %w", parseErr)
			}

			totalCalorieCount += calories
		} else {
			// if it is an empty line add total calorie count to list and reset
			totalCaloriesList = append(totalCaloriesList, totalCalorieCount)
			totalCalorieCount = 0
		}
	}

	return totalCaloriesList, nil
}

func sumOfSlice(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}

	return sum
}
