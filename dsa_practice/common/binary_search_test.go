package common

import "testing"

func TestBinarySearchIdxLesserThan(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5}
	findIdx := BinarySearchIdxLesserThan(3, ary)
	if findIdx != 2 {
		t.Errorf("Expected: %v Found: %v\n", 2, findIdx)
	}

	findIdx = BinarySearchIdxLesserThan(0, ary)
	if findIdx != -1 {
		t.Errorf("Expected: %v Found: %v\n", -1, findIdx)
	}

	findIdx = BinarySearchIdxLesserThan(6, ary)
	if findIdx != 4 {
		t.Errorf("Expected: %v Found: %v\n", 4, findIdx)
	}

	ary = []int{2, 4, 6, 8, 10}
	findIdx = BinarySearchIdxLesserThan(3, ary)
	if findIdx != 0 {
		t.Errorf("Expected: %v Found: %v\n", 0, findIdx)
	}

	ary = []int{2, 4, 6, 8, 10}
	findIdx = BinarySearchIdxLesserThan(9, ary)
	if findIdx != 3 {
		t.Errorf("Expected: %v Found: %v\n", 3, findIdx)
	}
}

func TestBinarySearchIdxLesserThanDuplicatesTypeGeneric(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5}
	findIdxs := BinarySearchIdxLesserThanDuplicatesTypeGeneric(3, ary, intComparator)
	if findIdxs[0] != findIdxs[1] {
		t.Errorf("Expected start and end match")
	}

	if findIdxs[0] != 2 {
		t.Errorf("Expected: %v Found: %v\n", 2, findIdxs[0])
	}

	findIdxs = BinarySearchIdxLesserThanDuplicatesTypeGeneric(0, ary, intComparator)
	if findIdxs[0] != findIdxs[1] {
		t.Errorf("Expected start and end match")
	}

	if findIdxs[0] != -1 {
		t.Errorf("Expected: %v Found: %v\n", -1, findIdxs[0])
	}

	findIdxs = BinarySearchIdxLesserThanDuplicatesTypeGeneric(6, ary, intComparator)
	if findIdxs[0] != findIdxs[1] {
		t.Errorf("Expected start and end match")
	}

	if findIdxs[0] != 4 {
		t.Errorf("Expected: %v Found: %v\n", 4, findIdxs[0])
	}

	ary = []int{2, 4, 6, 8, 10}
	findIdxs = BinarySearchIdxLesserThanDuplicatesTypeGeneric(3, ary, intComparator)
	if findIdxs[0] != findIdxs[1] {
		t.Errorf("Expected start and end match")
	}

	if findIdxs[0] != 0 {
		t.Errorf("Expected: %v Found: %v\n", 0, findIdxs[0])
	}

	ary = []int{2, 4, 6, 8, 10}
	findIdxs = BinarySearchIdxLesserThanDuplicatesTypeGeneric(9, ary, intComparator)
	if findIdxs[0] != findIdxs[1] {
		t.Errorf("Expected start and end match")
	}

	if findIdxs[0] != 3 {
		t.Errorf("Expected: %v Found: %v\n", 3, findIdxs[0])
	}

	ary = []int{1, 7, 18, 18, 27, 28, 35, 44, 45, 53, 59, 60, 63, 66, 72, 73, 83, 88, 94, 100}
	findIdxs = BinarySearchIdxLesserThanDuplicatesTypeGeneric(18, ary, intComparator)
	if findIdxs[0] == findIdxs[1] {
		t.Errorf("Expected start and end to not match")
	}

	if findIdxs[0] != 2 || findIdxs[1] != 3 {
		t.Errorf("Expected: %v, %v Found: %v, %v\n", 2, 3, findIdxs[0], findIdxs[1])
	}
}

func intComparator(i, j int) int {
	if i == j {
		return 0
	}

	if i < j {
		return -1
	}

	return 1
}
