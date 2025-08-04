package maxprofitassignment

import "testing"

func TestMaxProfitAssignment(t *testing.T) {
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	maxProfit := maxProfitAssignment(difficulty, profit, worker)
	assertExpectedProfit(100, maxProfit, t)

	difficulty = []int{85, 47, 57}
	profit = []int{24, 66, 99}
	worker = []int{40, 25, 25}
	maxProfit = maxProfitAssignment(difficulty, profit, worker)
	assertExpectedProfit(0, maxProfit, t)

	difficulty = []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63}
	profit = []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77}
	worker = []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}
	maxProfit = maxProfitAssignment(difficulty, profit, worker)
	assertExpectedProfit(1392, maxProfit, t)
}

func assertExpectedProfit(expected int, found int, t *testing.T) {
	if found != expected {
		t.Errorf("Expected: %v Found %v\n", expected, found)
	} else {
		t.Logf("\nMax profit = %v\n", found)
	}
}
