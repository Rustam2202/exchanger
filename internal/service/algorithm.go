package service

import (
	"errors"
	"slices"
)

func FindCombinations(amount int, banknotes []int) ([][]int, error) {
	if amount <= 0 {
		return nil, errors.New("amount should be great as zero")
	}
	if len(banknotes) == 0 {
		return nil, errors.New("banknotes must not be empty")
	}
	if amount < slices.Min(banknotes) {
		return nil, errors.New("amount is less minimum banknote")
	}
	banknotes=slices.Compact(banknotes)

	var result [][]int
	var combination []int
	findCombinationsRecursive(amount, banknotes, 0, combination, &result)
	if result == nil {
		return nil, errors.New("couldn`t found exchange variants")
	}
	return result, nil
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
