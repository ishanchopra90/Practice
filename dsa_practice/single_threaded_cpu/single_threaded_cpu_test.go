package singlethreadedcpu

import (
	"strconv"
	"testing"
)

func TestGetOrder(t *testing.T) {
	tasks := [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}
	output := getOrder(tasks)
	expectedOrder := []int{0, 2, 3, 1}
	assertOrder(expectedOrder, output, t)

	tasks = [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}
	output = getOrder(tasks)
	expectedOrder = []int{4, 3, 2, 0, 1}
	assertOrder(expectedOrder, output, t)

	tasks = [][]int{{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11}}
	output = getOrder(tasks)
	expectedOrder = []int{15, 14, 13, 1, 6, 3, 5, 12, 8, 11, 9, 4, 10, 7, 0, 2}
	assertOrder(expectedOrder, output, t)
}

func assertOrder(expectedOrder []int, actualOrder []int, t *testing.T) {
	expectedOrderStr := ""
	for i := 0; i < len(expectedOrder); i++ {
		expectedOrderStr = expectedOrderStr + strconv.Itoa(expectedOrder[i]) + ","
	}

	t.Logf("validating expected order %v\n", expectedOrderStr)
	if len(actualOrder) != len(expectedOrder) {
		t.Errorf("Expected output len %v found len %v\n", len(expectedOrder), len(actualOrder))
	}

	actualOrderStr := ""
	for i := 0; i < len(actualOrder); i++ {
		actualOrderStr = actualOrderStr + strconv.Itoa(actualOrder[i]) + ","
	}

	for i := 0; i < len(actualOrder); i++ {
		if actualOrder[i] != expectedOrder[i] {
			t.Errorf("Expected output %v\n returned output %v\n", expectedOrderStr, actualOrderStr)
			break
		}
	}
}
