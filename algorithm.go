package main

import "fmt"

func findCombinations(amount int, banknotes []int) [][]int {
	var result [][]int
	var combination []int
	findCombinationsRecursive(amount, banknotes, 0, combination, &result)
	return result
}

func findCombinationsRecursive(amount int, banknotes []int, start int, combination []int, result *[][]int) {
	if amount == 0 {
		comboCopy := make([]int, len(combination))
		copy(comboCopy, combination)
		*result = append(*result, comboCopy)
		return
	}

	for i := start; i < len(banknotes); i++ {
		if banknotes[i] <= amount {
			combination = append(combination, banknotes[i])
			findCombinationsRecursive(amount-banknotes[i], banknotes, i, combination, result)
			combination = combination[:len(combination)-1]
		}
	}
}

func main() {
	result := findCombinations(400, []int{5000, 2000, 1000, 500, 200, 100, 50})
	fmt.Printf("result: %v\n", result)
}
